package gits

import "testing"
import "gopkg.in/pipe.v2"
import "fmt"

func TestCreate(t *testing.T) {
	gs := NewService(&Config{ServerPath: "/Users/mjha/git/data"})
	url, err := gs.CreateRepo("tempting", &User{"123", "maddy"})
	want := "file:///Users/mjha/git/data/maddy/tempting.git"
	if url != want || err != nil {
		t.Errorf("%v", err)
		t.Errorf("Create: %v, want %v", url, want)
	}
}

func TestFiles(t *testing.T) {
	m := map[string][]byte{
		"abc.txt": []byte("Hello\n"),
		"pqr.txt": []byte("New stuff\n"),
	}
	p := addToRepoScript(m, "data")
	output, _ := pipe.CombinedOutput(p)
	fmt.Println(string(output))
}
