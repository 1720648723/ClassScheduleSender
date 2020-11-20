package main

import (

	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){

	resp,err:=http.Get("http://hubs.hust.edu.cn/hub.jsp;jsessionid=jefj_Q8pw0Z1WU4aIDpsWMizKPjVcb77uNKJcC-b5AzfoilfQlPb!788946090")
	//异常时进行panic
	if err!=nil{
		panic(err)
	}
	//延时关闭
	defer resp.Body.Close()

	//状态码错误时输出错误码
	if resp.StatusCode!=http.StatusOK{
		fmt.Printf("Error ststus code:%d",resp.StatusCode)
	}

	//借助io.reader获取流的信息
	result,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s",result)
	//href<我理解为链接 有这个标签就会有子层的目录 比方说http://www.shouxindehai.cn/index.php/cs/>
	//在这里可以在终端搜索这个目录发现可以找到 说明是可以爬到的
	//该段代码适用于utf-8的编码的网址
	//gbk会乱码的
	//Sendmail(result)

}