package parse

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"sync"
)

func Parse(htmlUrl string) {

	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("即将请求:", r.URL)

	})
	c.OnResponse(func(response *colly.Response) {
		fmt.Println("请求来的html->", response.StatusCode)
		fmt.Println("请求来的html->", response.Request.URL.String())
		fmt.Println("请求来的html->", htmlUrl)
		if response.Request.URL.String() != htmlUrl {
			fmt.Println("没有找到目标地址，请求被转发了，转发地址-->", response.Request.URL.String())
			fmt.Println("退出解析网页")
			return
		}
		htmlDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(response.Body))
		if err != nil {
			fmt.Println("解析网页出问题了~~~", err)
			log.Fatal(err)
		}
		Location := response.Headers.Get("Location")
		fmt.Println("响应头Location:", Location)

		elment := htmlDoc.Find(".g-playicon img")
		var wg sync.WaitGroup
		elment.Each(func(i int, s *goquery.Selection) {
			srcUrl, isExist := s.Attr("src")
			alt, _ := s.Attr("alt")
			if isExist {
				fmt.Println("找到的图片地址:", srcUrl)
				fmt.Println("找到的图片文案:", alt)
				gifUrl, _ := FindGifUrl(srcUrl)
				picInfo := NewPicInfo(htmlUrl, srcUrl, gifUrl, alt)
				wg.Add(1) //wg
				go startDown(picInfo, &wg)
			}
		})
		wg.Wait()
	})

	c.Visit(htmlUrl)
}

func startDown(info picInfo, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("开始下载:",info.GifUrl)
	fileName := path.Base(info.GifUrl)
	// Get the data
	resp, err := http.Get(info.GifUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	 _,existErr := os.Lstat("./gif")
	if os.IsNotExist(existErr) {
		os.Mkdir("gif",os.ModePerm)
	}
	// 创建一个文件用于保存
	out, err := os.Create("./gif/" + fileName)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
}
