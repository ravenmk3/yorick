package utils

import (
	"os"
	"path/filepath"
)

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
	err := MakeParentDir(dst)
	if err != nil {
		return err
	}
	return CopyFile(src, dst)
}

func SafeCopyDir(src, dst string) error {
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
