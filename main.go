package main

import (
	"fmt"

	"github.com/Ferlab-Ste-Justine/ferload-client-cli-go/config"
)

func main() {
	cfg := config.LoadDefaultConfigUser()
	fmt.Printf("%+v\n", cfg)
}
