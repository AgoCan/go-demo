package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {

	ctx := context.Background()
	cli, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	res, _ := cli.ImageList(ctx, types.ImageListOptions{All: true})
	for _, image := range res {
		fmt.Printf("%T", image.RepoTags[0])
	}
}
