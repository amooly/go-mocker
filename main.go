package main

import (
	"encoding/json"
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const sampleFilePath = "/src/github.com/amooly/go-mocker/sample/mock.json"

func main() {
	// 1. 读取文件
	var mockFile = getMockFile()

	// 2. 基于文件内容创建服务
	config, err := readConfig(mockFile)
	if err != nil {
		log.Fatal(err)
	}
	server := HttpServer{HttpConfig: config}

	// 3. 运行服务
	log.Fatal(server.Run())
}

func getMockFile() string {
	var mockFile string
	if len(os.Args) == 1 {
		gopath := os.Getenv("GOPATH")
		if gopath == "" {
			gopath = build.Default.GOPATH
		}
		mockFile = gopath + sampleFilePath
		log.Println("Using sample file:", mockFile)
	} else {
		mockFile = os.Args[1]
		if !strings.HasPrefix(mockFile, "/") {
			path, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			mockFile = path + "/" + mockFile
		}
		log.Println("Using user-defined file:", mockFile)
	}
	return mockFile
}

func readConfig(file string) (*Config, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	if err := json.Unmarshal(data, config); err != nil {
		return nil, err
	} else {
		return config, nil
	}
}
