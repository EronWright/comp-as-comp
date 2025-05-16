package main

import (
	"context"
	"fmt"
	"os"

	"github.com/pulumi/pulumi-go-provider/infer"
)

func main() {
	provider, err := infer.NewProviderBuilder().
		WithNamespace("mikhailshilkov").
		WithComponents(
			infer.ComponentF(NewRandomComponent),
			infer.ComponentF(NewStaticPage),
		).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}

	err = provider.Run(context.Background(), "go-components", "0.1.0")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}
