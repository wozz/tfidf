package tfidf

import (
	"hash/fnv"
	"math"
)

const sep = " "

// TFIDF holds info to calculate the TFIDF scores
type TFIDF struct {
	docIndex  map[uint64]int
	termFreqs []map[string]int
	termDocs  map[string]int
	n         int
}

// New creates a new empty TFIDF
func New() *TFIDF {
	return &TFIDF{
		docIndex:  make(map[uint64]int),
		termFreqs: make([]map[string]int, 0),
		termDocs:  make(map[string]int),
		n:         0,
	}
}

// AddDocs updates the TFIDF struct with the doc info. each doc should already be
// tokenized and represented as a []string
func (f *TFIDF) AddDocs(docs ...[]string) {
	for _, doc := range docs {
		h := hash(doc)
		if f.docHashPos(h) >= 0 {
			continue
		}

		termFreq := f.termFreq(doc)
		if len(termFreq) == 0 {
			continue
		}

		f.docIndex[h] = f.n
		f.n++

		f.termFreqs = append(f.termFreqs, termFreq)

		for term := range termFreq {
			f.termDocs[term]++
		}
	}
}

// Cal calculates tf-idf weight for specified document
func (f *TFIDF) Cal(doc []string) (weight map[string]float64) {
	weight = make(map[string]float64)

	var termFreq map[string]int

	docPos := f.docPos(doc)
	if docPos < 0 {
		termFreq = f.termFreq(doc)
	} else {
		termFreq = f.termFreqs[docPos]
	}

	docTerms := 0
	for _, freq := range termFreq {
		docTerms += freq
	}
	for term, freq := range termFreq {
		weight[term] = tfidf_(freq, docTerms, f.termDocs[term], f.n)
	}

	return weight
}

func (f *TFIDF) termFreq(doc []string) (m map[string]int) {
	m = make(map[string]int)

	for _, term := range doc {
		m[term]++
	}

	return
}

func (f *TFIDF) docHashPos(hash uint64) int {
	if pos, ok := f.docIndex[hash]; ok {
		return pos
	}

	return -1
}

func (f *TFIDF) docPos(doc []string) int {
	return f.docHashPos(hash(doc))
}

func hash(doc []string) uint64 {
	h := fnv.New64a()
	for _, term := range doc {
		h.Write([]byte(term))
		h.Write([]byte(sep))
	}
	return h.Sum64()
}

func tfidf_(termFreq, docTerms, termDocs, N int) float64 {
	tf := float64(termFreq) / float64(docTerms)
	idf := math.Log(float64(1+N) / (1 + float64(termDocs)))
	return tf * idf
}
