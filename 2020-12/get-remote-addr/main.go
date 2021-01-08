package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httptrace"
)

// 方法一：通过httptrace.ClientTrace获取服务IP地址
func m1() {
	trace := &httptrace.ClientTrace{
		DNSStart: func(_ httptrace.DNSStartInfo) {},
		DNSDone:  func(_ httptrace.DNSDoneInfo) {},
		ConnectStart: func(net, addr string) {
			fmt.Printf("ConnectStart addr=%s\n", addr)
		},
		ConnectDone: func(net, addr string, err error) {
			fmt.Printf("ConnectDone addr=%s\n", addr)
		},
		GotConn:              func(_ httptrace.GotConnInfo) {},
		GotFirstResponseByte: func() {},
		TLSHandshakeStart:    func() {},
		TLSHandshakeDone:     func(_ tls.ConnectionState, _ error) {},
	}

	req, err := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	req = req.WithContext(httptrace.WithClientTrace(context.Background(), trace))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

// 方法二：通过DialContext获取服务IP地址
func m2() {
	req, err := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				conn, err := net.Dial(network, addr)
				req.RemoteAddr = conn.RemoteAddr().String()
				return conn, err
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("RemoteAddr:", req.RemoteAddr)
}

func main() {
	m1()
	m2()

	//查找DNS A记录
	iprecords, _ := net.LookupIP("landv.cn")
	for _, ip := range iprecords {
		fmt.Println(ip)
	}
	//查找DNS CNAME记录
	canme, _ := net.LookupCNAME("www.baidu.com")
	fmt.Println(canme)
	//查找DNS PTR记录
	ptr, e := net.LookupAddr("8.8.8.8")
	if e != nil {
		fmt.Println(e)
	}
	for _, ptrval := range ptr {
		fmt.Println(ptrval)
	}
	//查找DNS NS记录
	nameserver, _ := net.LookupNS("baidu.com")
	for _, ns := range nameserver {
		fmt.Println("ns记录", ns)
	}
	//查找DNS MX记录
	mxrecods, _ := net.LookupMX("google.com")
	for _, mx := range mxrecods {
		fmt.Println("mx:", mx)
	}
	//查找DNS TXT记录
	txtrecords, _ := net.LookupTXT("baidu.com")

	for _, txt := range txtrecords {
		fmt.Println("txt:", txt)
	}
}
