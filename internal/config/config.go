package config
import (
	"os"//os package ask the operating system for environment variables
	"github.com/joho/godotenv"// godotenv package is used to load environment variables from a .env file
)
type Config struct {
	Port string
	Env  string //env means environment, like development, production, etc.
}


func MustLoad() Config{
	godotenv.Load()//load the .env file and set the environment variables, if the .env file is not found, it will not panic, it will just continue
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