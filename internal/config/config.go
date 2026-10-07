package config
import (
	"os"
	"github.com/joho/godotenv"
)
type Config struct {
	Port string
	Env  string //env means environment, like development, production, etc.
}


func MustLoad() Config{
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" { //check if the PORT environment variable is set
		panic("PORT environment variable is not set")
	}
	env := os.Getenv("ENV")
	if env == "" { //check if the ENV environment variable is set
		panic("ENV environment variable is not set")
	}
	return Config{
		Port: port,
		Env:  env,
	}
	
}