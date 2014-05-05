package darkseg

import (
	"github.com/3xian/darkseg/util"
	"testing"
)

func TestSegmentEmptyText(t *testing.T) {
	segmentor, err := NewSegmentor("data/20140502.dict", "data/hmm.model")
	util.Expect(t, err, nil)

	text := ""
	util.Expect(t, segmentor.Segment(text), []string{})
}

func TestSegmentNormalText(t *testing.T) {
	segmentor, err := NewSegmentor("data/20140502.dict", "data/hmm.model")
	util.Expect(t, err, nil)
	segmentor.SetDebug(true)

	text := "刷牙洗澡睡觉"
	util.Expect(t, FormatTerms(segmentor.Segment(text)), "刷牙|洗澡|睡觉")
}

func TestSegmentAlphabets(t *testing.T) {
	segmentor, err := NewSegmentor("data/20140502.dict", "data/hmm.model")
	util.Expect(t, err, nil)
	segmentor.SetDebug(true)

	text1 := "Hello world!"
	util.Expect(t, FormatTerms(segmentor.Segment(text1)), "Hello| |world|!")

	text2 := "After 10 years"
	util.Expect(t, FormatTerms(segmentor.Segment(text2)), "After| |10| |years")

	text3 := "23天"
	util.Expect(t, FormatTerms(segmentor.Segment(text3)), "23|天")
}
