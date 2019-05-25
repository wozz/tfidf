package tfidf

import (
	"bytes"
	"encoding/gob"
)

// Save exports the TFIDF struct for persisting to disk
func (f *TFIDF) Save() ([]byte, error) {
	var e bytes.Buffer
	enc := gob.NewEncoder(&e)
	t := struct {
		D  map[uint64]int
		TF []map[string]int
		TD map[string]int
		N  int
	}{f.docIndex, f.termFreqs, f.termDocs, f.n}
	err := enc.Encode(t)
	return e.Bytes(), err
}

// Load imports a previously saved TFIDF struct
func Load(data []byte) (*TFIDF, error) {
	t := struct {
		D  map[uint64]int
		TF []map[string]int
		TD map[string]int
		N  int
	}{}
	d := bytes.NewBuffer(data)
	dec := gob.NewDecoder(d)
	err := dec.Decode(&t)
	if err != nil {
		return nil, err
	}
	return &TFIDF{
		docIndex:  t.D,
		termFreqs: t.TF,
		termDocs:  t.TD,
		n:         t.N,
	}, nil
}
