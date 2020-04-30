package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

func main() {
	ctx := context.Background()
	fmt.Printf("%v,%T\n", ctx, ctx)
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println(1)
		panic(err)
	}
	fmt.Println("pull 1")
	reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	fmt.Println("pull 2")
	if err != nil {
		fmt.Println(2)
		panic(err)
	}
	io.Copy(os.Stdout, reader)
	fmt.Println("b")
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world"},
		Tty:   true,
	}, nil, nil, "")
	if err != nil {
		fmt.Println(3)
		panic(err)
	}
	fmt.Println("c")
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		fmt.Println(4)
		panic(err)
	}
	fmt.Println("d")
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			fmt.Println(5)
			panic(err)
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		fmt.Println(5)
		panic(err)
	}
	var tmp []byte
	out.Read(tmp)
	fmt.Println(tmp, statusCh)
	// fmt.Println("p")
	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	// fmt.Println("最后", w)
}
