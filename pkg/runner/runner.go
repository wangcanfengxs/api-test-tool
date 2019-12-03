package runner

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ApiTestRunner struct {
	Requests []ApiTestRequest `json:"requests"`
}

type RequestQueryParameter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (requestQueryParameter RequestQueryParameter) String() string {
	b, err := json.MarshalIndent(requestQueryParameter, "", "    ")
	if err != nil {
		return ""
	}
	return string(b)
}

type RequestHeader struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ExpectResponse struct {
	HttpStatusCode int       `json:"httpStatusCode"`
	Asserts        []*Assert `json:"asserts"`
}

type Assert struct {
	Jsonpath string `json:"jsonpath"`
	Match    string `json:"match"`
}

type ApiTestRequest struct {
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Url         string                  `json:"url"`
	QueryParams []RequestQueryParameter `json:"queryParams"`
	Data        string                  `json:"data"`
	HeaderData  []RequestHeader         `json:"headerData"`
	Method      string                  `json:"method"`
	Expect      *ExpectResponse         `json:"expect"`
}

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

func (apiTestRunner *ApiTestRunner) Run() {
	fmt.Println("Start to run api test...")
	success, fail := 0, 0
	for index, request := range apiTestRunner.Requests {
		fmt.Printf("-------------------------------START(%d)---------------------------------\n", index+1)
		fmt.Printf("RUN %s(%s)\n", request.Name, request.Description)
		var err error
		switch strings.ToUpper(request.Method) {
		case GET:
			err = doGet(&request)
		case POST:
			err = doPost(&request)
		case PUT:
			err = doPut(&request)
		case DELETE:
			err = doDelete(&request)
		default:
			err = fmt.Errorf("Do not support http method %s. ", request.Method)
		}
		if err == nil {
			success += 1
			fmt.Printf("Run successfully.\n")
		} else {
			fail += 1
			fmt.Printf("Run failed. reason: %s\n", err)
		}
		fmt.Printf("--------------------------------END-------------------------------\n")

	}
	fmt.Println("--------------------------------------------------------------------")
	fmt.Printf("Total requests: %d, successful requests: %d, failed requests: %d\n", success+fail, success, fail)
	fmt.Println("--------------------------------------------------------------------")

}

func (apiTestRunner ApiTestRunner) String() string {
	b, err := json.MarshalIndent(apiTestRunner, "", "    ")
	if err != nil {
		return ""
	}
	return string(b)
}
