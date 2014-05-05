package darkseg

import (
	"bufio"
	"github.com/3xian/darkseg/hmm"
	"github.com/3xian/darkseg/trie"
	"github.com/3xian/darkseg/util"
	"log"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Segmentor struct {
	hmm          *hmm.Hmm
	forwardTrie  *trie.Trie
	backwardTrie *trie.Trie
	debugFlag    bool
}

func NewSegmentor(dictPath string, hmmModelPath string) (segmentor *Segmentor, err error) {
	segmentor = &Segmentor{}

	segmentor.forwardTrie = trie.NewTrie()
	segmentor.backwardTrie = trie.NewTrie()

	dictFile, err := os.Open(dictPath)
	if err != nil {
		return nil, err
	}
	defer dictFile.Close()

	dictWordCounter := 0
	for scanner := bufio.NewScanner(dictFile); scanner.Scan(); {
		word := []rune(scanner.Text())
		segmentor.forwardTrie.Insert(word)
		segmentor.backwardTrie.Insert(util.ReverseRuneSlice(word))
		dictWordCounter++
	}
	log.Println("dictWordCount:", dictWordCounter)

	if segmentor.hmm, err = hmm.NewHmm(hmmModelPath); err != nil {
		return nil, err
	}
	return segmentor, nil
}

func (s *Segmentor) SetDebug(flag bool) {
	s.debugFlag = flag
}

func (s *Segmentor) Segment(text string) (terms []string) {
	if s.debugFlag {
		log.Println("text:", text)
	}

	textInRune := []rune(text)

	forward := s.maxMatch(textInRune)
	if s.debugFlag {
		log.Println("forward:", FormatTerms(offsetsToTerms(textInRune, forward)))
	}

	backward := s.reversedMaxMatch(textInRune)
	if s.debugFlag {
		log.Println("backward:", FormatTerms(offsetsToTerms(textInRune, backward)))
	}

	if util.IsIntSliceSame(forward, backward) {
		return offsetsToTerms(textInRune, forward)
	}

	return s.combineSegmentResults(textInRune, forward, backward)
}

func FormatTerms(terms []string) (str string) {
	return strings.Join(terms, "|")
}

func (s *Segmentor) combineSegmentResults(textInRune []rune, forward []int, backward []int) (terms []string) {
	var (
		subStart, subEnd int
		fi, fj, bi, bj   int
	)

	termEnd := func(offsets []int, index int) int {
		if index+1 >= len(offsets) {
			return len(textInRune)
		}
		return offsets[index+1]
	}

	for fi < len(forward) {
		fj = fi
		bj = bi
		for {
			forwardTermEnd := termEnd(forward, fj)
			for termEnd(backward, bj) < forwardTermEnd {
				bj++
			}
			if forwardTermEnd == termEnd(backward, bj) {
				break
			}
			fj++
		}
		fj++
		bj++

		subStart = forward[fi]
		if fj >= len(forward) {
			subEnd = len(textInRune)
		} else {
			subEnd = forward[fj]
		}
		if fi+1 == fj && bi+1 == bj {
			terms = append(terms, string(textInRune[subStart:subEnd]))
		} else {
			hmmResult := s.hmm.Segment(textInRune[subStart:subEnd])
			if s.debugFlag {
				log.Println("sub hmm:", FormatTerms(hmmResult))
			}
			terms = append(terms, hmmResult...)
		}

		fi = fj
		bi = bj
	}
	return
}

func (s *Segmentor) maxMatch(textInRune []rune) (offsets []int) {
	for i := 0; i < len(textInRune); {
		offsets = append(offsets, i)
		matchedWord := s.forwardTrie.LongestMatchedPrefix(textInRune[i:])
		if len(matchedWord) > 0 {
			i += len(matchedWord)
		} else {
			matchedWord = alphabetsAndDigitsPrefix(textInRune[i:])
			if len(matchedWord) > 0 {
				i += len(matchedWord)
			} else {
				i++
			}
		}
	}
	return
}

func (s *Segmentor) reversedMaxMatch(textInRune []rune) (offsets []int) {
	textInRune = util.ReverseRuneSlice(textInRune)
	for i := 0; i < len(textInRune); {
		matchedWord := s.backwardTrie.LongestMatchedPrefix(textInRune[i:])
		if len(matchedWord) > 0 {
			i += len(matchedWord)
		} else {
			matchedWord = alphabetsAndDigitsPrefix(textInRune[i:])
			if len(matchedWord) > 0 {
				i += len(matchedWord)
			} else {
				i++
			}
		}
		offsets = append(offsets, len(textInRune)-i)
	}
	return util.ReverseIntSlice(offsets)
}

func offsetsToTerms(textInRune []rune, offsets []int) (terms []string) {
	for i := 0; i < len(offsets); i++ {
		if i+1 < len(offsets) {
			terms = append(terms, string(textInRune[offsets[i]:offsets[i+1]]))
		} else {
			terms = append(terms, string(textInRune[offsets[i]:]))
		}
	}
	return
}

func alphabetsAndDigitsPrefix(textInRune []rune) (prefix []rune) {
	for _, r := range textInRune {
		if utf8.RuneLen(r) <= 2 && (unicode.IsLetter(r) || unicode.IsDigit(r)) {
			prefix = append(prefix, r)
		} else {
			break
		}
	}
	return
}
