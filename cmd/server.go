package main

import (
	"fmt"
	"os"

	"github.com/MokkeMeguru/golang-local-fs/internal/api/infrastructure/env"
	"github.com/MokkeMeguru/golang-local-fs/internal/api/infrastructure/router"
)

func main() {
	env, err := env.NewEnv()
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(0)
	}

	router, err := router.NewRouter(env)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(0)
	}

	router.Execute()
}
