package main

import (
	"context"
	"fmt"
	"os"

	"github.com/moby/moby/client"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "USAGE: %s CONTAINER-ID\n", os.Args[0])
		os.Exit(1)
	}

	id := os.Args[1]

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create docker client: %s\n", err)
		os.Exit(1)
	}

	cli.NegotiateAPIVersion(context.Background())

	container, err := cli.ContainerInspect(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "docker inspect for '%s' failed: %s\n", id, err)
		os.Exit(1)
	}

	fmt.Printf("%v;%d\n", container.State.Running, container.State.Pid)
}
