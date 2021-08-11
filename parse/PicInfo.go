package parse

import (
	"strings"
)

type picInfo struct {
	ParentUrl string
	JpgUrl    string
	GifUrl    string
	PicText   string
}

// FindGifUrl / 根据传入的jpg信息，构架gif 地址
func FindGifUrl(jpgUrl string) (string,bool){
	if len(jpgUrl)>0 {
		if strings.HasSuffix(jpgUrl,"_sub.jpg") {
			return strings.Replace(jpgUrl,"_sub.jpg",".gif",-1),true
		}
	}
	return jpgUrl ,false
}

// NewPicInfo 创建一个图片信息结构体
func NewPicInfo(parentUrl string,jpgUrl string,gifUrl string,picText string) picInfo  {
	return picInfo{ parentUrl,jpgUrl,gifUrl,picText}
}
