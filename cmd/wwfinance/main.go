package main

import (
	"fmt"

	dotenv "github.com/joho/godotenv"
	"github.com/wenerme/go-wecom/WeWorkFinanceSDK"
)

func main() {
	_ = dotenv.Load()

	client, err := WeWorkFinanceSDK.NewClientFromEnv()
	if err != nil {
		panic(err)
	}

	data, err := client.GetChatData(WeWorkFinanceSDK.GetChatDataOptions{
		Limit:   10,
		Timeout: 5,
	})
	if err != nil {
		panic(err)
	}
	for _, v := range data {
		fmt.Println(v.Message)
	}
}
