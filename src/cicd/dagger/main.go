package main

import (
	"context"
	"os"

	"dagger.io/dagger"
	_ "dagger.io/dagger"
)

func main() {
	ctx := context.Background()

	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	repo := client.Git("https://github.com/koki-develop/dagger-go-sdk-example.git")
	src := repo.Branch("main").Tree()

	golang := client.Container().From("golang:1.19")
	golang = golang.WithMountedDirectory("/app", src).WithWorkdir("/app")

	golang = golang.Exec(dagger.ContainerExecOpts{
		Args: []string{"go", "build", "-o", "build/"},
	}).Exec(dagger.ContainerExecOpts{
		Args: []string{"ls", "./build"},
	})

	if _, err := golang.ExitCode(ctx); err != nil {
		panic(err)
	}
}
