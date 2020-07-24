package main

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

// func build() {
// 	ctx := context.Background()
// 	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
// 	fmt.Println(err)

// }

var wg sync.WaitGroup

func pull(ctx context.Context, cli *client.Client) {
	reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)
	wg.Done()
}

func exec(ctx context.Context, cli *client.Client, containerID string) {
	configExec := types.ExecConfig{}
	configExec.Tty = false
	configExec.AttachStdin = true
	configExec.AttachStdout = true
	configExec.AttachStderr = true
	configExec.Cmd = []string{"touch", "/1.txt"}
	res, err := cli.ContainerExecCreate(ctx, containerID, configExec)
	fmt.Println(err)
	fmt.Println(res)
	r, err := cli.ContainerExecAttach(ctx, res.ID, types.ExecStartCheck{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, r.Reader)
}

func build(ctx context.Context, cli *client.Client, dockerfilePath, imageName string) (err error) {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()
	dockerFile := "myDockerfile"
	dockerFileReader, err := os.Open(dockerfilePath)
	if err != nil {
		log.Fatal(err, " :unable to open Dockerfile")
	}
	readDockerFile, err := ioutil.ReadAll(dockerFileReader)
	if err != nil {
		log.Fatal(err, " :unable to read dockerfile")
	}
	tarHeader := &tar.Header{
		Name: dockerFile,
		Size: int64(len(readDockerFile)),
	}
	err = tw.WriteHeader(tarHeader)
	if err != nil {
		log.Fatal(err, " :unable to write tar header")
	}
	_, err = tw.Write(readDockerFile)
	if err != nil {
		log.Fatal(err, " :unable to write tar body")
	}
	dockerFileTarReader := bytes.NewReader(buf.Bytes())

	imageBuildResponse, err := cli.ImageBuild(
		ctx,
		dockerFileTarReader,
		types.ImageBuildOptions{
			Context:    dockerFileTarReader,
			Dockerfile: dockerFile,
			Remove:     true,
			Tags:       []string{imageName},
		})
	if err != nil {
		log.Fatal(err, " :unable to build docker image")
	}
	defer imageBuildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	if err != nil {
		log.Fatal(err, " :unable to read image build response")
	}
	return nil
}

func main() {
	config := container.Config{}
	config.Image = "alpine"
	config.Cmd = []string{"sh"}
	config.Tty = true
	config.ExposedPorts = nat.PortSet{"80/tcp": struct{}{}}
	config.Env = []string{
		"eeee=111",
		"happy=no",
	}
	config.Env = append(config.Env, "abc=111")
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
	wg.Add(1)
	go pull(ctx, cli)
	wg.Wait()
	resp, err := cli.ContainerCreate(ctx, &config, &hostConfig, nil, "")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.ID)
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	// commitResp, err := cli.ContainerCommit(ctx, resp.ID, types.ContainerCommitOptions{Reference: "helloworld"})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(commitResp.ID)
	err = build(ctx, cli, "./Dockerfile", "test:v-go")

	fmt.Println(err)
}
