package repository

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jayakrishnanMurali/kit/pkg/helper"
	"gopkg.in/ini.v1"
)

var ErrNotGitRepository = errors.New("not a git repository")
var ErrNoConfigFile = errors.New("no config file")
var ErrUnSupportedVersion = errors.New("unsupported repository version")
var ErrUnableToCreate = errors.New("unable to create repository")
var ErrNoRepo = errors.New("no repository found")

type GitRepository struct {
	worktree string
	Gitdir   string
	Conf     *ini.File
}

func NewGitRepository(path string, force bool) (*GitRepository, error) {

	dir := filepath.Join(path, ".git")

	if !force && helper.IsDir(dir) {
		return nil, ErrNotGitRepository
	}

	repo := &GitRepository{
		worktree: path,
		Gitdir:   dir,
	}

	configFile := filepath.Join(repo.Gitdir, "config")

	if helper.IsFile(configFile) {
		conf, err := ini.Load(configFile)
		if err != nil {
			return nil, ErrNoConfigFile
		}

		repo.Conf = conf
	} else if !force {
		return nil, ErrNoConfigFile
	}

	if !force {
		vers := repo.Version()

		if vers != "0" {
			return nil, ErrUnSupportedVersion
		}
	}

	return repo, nil

}

func (repo *GitRepository) Version() string {
	return repo.Conf.Section("core").Key("repositoryformatversion").String()
}

func RepoPath(repo *GitRepository, paths ...string) string {
	allPaths := append([]string{repo.Gitdir}, paths...)
	return filepath.Join(allPaths...)
}

func RepoFile(repo *GitRepository, mkdir bool, paths ...string) (string, error) {
	dirPath := RepoDir(repo, mkdir, paths[:len(paths)-1]...)
	if dirPath == "" {
		return "", os.ErrNotExist
	}
	return RepoPath(repo, paths...), nil
}

func RepoDir(repo *GitRepository, mkdir bool, paths ...string) string {
	path := RepoPath(repo, paths...)

	if helper.IsDir(path) {
		return path
	}

	if mkdir {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return ""
		}
		return path
	}

	return ""
}

func RepoCreate(path string) (*GitRepository, error) {
	repo, err := NewGitRepository(path, true)
	if err != nil {
		return nil, ErrUnableToCreate
	}

	if helper.IsExist(repo.worktree) {
		if !helper.IsDir(repo.worktree) {
			return nil, fmt.Errorf("%s is not a directory", path)
		}
		if helper.IsDir(repo.Gitdir) && !helper.IsEmpty(repo.Gitdir) {
			return nil, fmt.Errorf("%s is not empty", path)
		}
	} else {
		err := os.MkdirAll(repo.worktree, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	if RepoDir(repo, true, "branches") == "" {
		return nil, fmt.Errorf("failed to create branches directory")
	}
	if RepoDir(repo, true, "objects") == "" {
		return nil, fmt.Errorf("failed to create objects directory")
	}
	if RepoDir(repo, true, "refs", "tags") == "" {
		return nil, fmt.Errorf("failed to create refs/tags directory")
	}
	if RepoDir(repo, true, "refs", "heads") == "" {
		return nil, fmt.Errorf("failed to create refs/heads directory")
	}

	descriptionFile, err := RepoFile(repo, true, "description")
	if err != nil {
		return nil, err
	}
	err = os.WriteFile(descriptionFile, []byte("Unnamed repository; edit this file 'description' to name the repository.\n"), os.ModePerm)
	if err != nil {
		return nil, err
	}

	headFile, err := RepoFile(repo, true, "HEAD")
	if err != nil {
		return nil, err
	}
	err = os.WriteFile(headFile, []byte("ref: refs/heads/master\n"), os.ModePerm)
	if err != nil {
		return nil, err
	}

	configFile, err := RepoFile(repo, true, "config")
	if err != nil {
		return nil, err
	}
	config := helper.DefaultINIConfig()
	err = config.SaveTo(configFile)
	if err != nil {
		return nil, err
	}

	return repo, nil
}
