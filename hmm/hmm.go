package hmm

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/3xian/darkseg/util"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

type Hmm struct {
	startProb []float64
	transProb [][]float64
	emitProb  []map[rune]float64
}

func NewHmm(filepath string) (hmm *Hmm, err error) {
	hmm = &Hmm{}
	if err := hmm.loadModel(filepath); err != nil {
		return nil, err
	}
	return hmm, nil
}

func (h *Hmm) Segment(textInRune []rune) (terms []string) {
	if len(textInRune) == 0 {
		return
	}

	weight := util.NewFloatMatrix(len(textInRune), h.stateSize())
	path := util.NewIntMatrix(len(textInRune), h.stateSize())

	for i, char := range textInRune {
		for j := 0; j < h.stateSize(); j++ {
			if i == 0 {
				weight[0][j] = h.startProb[j] + h.getEmitProb(textInRune[0], j)
			} else {
				weight[i][j] = -1e10
				for k := 0; k < h.stateSize(); k++ {
					if w := weight[i-1][k] + h.transProb[k][j] + h.getEmitProb(char, j); w > weight[i][j] {
						weight[i][j] = w
						path[i][j] = k
					}
				}
			}
		}
	}

	lastState := 0
	for j := 1; j < h.stateSize(); j++ {
		if weight[len(textInRune)-1][j] > weight[len(textInRune)-1][lastState] {
			lastState = j
		}
	}
	terms = h.getTermsFromPath(textInRune, path, lastState)
	return terms
}

func (h *Hmm) readStartProb(r *bufio.Reader) error {
	line, err := readLine(r)
	if err != nil {
		return err
	}

	for _, token := range strings.Split(line, " ") {
		var prob float64
		fmt.Sscan(token, &prob)
		h.startProb = append(h.startProb, prob)
	}

	log.Println("startProb:", h.startProb)
	return nil
}

func (h *Hmm) readTransProb(r *bufio.Reader) error {
	h.transProb = util.NewFloatMatrix(h.stateSize(), h.stateSize())

	for i := 0; i < h.stateSize(); i++ {
		line, err := readLine(r)
		if err != nil {
			return err
		}

		tokens := strings.Split(line, " ")
		if len(tokens) != h.stateSize() {
			return errors.New("bad trans prob matrix")
		}

		for j, token := range tokens {
			fmt.Sscan(token, &h.transProb[i][j])
		}
	}

	log.Println("transProb:", h.transProb)
	return nil
}

func (h *Hmm) readEmitProb(r *bufio.Reader) error {
	h.emitProb = make([]map[rune]float64, h.stateSize())
	for i := 0; i < h.stateSize(); i++ {
		h.emitProb[i] = make(map[rune]float64)
	}

	for i := 0; i < h.stateSize(); i++ {
		line, err := readLine(r)
		if err != nil {
			return err
		}

		tokens := strings.Split(line, ",")
		for _, token := range tokens {
			kv := strings.Split(token, ":")
			if len(kv) != 2 {
				return errors.New("bad emit prob matrix")
			}

			var (
				char rune
				prob float64
			)
			char, _ = utf8.DecodeRuneInString(kv[0])
			fmt.Sscan(kv[1], &prob)
			h.emitProb[i][char] = prob
		}
		log.Printf("emitProb[%d] size: %d", i, len(tokens))
	}
	return nil
}

func (h *Hmm) stateSize() int {
	return len(h.startProb)
}

func (h *Hmm) getTermsFromPath(textInRune []rune, path [][]int, lastState int) (terms []string) {
	state := lastState
	end := len(textInRune)
	for i := len(textInRune) - 1; i >= 0; i-- {
		if i == 0 || state == 0 || state == 3 {
			terms = append(terms, string(textInRune[i:end]))
			end = i
		}
		state = path[i][state]
	}

	i := 0
	j := len(terms) - 1
	for i < j {
		terms[i], terms[j] = terms[j], terms[i]
		i++
		j--
	}
	return
}

func (h *Hmm) loadModel(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	if err = h.readStartProb(reader); err != nil {
		return err
	}
	if err = h.readTransProb(reader); err != nil {
		return err
	}
	if err = h.readEmitProb(reader); err != nil {
		return err
	}
	return nil
}

func (h *Hmm) getEmitProb(char rune, state int) float64 {
	if prob, ok := h.emitProb[state][char]; ok {
		return prob
	} else {
		return 0.0
	}
}

func readLine(r *bufio.Reader) (line string, err error) {
	for {
		line, err = r.ReadString('\n')
		if err != nil {
			return
		}
		if len(line) > 0 && line[0] != '#' {
			return
		} else {
			continue
		}
	}
}
