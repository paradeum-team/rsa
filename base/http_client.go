package base

import (
	"bfs-rsa/utils"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

/**
 * HTTPS Basic Authentication in Go :Gish
 *
 * https://bl.ocks.org/nicerobot/4375261
 */
var (
	BFS_API_URL="http://39.100.9.55:5144"
)

func init()  {
	BFS_API_URL="http://39.100.9.55:5144" //dev 环境的rn,应该从环境变量获取
}

func ExecuteGet(api string, auth string) (int,[] byte, error) {
	url :=BFS_API_URL+api
	return HttpGetExecute(url, auth)
}


func ExecutePostFile(api string,params map[string]string, filePath string, auth string) (int,[] byte, error) {
	url :=BFS_API_URL+api
	return HttpPostFileExecute(url,params,filePath,auth)
}


func ExecutePost(api string,params map[string]string, auth string) (int,[] byte, error) {
	url :=BFS_API_URL+api
	return HttpPostExecute(url,params,auth)
}


func ExecutePut(api string,params map[string]string, auth string) (int,[] byte, error) {
	url :=BFS_API_URL+api
	return HttpPutExecute(url,params,auth)
}


func  HttpPostExecute(url string ,params map[string]string, auth string)(int,[] byte, error){
	client := GetHttpClient()

	data,err:=json.Marshal(params)
	if err != nil {
		log.Printf("[HttpPostExecute]  json Parse is  wrong. %v \n",err)
		return 100,nil, err
	}
	requestBody:=strings.NewReader(string(data))

	req, err := http.NewRequest("POST", url, requestBody)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	if auth != ""{
		req.Header.Set("Authorization", auth)
	}

	response, err := client.Do(req)

	if err != nil {
		log.Println("请求错误：", err)
		return 1200,nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("读取body失败：", err)
		return response.StatusCode,  nil, err
	}

	return response.StatusCode,body, nil
}


func  HttpPutExecute(url string ,params map[string]string, auth string)(int,[] byte, error){
	client := GetHttpClient()

	data,err:=json.Marshal(params)
	if err != nil {
		log.Printf("[HttpPutExecute]  json Parse is  wrong. %v \n",err)
		return 100,nil, err
	}

	requestBody:=strings.NewReader(string(data))
	req, err := http.NewRequest("PUT", url, requestBody)

	if err !=nil {
		fmt.Println("put request error ",err.Error())
		return 1200,nil, err
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	if auth != ""{
		req.Header.Set("Authorization", auth)
	}
	response, err := client.Do(req)

	if err != nil {
		log.Println("请求错误：", err)
		return 1200,nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("读取body失败：", err)
		return response.StatusCode,  nil, err
	}

	return response.StatusCode,body, nil
}


/**
 * http client get request
 * return httpStatusCode ,body ,err
 */
func HttpPostFileExecute(url string,params map[string]string, filePath string, auth string) (int,[] byte, error) {
	client := GetHttpClient()
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("[HttpPostExecute]  open file is  error. %v \n",err)
		return 100,nil, err
	}
	defer file.Close()


	requestBody := &bytes.Buffer{}
	writer := multipart.NewWriter(requestBody)
	part, err := writer.CreateFormFile("file", getTempPath()+file.Name())
	if err != nil {
		log.Printf("[HttpPostExecute]  CreateFormFile is  wrong. %v \n",err)
		return 100,nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		log.Printf("[HttpPostExecute]  close the writer is  wrong. %v \n",err)
		return 100,nil, err
	}

	req, err := http.NewRequest("POST", url, requestBody)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if auth != ""{
		req.Header.Set("Authorization", auth)
	}

	if err != nil {
		log.Println("[HttpPostExecute] Create Request is wrong :", err)
		return 100,nil, err
	}

	response, err := client.Do(req)

	if err != nil {
		log.Println("请求错误：", err)
		return 1200,nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("读取body失败：", err)
		return response.StatusCode,  nil, err
	}

	return response.StatusCode,body, nil

}

func HttpGetExecute(url string, auth string) (int,[] byte, error) {
	client := GetHttpClient()

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")

	if auth != ""{
		req.Header.Set("Authorization", auth)
	}

	if err != nil {
		log.Fatalln("New Request is wrong :", err)
		return 1200,nil, err
	}

	response, err := client.Do(req)

	if err != nil {
		log.Println("请求错误：", err)
		return 100,nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("读取body失败：", err)
		return response.StatusCode,  nil, err
	}

	return response.StatusCode,body, nil

}


/**
 * 获取 http client
 * 忽略证书
 *
 */
func GetHttpClient() *http.Client {
	proxyUrl := "socks5://127.0.0.1:10020"
	proxy, _ := url.Parse(proxyUrl)
	tr := &http.Transport{
		//Proxy: http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		MaxIdleConns:        20,
		MaxIdleConnsPerHost: 20,
	}

	//if pldconf.AppConfig.Server.ProxyModel{//本地开发模式，需要使用代理
	if false {
		tr.Proxy=http.ProxyURL(proxy)
	}
	client := &http.Client{
		Transport: tr,
		//Timeout:   time.Second * 5, //超时时间
	}
	return client
}

func getTempPath()string {
	timeContent:=utils.GetCurrentTimeUnix()
	return "./data/temp/" +timeContent

}