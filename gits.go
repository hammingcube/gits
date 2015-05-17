package gits

import (
	"path"
	"os"
	"fmt"
	"gopkg.in/pipe.v2"
)

type User struct {
	Id		string
	Name	string
}

type GitService struct {
	Path	string
}

type Config struct {
	ServerPath	string
}

const SERVER_PATH = "/git/data/"

func NewService(cfg *Config) *GitService {
	gs := &GitService{}
	if cfg == nil {
		gs.Path = SERVER_PATH
	} else {
		gs.Path = cfg.ServerPath
	}
	return gs
}

func createRepo(path string) pipe.Pipe {
	cmd := []string{"git", "init", "--bare"}
	p := pipe.Line(pipe.ChDir(path), pipe.Exec(cmd[0], cmd[1:]...))
	return p
}

func (gs *GitService) CreateRepo(repo string, user *User) (string, error) {
	fullPath := path.Join(gs.Path, user.Name, repo+".git")
	fmt.Println(fullPath)
	err := os.MkdirAll(fullPath, 0777)
	p := createRepo(fullPath)
	output, err := pipe.CombinedOutput(p)
	fmt.Println(string(output))
	return "file://" + fullPath, err
}