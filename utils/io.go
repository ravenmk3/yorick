package utils

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func ExpandUser(path string) (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path = strings.ReplaceAll(path, "~", dir)
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

func SafeCopyDirEx(src, dst string, excludes []string) error {
	files, err := ListFiles(src, true, -1)
	if err != nil {
		return err
	}

	if excludes != nil && len(excludes) > 0 {
		files, err = filterExcludedFiles(files, excludes)
		if err != nil {
			return err
		}
	}

	for _, file := range files {
		srcFile := filepath.Join(src, file)
		dstFile := filepath.Join(dst, file)
		err := SafeCopyFile(srcFile, dstFile)
		if err != nil {
			return err
		}
	}

	return nil
}

func filterExcludedFiles(files, patterns []string) ([]string, error) {
	result := []string{}
	regexps := []*regexp.Regexp{}
	for _, pattern := range patterns {
		expr, err := regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}
		regexps = append(regexps, expr)
	}

	for _, file := range files {
		excluded := isFileExcluded(file, regexps)
		if excluded {
			continue
		}
		result = append(result, file)
	}
	return result, nil
}

func isFileExcluded(file string, regexps []*regexp.Regexp) bool {
	for _, re := range regexps {
		matched := re.MatchString(file)
		if matched {
			return true
		}
	}
	return false
}
