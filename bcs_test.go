package bcs

import (
	//"fmt"
	"testing"
)

const AK string = "G1HvMoOfTqhrEpmdGM5GXC"
const SK string = "TJztiP1YXmlB1GVgnCkSpDZ39T9"
const HOST string = "bcs-sandbox.baidu.com"
const PUBHOST string = "bcs-sandbox.baidu.com"

/*
func TestFormatSign(t *testing.T) {
	bcs := NewBaiduBcs(AK, SK, HOST, PUBHOST)
	opt := make(map[string]string)
	opt["Method"] = "GET"
	opt["Bucket"] = "uaq-hijack-files"
	opt["Object"] = "/1.jpg"
	//opt["Ip"] = "220.234.21.2"
	sign := bcs.formatSign(opt)
	fmt.Println(sign)
}

func TestFormatUrl(t *testing.T) {
	bcs := NewBaiduBcs(AK, SK, HOST, PUBHOST)
	opt := make(map[string]string)
	opt["Method"] = "GET"
	opt["Bucket"] = "uaq-hijack-files"
	opt["Object"] = "/1.jpg"

	url := bcs.formatUrl(opt)
	fmt.Println(url)
}

func TestCreateObjectByFile(t *testing.T) {
	bcs := NewBaiduBcs(AK, SK, HOST, PUBHOST)
	opt := make(map[string]string)
	opt["Method"] = "GET"
	opt["Bucket"] = "uaq-hijack-files"
	opt["Object"] = "/4.jpg"

	_, _, url := bcs.CreateObjectByFile("uaq-hijack-files", "/11.jpg", "/Users/zhujianfeng/Desktop/1.jpg")
	fmt.Println(url)
}

func TestCreateObjectByText(t *testing.T) {
	bcs := NewBaiduBcs(AK, SK, HOST, PUBHOST)
	opt := make(map[string]string)
	opt["Method"] = "GET"
	opt["Bucket"] = "uaq-hijack-files"
	opt["Object"] = "/4.jpg"

	_, _, url := bcs.CreateObjectByText("uaq-hijack-files", "/11.json", "/Users/zhujianfeng/Desktop/1.jpg")
	fmt.Println(url)
}
*/
func TestGetObjectAndSave(t *testing.T) {
	bcs := NewBaiduBcs(AK, SK, HOST, PUBHOST)
	opt := make(map[string]string)
	opt["Method"] = "GET"
	opt["Bucket"] = "uaq-hijack-files"
	opt["Object"] = "/1.jpg"

	bcs.GetObjectAndSave("uaq-hijack-files", "/1.jpg", "/Users/zhujianfeng/Desktop/1.jpg")

}

func TestDeleteObject(t *testing.T) {
	bcs := NewBaiduBcs(AK, SK, HOST, PUBHOST)
	opt := make(map[string]string)
	opt["Method"] = "DELETE"
	opt["Bucket"] = "uaq-hijack-files"
	opt["Object"] = "/1.jpg"

	bcs.DeleteObject("uaq-hijack-files", "/1.jpg")
}
