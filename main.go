package main

import (
	"fmt"
	"log"
	"os"
)

func main() { 
  // This fails because PORT is not in the current shell env.
  // It's not taken from the .env file
  // we can add the env by using export command
  // $ export PORT=8000
  // the above command will add the env variable to the current shell 
  // but to get it from the .env file we are going to use the godotenv package.
  port := os.Getenv("PORT")
  if port == ""{
    log.Fatal("Port is not in the env")
  }
  fmt.Println("PORT: ",port)
}
