package main

import (
	"fmt"
	"github.com/3xian/darkseg"
)

var (
	segmentor *darkseg.Segmentor
)

func init() {
	var err error
	segmentor, err = darkseg.NewSegmentor("../data/20140502.dict", "../data/hmm.model")
	if err != nil {
		panic(err)
	}
	segmentor.SetDebug(true)
}

func segment(text string) {
	terms := segmentor.Segment(text)
	fmt.Printf("\033[35m%s => %v\033[0m\n", text, darkseg.FormatTerms(terms))
}

func main() {
	segment("刷牙洗澡睡觉")
	segment("我在1号店买了一台CALL机")
	segment("我来到北京清华大学")
	segment("他来到了网易杭研大厦")
	segment("小明硕士毕业于中国科学院计算所，后在日本京都大学深造")
	segment("龟头疼怎么办")
	segment("百度91博弈全过程：将成中国互联网最大并购")
	segment("来源：21世纪经济报道")
	segment("7月16日上午，百度宣布拟全资收购在香港上市的网龙公司旗下的91无线业务，购买总价为19亿美元，相关各方已经签署谅解备忘录。91无线方面随后向记者证实了此事。该交易如果完成，将成为中国互联网历史上最大的一笔并购案。")
	segment("南京市长江大桥")
	segment("浙江省了大批投资")
	segment("这事儿的确定不下来")
	segment("发展中国家服装需求大增")
	segment("我们提供高档设备和服务。")
	segment("我是个性开放的人")
	segment("展开我党先进性教育工作")
	segment("How could this happen to me 啊？")
}
