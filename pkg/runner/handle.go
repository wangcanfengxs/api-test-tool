package runner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/yalp/jsonpath"
)

func getUrl(request *ApiTestRequest) (string, error) {
	urlParsed, err := url.Parse(request.Url)
	if err != nil {
		return "", err
	}
	values := urlParsed.Query()
	for _, queryParam := range request.QueryParams {
		values.Set(queryParam.Key, queryParam.Value)
	}
	urlParsed.RawQuery = values.Encode()
	return urlParsed.String(), nil
}

func doGet(request *ApiTestRequest) error {
	rawUrl, err := getUrl(request)
	if err != nil {
		return err
	}
	client := http.Client{}
	httpRequest, err := http.NewRequest(GET, rawUrl, nil)
	if err != nil {
		return err
	}
	for _, header := range request.HeaderData {
		httpRequest.Header.Add(header.Key, header.Value)
	}
	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return err
	}
	return checkHttpResponse(request.Expect, httpResponse)
}

func doPost(request *ApiTestRequest) error {
	rawUrl, err := getUrl(request)
	if err != nil {
		return err
	}
	client := http.Client{}
	httpRequest, err := http.NewRequest(POST, rawUrl, strings.NewReader(request.Data))
	if err != nil {
		return err
	}
	for _, header := range request.HeaderData {
		httpRequest.Header.Add(header.Key, header.Value)
	}
	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return err
	}
	return checkHttpResponse(request.Expect, httpResponse)
}

func doPut(request *ApiTestRequest) error {
	rawUrl, err := getUrl(request)
	if err != nil {
		return err
	}
	client := http.Client{}
	httpRequest, err := http.NewRequest(PUT, rawUrl, strings.NewReader(request.Data))
	if err != nil {
		return err
	}
	for _, header := range request.HeaderData {
		httpRequest.Header.Add(header.Key, header.Value)
	}
	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return err
	}
	return checkHttpResponse(request.Expect, httpResponse)
}

func doDelete(request *ApiTestRequest) error {
	rawUrl, err := getUrl(request)
	if err != nil {
		return err
	}
	client := http.Client{}
	httpRequest, err := http.NewRequest(DELETE, rawUrl, nil)
	if err != nil {
		return err
	}
	for _, header := range request.HeaderData {
		httpRequest.Header.Add(header.Key, header.Value)
	}
	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return err
	}
	return checkHttpResponse(request.Expect, httpResponse)
}

func checkHttpResponse(expect *ExpectResponse, httpResponse *http.Response) error {
	if expect != nil {
		if expect.HttpStatusCode != httpResponse.StatusCode {
			return fmt.Errorf("Unexpected http status code: %d ", httpResponse.StatusCode)
		}
		if len(expect.Asserts) > 0 {
			body, err := ioutil.ReadAll(httpResponse.Body)
			if err != nil {
				return err
			}
			for _, assert := range expect.Asserts {
				var data interface{}
				err := json.Unmarshal(body, &data)
				if err != nil {
					return err
				}
				value, err := jsonpath.Read(data, assert.Jsonpath)
				if err != nil {
					return err
				}
				if fmt.Sprintf("%v", value) != fmt.Sprintf("%v", assert.Match) {
					return fmt.Errorf("response body not match, jsonpath %q, value: %q, expect: %q ", assert.Jsonpath, value, assert.Match)
				}
			}
		}
	}
	return nil
}
