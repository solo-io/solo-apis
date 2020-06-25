package main

import (
	"github.com/solo-io/solo-apis/codegen"
	"log"
)

func main() {
	log.Println("Starting skv2 code generation")
	skv2Cmd := codegen.Command()
	if err := skv2Cmd.Execute(); err != nil {
		log.Fatal(err)
	}
	log.Println("Finished generating skv2 code")
}
