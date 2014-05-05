package trie

import (
	"github.com/3xian/darkseg/util"
	"testing"
)

func TestLongestMatchedPrefix(t *testing.T) {
	trie := NewTrie()
	trie.Insert([]rune("吃饭"))
	util.Expect(t, string(trie.LongestMatchedPrefix([]rune("吃饭啦"))), "吃饭")
	util.Expect(t, string(trie.LongestMatchedPrefix([]rune("吃鸡蛋"))), "")
	util.Expect(t, string(trie.LongestMatchedPrefix([]rune("吃饭"))), "吃饭")
	util.Expect(t, string(trie.LongestMatchedPrefix([]rune("鸡蛋"))), "")

	trie.Insert([]rune("吃鸡蛋"))
	util.Expect(t, string(trie.LongestMatchedPrefix([]rune("吃饭啦"))), "吃饭")
	util.Expect(t, string(trie.LongestMatchedPrefix([]rune("吃鸡蛋"))), "吃鸡蛋")
	util.Expect(t, string(trie.LongestMatchedPrefix([]rune("吃"))), "")
	util.Expect(t, string(trie.LongestMatchedPrefix([]rune("鸡蛋"))), "")
}
