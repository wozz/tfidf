package tfidf

import "testing"

func TestSaveLoadEmpty(t *testing.T) {
	f := New()
	b, err := f.Save()
	if err != nil {
		t.Fatalf("could not save: %v", err)
	}
	f2, err := Load(b)
	if err != nil {
		t.Fatalf("could not load: %v", err)
	}
	if f.n != f2.n {
		t.Error("saved/loaded not same")
	}
}

func TestSaveLoadWithDocs(t *testing.T) {
	f := New()
	f.AddDocs(Document{"a", "b"}, Document{"b", "c"})
	b, err := f.Save()
	if err != nil {
		t.Fatalf("could not save: %v", err)
	}
	f2, err := Load(b)
	if err != nil {
		t.Fatalf("could not load: %v", err)
	}
	if f.n != f2.n {
		t.Error("saved/loaded not same")
	}
	cal1 := f.Cal(Document{"a"})
	cal2 := f2.Cal(Document{"a"})
	acal1, ok := cal1["a"]
	if !ok {
		t.Error("calculation incorrect, missing val")
	}
	acal2, ok := cal2["a"]
	if !ok {
		t.Error("calculation incorrect, missing val")
	}
	if acal1 != acal2 {
		t.Error("calculation incorrect, not equal")
	}
}
