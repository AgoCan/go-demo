package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// ContainerInfos 容器信息
type ContainerInfos struct {
	Container []ContainerInfo `json:"container"`
}

// ContainerInfo 单个容器信息
type ContainerInfo struct {
	ID     string `json:"id"`
	Image  string `json:"image"`
	Size   int64  `json:"size"`
	Status string `json:"status"`
}

func main() {

	var conList ContainerInfos
	var conInfo ContainerInfo
	ctx := context.Background()
	cli, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	containers, _ := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	for _, con := range containers {
		conJSON, _, _ := cli.ContainerInspectWithRaw(ctx, con.ID, true)
		conInfo.ID = con.ID
		conInfo.Image = con.Image
		conInfo.Status = con.Status
		conInfo.Size = (*conJSON.SizeRootFs) / 1000 / 1000
		conList.Container = append(conList.Container, conInfo)
	}
	fmt.Printf("%#v", conList)
}
