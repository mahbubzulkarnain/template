package main

import (
	"fmt"
	"log"

	"template-go/pkg/dotenv"
)

var PORT = ":" + dotenv.Load("PORT", "80")

func main() {
	log.Println(fmt.Sprintf("0.0.0.0%v 🚀 server started...", PORT))
}
