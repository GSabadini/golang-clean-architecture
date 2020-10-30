package main

import (
	_ "encoding/json"
	"github.com/GSabadini/go-challenge/infrastructure"
)

func main() {
	infrastructure.NewHTTPServer().Start()
}
