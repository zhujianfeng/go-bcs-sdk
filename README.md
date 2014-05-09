go-bcs-sdk
====================
go-bcs-sdk是使用go语言实现的一个百度云存储的sdk。目前百度云存储官方提供PHP、Java、python和C/C++的版本（详见[这里](http://developer.baidu.com/wiki/index.php?title=docs/cplat/bcs/sdk)），本sdk为非官方实现，功能暂时只提供以下几个：

* 创建object
* 删除object
* 下载object

bucket操作、其他object操作等功能暂未实现。

安装
----
	go get github.com/zhujianfeng/go-bcs-sdk
	
使用
----
### 引入
	import "github.com/zhujianfeng/go-bcs-sdk"
	
### 初始化
	ak := "your ak"
	sk := "your sk"
	host := "bcs server" //bcs的地址
	pubhost := "bcs server" //object上传后生成访问地址重的域名，一般和host相同，在区分内外网时不同
	baiduBCS := bcs.NewBaiduBcs(ak, sk, host, pubhost)
	
方法列表
------
创建object

	func (bcs *BaiduBcs) CreateObject(bucket, object string, body []byte) (int, map[string][]string, string)
	
* bucket：object所在的bucket，需要保证该bucket存在
* object: 待创建的object
* body: object的内容

* 返回值分别为本次上传的http状态码、http头信息和object地址


根据文件路径创建object

	func (bcs *BaiduBcs) CreateObjectByFile(bucket, object, path string) (int, map[string][]string, string)
	
* bucket：object所在的bucket，需要保证该bucket存在
* object: 待创建的object
* path: object在本地的路径

* 返回值分别为本次上传的http状态码、http头信息和object地址

根据文本内容创建object

	func (bcs *BaiduBcs) CreateObjectByText(bucket, object, text string) (int, map[string][]string, string)
	
* bucket：object所在的bucket，需要保证该bucket存在
* object: 待创建的object
* text: 文本内容

* 返回值分别为本次上传的http状态码、http头信息和object地址

下载object

	func (bcs *BaiduBcs) GetObject(bucket string, object string) ([]byte, error)
	
* bucket：object所在的bucket，需要保证该bucket存在
* object: 待下载的object

* 返回值分别为object内容和错误信息

下载object并保存到文件

	func (bcs *BaiduBcs) GetObjectAndSave(bucket, object, path string) error
	
* bucket：object所在的bucket，需要保证该bucket存在
* object: 待下载的object
* path：object存到本地的文件路径

* 返回值：错误信息

删除object

	func (bcs *BaiduBcs) DeleteObject(bucket, object string) error 
	
* bucket：object所在的bucket，需要保证该bucket存在
* object: 待删除的object

* 返回值：错误信息

License
-------

使用 [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html).
