package main

import (
	"context"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func main() {
	var err error
	u, _ := url.Parse("https://example.cos.ap-chengdu.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	// 1.永久密钥
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "SecretID",
			SecretKey: "SecretKey",
		},
	})
	name := "test/go.mod"
	_, err = client.Object.PutFromFile(context.Background(), name, "go.mod", nil)
	if err != nil {
		panic(err)
	}
}
