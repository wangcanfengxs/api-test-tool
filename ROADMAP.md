# v1.0.0

## 支持命名行运行，传入文件（json格式）

文件格式如下：
```json
{
    "requests":[
        {
            "name":"test_case_1",
            "url":"http://api.juheapi.com/japi/toh",
            "queryParams":[
                {
                    "key":"value1",
                    "value":"value2"
                }
            ],
            "data":"{}",
            "headerData":[
                {
                    "key":"Content-Type",
                    "value":"application/json"
                }
            ],
            "method":"get",
            "description":"测试接口, get请求",
            "expect":{
                "httpStatusCode": 200, 
                "asserts": [
                    {
                        "jsonpath": 
                        "match": 
                    }    
                ]   
            }
        }
    ]
}
```

url中如果也包含了参数，query会覆盖之

## 支持简单的统计

