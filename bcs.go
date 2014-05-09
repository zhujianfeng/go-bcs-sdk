package bcs

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type BaiduBcs struct {
	Ak      string
	Sk      string
	Host    string
	PubHost string
}

//bcs构造函数
func NewBaiduBcs(ak, sk, host, pubHost string) *BaiduBcs {
	return &BaiduBcs{ak, sk, host, pubHost}
}

//生成签名
func (bcs *BaiduBcs) formatSign(opt map[string]string) string {
	flags := "MBO"
	content := ""
	if opt["Method"] != "" && opt["Bucket"] != "" && opt["Object"] != "" {
		content += "Method=" + opt["Method"] + "\n"
		content += "Bucket=" + opt["Bucket"] + "\n"
		content += "Object=" + opt["Object"] + "\n"
	} else {
		return ""
	}
	if opt["Ip"] != "" {
		flags += "I"
		content += "Ip=" + opt["Ip"] + "\n"
	}
	if opt["Time"] != "" {
		flags += "T"
		content += "Time=" + opt["Time"] + "\n"
	}
	if opt["Size"] != "" {
		flags += "Size=" + opt["Size"] + "\n"
	}
	content = flags + "\n" + content
	mac := hmac.New(sha1.New, []byte(bcs.Sk))
	mac.Write([]byte(content))
	sign := url.QueryEscape(base64.StdEncoding.EncodeToString(mac.Sum(nil)))

	return "sign=" + flags + ":" + bcs.Ak + ":" + sign
}

//生成URL
func (bcs *BaiduBcs) formatUrl(opt map[string]string) string {
	sign := bcs.formatSign(opt)
	if sign == "" {
		return ""
	}
	url := "http://" + bcs.Host + "/" + opt["Bucket"]
	if opt["Object"] != "/" {
		url += opt["Object"]
	}
	url += "?" + sign

	return url
}

//根据[]byte内容创建object
func (bcs *BaiduBcs) CreateObject(bucket, object string, body []byte) (int, map[string][]string, string) {
	opt := make(map[string]string)
	opt["Method"] = "PUT"
	opt["Bucket"] = bucket
	opt["Object"] = object

	url := bcs.formatUrl(opt)

	req, err := http.NewRequest("PUT", url, bytes.NewReader(body))
	if err != nil {
		return 0, nil, ""
	}

	bodyLen := len(body)
	req.Header.Set("Content-Length", strconv.Itoa(bodyLen))
	req.Header.Set("x-bs-acl", "public-read")

	contentType := ""
	fileNameArr := strings.Split(object, ".")
	arrLen := len(fileNameArr)
	if arrLen > 1 {
		suffix := fileNameArr[arrLen-1]
		contentType = mimeTypes[suffix]
	}
	if contentType == "" {
		contentType = "binary/octet-stream"
	}
	req.Header.Set("Content-Type", contentType)

	var client http.Client
	res, err := client.Do(req)
	if err != nil {
		return 0, nil, ""
	}

	pubObjUrl := ""
	if res.StatusCode == 200 {
		pubObjUrl = "http://" + bcs.PubHost + "/" + bucket + object
	}
	defer res.Body.Close()
	return res.StatusCode, res.Header, pubObjUrl
}

//根据文件路径创建object
func (bcs *BaiduBcs) CreateObjectByFile(bucket string, object string, path string) (int, map[string][]string, string) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, nil, ""
	}
	return bcs.CreateObject(bucket, object, body)
}

//根据文本内容创建object
func (bcs *BaiduBcs) CreateObjectByText(bucket string, object string, text string) (int, map[string][]string, string) {
	body := []byte(text)
	return bcs.CreateObject(bucket, object, body)
}

//下载object
func (bcs *BaiduBcs) GetObject(bucket string, object string) ([]byte, error) {
	opt := make(map[string]string)
	opt["Method"] = "GET"
	opt["Bucket"] = bucket
	opt["Object"] = object

	url := bcs.formatUrl(opt)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var client http.Client
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New("get object fail")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

//下载object并保存到文件
func (bcs *BaiduBcs) GetObjectAndSave(bucket, object, path string) error {
	body, err := bcs.GetObject(bucket, object)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, body, 0644)
	if err != nil {
		return err
	}
	return nil
}

//删除object
func (bcs *BaiduBcs) DeleteObject(bucket, object string) error {
	opt := make(map[string]string)
	opt["Method"] = "DELETE"
	opt["Bucket"] = bucket
	opt["Object"] = object

	url := bcs.formatUrl(opt)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	var client http.Client
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}
