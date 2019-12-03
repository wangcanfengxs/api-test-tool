package runner

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/yalp/jsonpath"
	"testing"
)

func Test_getUrl(t *testing.T) {
	tests := []struct {
		name    string
		request *ApiTestRequest
		expect  string
	}{
		{
			name: "Test_getQueryParameterFromUrl_1",
			request: &ApiTestRequest{
				Url:         "http://example.com?name=example&key=example&key=com",
				QueryParams: nil,
			},
			expect: "http://example.com?key=example&key=com&name=example",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualUrl, err := getUrl(test.request)
			if err != nil {
				t.Error(err)
				return
			}
			if !cmp.Equal(test.expect, actualUrl) {
				t.Errorf("expected %s, got %s ", test.expect, actualUrl)
				return
			}
		})
	}
}

func Test_jsonpath(t *testing.T) {
	tests := []struct {
		name     string
		json     []byte
		jsonpath string
		expect   interface{}
	}{
		{
			name:     "Test_jsonpath_1",
			json:     []byte(`{"hello": "world"}`),
			jsonpath: "$.hello",
			expect:   "world",
		},
		{
			name:     "Test_jsonpath_2",
			json:     []byte(`{"hello": 2}`),
			jsonpath: "$.hello",
			expect:   2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var data interface{}
			err := json.Unmarshal(test.json, &data)
			if err != nil {
				t.Error(err)
				return
			}
			actualValue, err := jsonpath.Read(data, test.jsonpath)
			if err != nil {
				t.Error(err)
				return
			}
			if fmt.Sprintf("%v", test.expect) != fmt.Sprintf("%v", actualValue) {
				t.Errorf("expected %v, got %v ", test.expect, actualValue)
				return
			}
		})
	}
}

func Test(t *testing.T) {
	raw := []byte(`{"hello":"world"}`)

	//var data interface{}
	//json.Unmarshal(raw, &data)

	out, err := jsonpath.Read(string(raw), "$.hello")
	if err != nil {
		panic(err)
	}

	fmt.Print(out)
}
