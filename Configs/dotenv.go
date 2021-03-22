package Configs

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/joho/godotenv"
)

func DotEnv(input string) string {

	err := godotenv.Load(filepath.Join("/home/go/src/mongo_sample_project/Configs/", "setup.env"))
	if err != nil {
		fmt.Println(err)
	}
	return os.Getenv(input)
}
