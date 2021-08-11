package config

import "fmt"

// ALL_URL 全部
const all = "http://www.giftu.net/?currentPage=%d"
const jiongtu = "http://www.giftu.net/gif/jiongtu?currentPage=%d"
const xiaoyuan = "http://www.giftu.net/gif/xiaoyuan?currentPage=%d"
const mingxin = "http://www.giftu.net/gif/mingxin?currentPage=%d"
const meinv = "http://www.giftu.net/gif/meinv?currentPage=%d"
const liangdian = "http://www.giftu.net/gif/liangdian?currentPage=%d"
const leiren = "http://www.giftu.net/gif/leiren?currentPage=%d"
const neihan = "http://www.giftu.net/gif/neihan?currentPage=%d"
const xiaohai = "http://www.giftu.net/gif/xiaohai?currentPage=%d"
const diaobao = "http://www.giftu.net/gif/diaobao?currentPage=%d"
const danteng = "http://www.giftu.net/gif/danteng?currentPage=%d"
const dongwu = "http://www.giftu.net/gif/dongwu?currentPage=%d"
const dongman = "http://www.giftu.net/gif/dongman?currentPage=%d"
const cuojue = "http://www.giftu.net/gif/cuojue?currentPage=%d"
const chuangyi = "http://www.giftu.net/gif/jiongtu?chuangyi=%d"

var curMap = make(map[int]string)

func InitUrl()  {
	curMap[1] = all
	curMap[2] = jiongtu
	curMap[3] = xiaoyuan
	curMap[4] = mingxin
	curMap[5] = meinv
	curMap[6] = liangdian
	curMap[7] = leiren
	curMap[8] = neihan
	curMap[9] = xiaohai
	curMap[10] = diaobao
	curMap[11] = danteng
	curMap[12] = dongwu
	curMap[12] = dongman
	curMap[12] = cuojue
	curMap[12] = chuangyi
}

func GetUrl(t int,page int) string {
	url := curMap[t]
	return fmt.Sprintf(url,page)
}