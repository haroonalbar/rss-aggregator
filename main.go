package main

import (
	"fmt"
	"log"
	"os"
)

func main() { 
  // This fails because PORT is not in the current shell env.
  // It's not taken from the .env file
  port := os.Getenv("PORT")
  if port == ""{
    log.Fatal("Port is not in the env")
  }
  fmt.Println("PORT: ",port)
}
