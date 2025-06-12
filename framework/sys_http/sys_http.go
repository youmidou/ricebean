package sys_http

import (
	"errors"
	"fmt"
	//"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// 接收体
type ReceiveBody struct {
	body []byte
}

func (t *ReceiveBody) Json() {
	//sys_json.Unmarshal(t.body,interface{})
}

func Get(url string) []byte {
	// 发送 GET 请求
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil
	}
	return body
}
func UrlPostRetry(url string, payload io.Reader, retry int, ContentType string) ([]byte, error) {
	for i := 0; i < retry; i++ {
		ret, err := Post(url, payload, ContentType)
		if err == nil {
			continue
		}
		return ret, nil
	}
	return nil, errors.New(fmt.Sprintf("UrlPostRetry %s 失败", url))
}

func Post(url string, payload io.Reader, ContentType string) ([]byte, error) {
	//url := "http://www.baidu.com"
	//payload := strings.NewReader("key1=value1&key2=value2") // 设置 POST 请求的内容
	requestTimeout := 10 * time.Second

	// 创建一个 POST 请求
	request, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}
	// 设置请求头
	if ContentType == "" {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else { //"text/Json; charset=utf-8"
		request.Header.Set("Content-Type", ContentType)
	}

	// 发送请求
	client := &http.Client{
		Timeout: requestTimeout,
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}
	//fmt.Println("Response from", url, ":", string(body))
	return body, nil
}
