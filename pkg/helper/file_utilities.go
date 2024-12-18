package helper

import (
	"os"

	"gopkg.in/ini.v1"
)

func IsDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func IsFile(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func IsExist(path string) bool {
	_, err := os.Stat(path + "/")
	return err == nil
}

func IsEmpty(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	return err != nil
}

func DefaultINIConfig() *ini.File {
	cfg := ini.Empty()
	cfg.Section("core").Key("repositoryformatversion").SetValue("0")
	cfg.Section("core").Key("filemode").SetValue("false")
	cfg.Section("core").Key("bare").SetValue("false")
	return cfg
}
