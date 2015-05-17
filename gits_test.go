package gits

import "testing"

func TestCreate(t *testing.T) {	
	gs := NewService(&Config{ServerPath: "/Users/mjha/git/data"})
	url, err := gs.CreateRepo("tempting", &User{"123", "maddy"})
	want := "file:///Users/mjha/git/data/maddy/tempting.git"
	if url != want || err != nil {
		t.Errorf("%v", err)
		t.Errorf("Create: %v, want %v", url, want)
	}
}