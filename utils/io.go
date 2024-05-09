package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func ExpandUser(path string) (string, error) {
	if !strings.HasPrefix(path, "~") {
		return path, nil
	}
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path = dir + path[1:]
	return path, nil
}

func IsFile(name string) (bool, error) {
	info, err := os.Stat(name)
	if err == nil {
		return !info.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func IsDir(name string) (bool, error) {
	info, err := os.Stat(name)
	if err == nil {
		return info.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CopyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0o644)
}

func MakeParentDir(filePath string) error {
	dir := filepath.Dir(filePath)
	return os.MkdirAll(dir, 0o755)
}

func SafeCopyFile(src, dst string) error {
	src, err := ExpandUser(src)
	if err != nil {
		return err
	}
	err = MakeParentDir(dst)
	if err != nil {
		return err
	}
	return CopyFile(src, dst)
}

func SafeCopyDir(src, dst string) error {
	src, err := ExpandUser(src)
	if err != nil {
		return err
	}
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = SafeCopyDir(srcPath, dstPath)
			if err != nil {
				return err
			}
		} else {
			err = SafeCopyFile(srcPath, dstPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
