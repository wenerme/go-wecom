package libs

import (
	_ "embed"
	"os"
)

//go:embed libWeWorkFinanceSdk_C.so
var Linux []byte

//go:embed version.txt
var Version string

//go:embed md5sum.txt
var Checksum string

func ExtractTo(dir string) error {
	files := map[string][]byte{
		dir + "/libWeWorkFinanceSdk_C.so": Linux,
		dir + "/version.txt":              []byte(Version),
		dir + "/md5sum.txt":               []byte(Checksum),
	}
	for name, data := range files {
		if _, err := os.Stat(name); err == nil {
			continue
		}
		if err := os.WriteFile(name, data, 0o600); err != nil {
			return err
		}
	}
	return nil
}
