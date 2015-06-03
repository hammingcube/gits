package gits

import (
	"fmt"
	"gopkg.in/pipe.v2"
	_ "os"
	"path"
)

type User struct {
	Id   string
	Name string
}

type GitService struct {
	Path string
}

type Config struct {
	ServerPath string
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

func createRepoScript(path string) pipe.Pipe {
	p := pipe.Script(
		pipe.MkDirAll(path, 0777),
		pipe.ChDir(path),
		pipe.Exec("git", "init", "--bare"),
	)
	return p
}


func (gs *GitService) CreateRepo(repo string, user *User) (string, error) {
	fullPath := path.Join(gs.Path, user.Name, repo+".git")
	p := createRepoScript(fullPath)
	output, err := pipe.CombinedOutput(p)
	fmt.Println(string(output))
	return "file://" + fullPath, err
}

func (gs *GitService) AddRepo(repo string, user *User) (string, error) {
	fullPath := path.Join(gs.Path, user.Name, repo+".git")
	p := createRepoScript(fullPath)
	output, err := pipe.CombinedOutput(p)
	fmt.Println(string(output))
	return "file://" + fullPath, err
}
