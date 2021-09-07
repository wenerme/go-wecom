package wecom

import (
	"testing"
	"time"

	"github.com/pkg/errors"

	"github.com/stretchr/testify/assert"
)

func TestGenericToken(t *testing.T) {
	var token *GenericToken
	assert.False(t, token.IsValid())
	token = &GenericToken{}
	changed, err := token.Refresh(&GenericToken{}, func() (Token, error) {
		return nil, nil
	})
	assert.False(t, changed)
	assert.NoError(t, err)

	token.SetFromToken(&GenericToken{
		Token:     "ABC",
		ExpiresIn: 1200,
	})
	now := timeNow().Unix()
	assert.Equal(t, token, &GenericToken{
		Token:     "ABC",
		ExpiresIn: 1200,
		ExpiresAt: now + 1200,
	})
}

// nolint:funlen
func TestTokenCache(t *testing.T) {
	store := &SyncMapStore{}
	{
		data := store.Dump()
		assert.NoError(t, store.Restore(data))
	}
	{
		var k, v, neo string
		v = "VAL"
		assert.NoError(t, store.Set(k, v))
		found, err := store.Get(k, &neo)
		assert.NoError(t, err)
		assert.True(t, found)
		assert.Equal(t, v, neo)

		assert.NoError(t, store.Set(k, nil))

		found, err = store.Get(k, &neo)
		assert.NoError(t, err)
		assert.False(t, found)
	}

	tc := &TokenCache{
		Store: store,
	}

	// no expire
	{
		token, err := tc.Refresh(&GenericToken{}, nil)
		assert.Error(t, err)
		assert.Equal(t, "", token)

		token, err = tc.Refresh(&GenericToken{
			Name: "A",
		}, func() (Token, error) {
			return &GenericToken{Token: "1"}, nil
		})
		assert.NoError(t, err)
		assert.Equal(t, "1", token)

		// use cache
		token, err = tc.Refresh(&GenericToken{
			Name: "A",
		}, func() (Token, error) {
			return &GenericToken{Token: "2"}, nil
		})
		assert.NoError(t, err)
		assert.Equal(t, "1", token)

		// owner change
		token, err = tc.Refresh(&GenericToken{
			Name:    "A",
			OwnerID: "1",
		}, func() (Token, error) {
			return &GenericToken{Token: "2"}, nil
		})
		assert.NoError(t, err)
		assert.Equal(t, "2", token)

		// cached
		token, err = tc.Refresh(&GenericToken{
			Name:    "A",
			OwnerID: "1",
		}, func() (Token, error) {
			return &GenericToken{Token: "3"}, nil
		})
		assert.NoError(t, err)
		assert.Equal(t, "2", token)
		// deps
		token, err = tc.Refresh(&GenericToken{
			Name:    "A",
			OwnerID: "1",
			Depends: "2",
		}, func() (Token, error) {
			return &GenericToken{Token: "3"}, nil
		})
		assert.NoError(t, err)
		assert.Equal(t, "3", token)

		// cached
		token, err = tc.Refresh(&GenericToken{
			Name:    "A",
			OwnerID: "1",
			Depends: "2",
		}, func() (Token, error) {
			return &GenericToken{Token: "4"}, nil
		})
		assert.NoError(t, err)
		assert.Equal(t, "3", token)
		// no deps
		token, err = tc.Refresh(&GenericToken{
			Name:    "A",
			OwnerID: "1",
		}, func() (Token, error) {
			return &GenericToken{Token: "4"}, nil
		})
		assert.NoError(t, err)
		assert.Equal(t, "4", token)
	}
	// with expire
	{
		n := time.Now()
		timeNow = func() time.Time {
			return n
		}
		defer func() {
			timeNow = time.Now
		}()

		token, err := tc.Refresh(&GenericToken{
			Name: "B",
		}, func() (Token, error) {
			return &GenericToken{Token: "1", ExpiresIn: 100}, nil
		})
		assert.NoError(t, err)
		assert.Equal(t, "1", token)

		n = n.Add(50 * time.Second)

		// still valid
		token, err = tc.Refresh(&GenericToken{
			Name: "B",
		}, func() (Token, error) {
			return &GenericToken{Token: "2", ExpiresIn: 100}, nil
		})
		assert.NoError(t, err)
		assert.Equal(t, "1", token)

		n = n.Add(30 * time.Second)

		// 80%
		token, err = tc.Refresh(&GenericToken{
			Name: "B",
		}, func() (Token, error) {
			return nil, errors.New("still valid")
		})
		assert.NoError(t, err)
		assert.Equal(t, "1", token)

		token, err = tc.Refresh(&GenericToken{
			Name: "B",
		}, func() (Token, error) {
			return &GenericToken{Token: "2", ExpiresIn: 100}, nil
		})
		assert.NoError(t, err)
		assert.Equal(t, "2", token)

		n = n.Add(50 * time.Second)
		// still valid
		token, err = tc.Refresh(&GenericToken{
			Name: "B",
		}, func() (Token, error) {
			return &GenericToken{Token: "3", ExpiresIn: 100}, nil
		})
		assert.NoError(t, err)
		assert.Equal(t, "2", token)

		n = n.Add(51 * time.Second)

		// expired
		_, err = tc.Refresh(&GenericToken{
			Name: "B",
		}, func() (Token, error) {
			return &GenericToken{Token: "3", ExpiresIn: 100}, errors.New("failed")
		})
		assert.Error(t, err)
		//
		token, err = tc.Refresh(&GenericToken{
			Name: "B",
		}, func() (Token, error) {
			return &GenericToken{Token: "3", ExpiresIn: 100}, nil
		})
		assert.NoError(t, err)
		assert.Equal(t, "3", token)
	}
}
