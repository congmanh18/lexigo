package utils

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	// mặc định file .env sẽ nằm ở ./conf/.env
	envFile string
)

func init() {
	// khai báo flag -env
	flag.StringVar(&envFile, "env", "./conf/.env", "path to .env file")
	// parse luôn ở init để flag có hiệu lực trước khi main chạy
	flag.Parse()

	// load envFile
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("⚠️  Warning: no .env file found at %s (will fallback to system env)\n", envFile)
	}
}

// GetEnv đọc biến môi trường, nếu không có trả về defaultValue
func GetEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}
