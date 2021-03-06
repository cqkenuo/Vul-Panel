package main

import (
	"crypto/md5"
	"fmt"
	"strings"
	"time"
)

func (s *Service) writeHTML(data VulInfo) {
	// 生成漏洞详情html文件名
	hostSlice := strings.Split(data.Detail.Host, ".")
	md5s := []byte(data.Detail.Url + data.Detail.Payload + time.Now().String())
	has := md5.Sum(md5s)
	ext := fmt.Sprintf("%x", has) //将[]byte转成16进制
	s.Conf.TempFileName = hostSlice[len(hostSlice)-2] + "-" + ext
	//contents, err := ioutil.ReadFile("./conf/template.html")
	//if err != nil {
	//	fmt.Println("本地文件读取失败")
	//}
	//
	//template := string(contents)
	//template = fmt.Sprintf(template,
	//	data.Detail.Url,
	//	data.VulClass,
	//	html.EscapeString(data.Detail.Request),
	//	html.EscapeString(data.Detail.Response),
	//)
	//
	//if err := ioutil.WriteFile(s.Conf.Base.WebRoot+s.Conf.TempFileName+".html", []byte(template), 0777); err != nil {
	//	log.Println(err.Error())
	//	return
	//}
	//fmt.Println(s.Conf.TempFileName + ".html" + "静态页面写入成功")
}
