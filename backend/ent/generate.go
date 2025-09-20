//go:build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	err := entc.Generate("./src/models", &gen.Config{})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}