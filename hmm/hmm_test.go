package hmm

import (
	"github.com/3xian/darkseg/util"
	"testing"
)

func TestLoadModel(t *testing.T) {
	hmm, err := NewHmm("../data/hmm.model")
	util.Expect(t, err, nil)
	util.Expect(t, hmm.stateSize(), 4)
}
