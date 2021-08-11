package main

import (
	"Spider/config"
	"Spider/parse"
	"fmt"
)

func main() {

	//ch := make(chan string, 2)
	////go read(ch)
	//go write(ch)
	//d1:=<-ch
	//d2:=<-ch
	//d3:=<-ch
	//d4:=<-ch
	//fmt.Println("读取到数据:",d1,d2,d3,d4)
	////time.Sleep(2000)
	////for v := range ch {
	////	fmt.Println("read value", v,"from ch")
	////	//time.Sleep(2 * time.Second)
	////
	////}
	////ch <- "naveen1"
	////ch <- "naveen2"
	////ch <- "naveen3"
	////ch <- "naveen4"
	////ch <- "paul5"
	//fmt.Println("capacity is", cap(ch))
	//fmt.Println("length is", len(ch))
	//
	//c1 := make(chan string,1)
	//c1 <- "hello"
	////c1 <- "hello"


	num := 0
	config.InitUrl()
	fmt.Println(" 请选择下载的内容:  ")
	fmt.Println("===============Spider===============")
	fmt.Println("|            1、 全部                |")
	fmt.Println("|            2、 囧图                |")
	fmt.Println("|            3、 校园                |")
	fmt.Println("|            4、 明星                |")
	fmt.Println("|            5、 美女                |")
	fmt.Println("|            6、 亮点                |")
	fmt.Println("|            7、 雷人                |")
	fmt.Println("|            8、 内涵                |")
	fmt.Println("|            9、熊孩子               |")
	fmt.Println("|            10、 碉堡               |")
	fmt.Println("|            11、 动物               |")
	fmt.Println("|            12、 动漫               |")
	fmt.Println("|            13、 错觉               |")
	fmt.Println("|            14、 创意               |")
	fmt.Println("====================================")
	fmt.Scanln(&num)
	parse.Parse(config.GetUrl(num, 2))
	fmt.Println("下载完成，输入任意字符回车退出程序")
	fmt.Scanln(&num)

}


func write(ch chan string) {
	ch<-"hello_1"
	ch<-"hello_2"
	ch<-"hello_3"
	ch<-"hello_4"
	ch<-"hello_5"
	ch<-"hello_6"
	//ch<-"hello_5"
}

func read(ch chan string) {
	data1:=<-ch
	data2:=<-ch
	data3:=<-ch
	fmt.Println("读取到的信道数据-->",data1,data2,data3)
}
