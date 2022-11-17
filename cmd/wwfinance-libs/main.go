package main

import (
	"log"
	"os"

	dotenv "github.com/joho/godotenv"
	"github.com/wenerme/go-wecom/WeWorkFinanceSDK/libs"
)

func main() {
	_ = dotenv.Load()
	libPath, _ := os.LookupEnv("WWF_LIBRARY_PATH")
	if libPath == "" {
		libPath, _ = os.Getwd()
	}
	if libPath != "" {
		log.Printf("Extract library to %s", libPath)
		_ = os.MkdirAll(libPath, 0o755)
		err := libs.ExtractTo(libPath)
		if err != nil {
			panic(err)
		}
	}
}
