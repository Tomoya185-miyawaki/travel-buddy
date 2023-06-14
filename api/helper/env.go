package helper

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("読み込みができませんでした")
	}
	return os.Getenv("PORT")
}
