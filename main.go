package main

import (
	_ "encoding/json"
	"github.com/GSabadini/golang-clean-architecture/infrastructure"
)

func main() {
	infrastructure.NewHTTPServer().Start()
}
