package gits

import (
	"bytes"
	"fmt"
	"gopkg.in/pipe.v2"
	"os"
	"path"
	"strings"
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

func writeFilesScript(files map[string][]byte, dest string) pipe.Pipe {
	pipes := []pipe.Pipe{pipe.ChDir(dest)}
	for filename, blob := range files {
		p := pipe.Line(
			pipe.Read(bytes.NewReader(blob)),
			pipe.WriteFile(filename, 0644),
		)
		pipes = append(pipes, p)
	}
	p := pipe.Script(pipes...)
	return p
}

func cloneRepoScript(repo, dest string) pipe.Pipe {
	p := pipe.Exec("git", "clone", repo, dest)
	return p
}

func commitAllScript(dest string) pipe.Pipe {
	return pipe.Script(
		pipe.ChDir(dest),
		pipe.Exec("git", "add", "."),
		pipe.Exec("git", "commit", "-m", "More updates"),
		pipe.Exec("git", "push", "-u", "origin", "master"),
	)
}

func addToRepoScript(files map[string][]byte, repo string, dest string) pipe.Pipe {
	p := pipe.Script(
		writeFilesScript(files, dest),
		commitAllScript(dest),
	)
	return p
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func checkStatus(dest string) (bool, error) {
	p := pipe.Script(
		pipe.ChDir(dest),
		pipe.Exec("git", "status"),
	)
	output, err := pipe.CombinedOutput(p)
	fmt.Println(string(output))
	if err != nil {
		return false, err
	}
	if strings.Contains(string(output), "Your branch is up-to-date") {
		return true, nil
	} else {
		return false, nil
	}
}

func (gs *GitService) PrepareRepo(repo string, dest string) error {
	ok, err := exists(dest)
	if err != nil {
		return err
	}
	if ok {
		if good, err := checkStatus(dest); err != nil {
			return err
		} else if good {
			return nil
		}
	}
	err = os.RemoveAll(dest)
	output, err := pipe.CombinedOutput(cloneRepoScript(repo, dest))
	fmt.Println(string(output))
	return err
}

func (gs *GitService) AddToRepo(repo string, user *User, files map[string][]byte) (string, error) {
	dest := "coolnew"
	repoURL := "file://" + gs.RepoFullPath(repo, user)
	gs.PrepareRepo(repoURL, dest)
	p := addToRepoScript(files, repoURL, dest)
	output, err := pipe.CombinedOutput(p)
	fmt.Println(string(output))
	return repoURL, err
}

func (gs *GitService) RepoFullPath(repo string, user *User) string {
	return path.Join(gs.Path, user.Name, repo+".git")
}

func (gs *GitService) CreateRepo(repo string, user *User) (string, error) {
	fullPath := gs.RepoFullPath(repo, user)
	p := createRepoScript(fullPath)
	output, err := pipe.CombinedOutput(p)
	fmt.Println(string(output))
	return "file://" + fullPath, err
}
