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

func createRepoCmd(path string) pipe.Pipe {
	fmt.Println(path)
	cmd := []string{"git", "init", "--bare"}
	p := pipe.Script(pipe.MkdirAll(path, 0777), pipe.ChDir(path), pipe.Exec(cmd[0], cmd[1:]...))
	return p
}

func (gs *GitService) CreateRepo(repo string, user *User) (string, error) {
	fullPath := path.Join(gs.Path, user.Name, repo+".git")
	output, err := pipe.CombinedOutput(createRepoCmd(fullPath))
	return "file://" + fullPath, err
}
