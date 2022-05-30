package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func EnvLoad() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
		panic(err)
	}
}
