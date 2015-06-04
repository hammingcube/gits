package gits

import "testing"
import _ "gopkg.in/pipe.v2"
import "fmt"

const REPO = "file:///Users/mjha/git/data/maddy/tempting.git"
const DEST = "cool"

func _TestCreate(t *testing.T) {
	gs := NewService(&Config{ServerPath: "/Users/mjha/git/data"})
	url, err := gs.CreateRepo("tempting", &User{"123", "maddy"})
	want := "file:///Users/mjha/git/data/maddy/tempting.git"
	if url != want || err != nil {
		t.Errorf("%v", err)
		t.Errorf("Create: %v, want %v", url, want)
	}
}

func _Test1(t *testing.T) {
	gs := NewService(&Config{ServerPath: "/Users/mjha/git/data"})
	err := gs.PrepareRepo(REPO, DEST)
	fmt.Println(err)
}

func TestFiles(t *testing.T) {
	gs := NewService(&Config{ServerPath: "/Users/mjha/git/data"})
	files := map[string][]byte{
		"abc.txt": []byte("Hello\nHow do you do?\n"),
		"pqr.txt": []byte("New stuff\n"),
		"uvw.txt": []byte("Totally new file\n"),
	}
	gs.AddToRepo("tempting", &User{"123", "maddy"}, files)
}
