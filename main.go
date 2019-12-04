package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/wangcanfengxs/api-test-tool/pkg/runner"
	"io/ioutil"
)

var (
	file      string
	directory string
)

func usage() string {
	return fmt.Sprintf("Tips: apitest -f test.json")
}

func getAllFiles(files *[]string, dirPath string) {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Println(err)
	}

	for _, fileInfo := range dir {
		if fileInfo.IsDir() {
			getAllFiles(files, dirPath+"/"+fileInfo.Name())
		} else {
			*files = append(*files, dirPath+"/"+fileInfo.Name())
		}
	}
}

func main() {
	flag.StringVar(&file, "f", "", "-f <file>")
	flag.StringVar(&directory, "d", "", "-d <directory>")
	flag.Parse()

	if file == "" && directory == "" {
		fmt.Println(usage())
		return
	}

	files := make([]string, 0)
	if directory != "" {
		getAllFiles(&files, directory)
	} else {
		files = append(files, file)
	}

	for _, filePath := range files {
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println(err)
		}

		// check data
		apiTestRunner, err := checkFileData(data)
		// fmt.Println(apiTestRunner)
		if err != nil {
			fmt.Println(err)
		}

		apiTestRunner.Run()
		fmt.Printf("Finished running requests in %s\n", filePath)
	}
}

func checkFileData(data []byte) (*runner.ApiTestRunner, error) {
	apiTestRunner := &runner.ApiTestRunner{}
	return apiTestRunner, json.Unmarshal(data, apiTestRunner)
}
