package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/wangcanfengxs/api-test-tool/pkg/runner"
	"io/ioutil"
)

var (
	file string
)

func usage() string {
	return fmt.Sprintf("Tips: apitest -f test.json")
}

func main() {
	flag.StringVar(&file, "f", "", "-f [file]")
	flag.Parse()
	if file == "" {
		fmt.Println(usage())
		return
	}

	// 读取文件
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	// check data
	apiTestRunner, err := checkFileData(data)
	//fmt.Println(apiTestRunner)
	if err != nil {
		panic(err)
	}

	apiTestRunner.Run()
}

func checkFileData(data []byte) (*runner.ApiTestRunner, error) {
	apiTestRunner := &runner.ApiTestRunner{}
	return apiTestRunner, json.Unmarshal(data, apiTestRunner)
}
