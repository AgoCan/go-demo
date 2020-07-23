package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/go-connections/nat"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main() {
	config := container.Config{}
	config.Image = "alpine"
	config.Cmd = []string{"sh"}
	config.Tty = true
	config.ExposedPorts = nat.PortSet{"80/tcp": struct{}{}}

	hostConfig := container.HostConfig{}
	hostConfig.Privileged = true
	hostConfig.PortBindings = nat.PortMap{
		"80/tcp": []nat.PortBinding{{
			HostIP:   "0.0.0.0",
			HostPort: "828/tcp",
		}},
	}
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)

	resp, err := cli.ContainerCreate(ctx, &config, &hostConfig, nil, "")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.ID)
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	commitResp, err := cli.ContainerCommit(ctx, resp.ID, types.ContainerCommitOptions{Reference: "helloworld"})
	if err != nil {
		panic(err)
	}
	fmt.Println(commitResp.ID)
}
