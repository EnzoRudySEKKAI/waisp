package main

import (
	"fmt"

	"github.com/EnzoRudySEKKAI/waisp/internal/brain"
)

func main() {
	companyBrain := brain.New()

	fmt.Println("waïsp Company Brain running")
	fmt.Printf("users=%d assistants=%d domains=%d\n", len(companyBrain.Users), len(companyBrain.Assistants), len(companyBrain.Domains))
}
