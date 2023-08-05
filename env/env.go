package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("./variables"); err != nil {
		log.Fatalln(err)
	}
}

func GetEnv(key, def string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}

	log.Println(key, ": valor default")
	return def
}
