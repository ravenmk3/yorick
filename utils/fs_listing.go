package utils

import (
	"os"
	"path/filepath"
	"time"
)

func ListDirs(dir string, relative bool, maxDepth int) ([]string, error) {
	dir, err := ExpandUser(dir)
	if err != nil {
		return nil, err
	}
	return listDirs(dir, dir, relative, maxDepth, 1)
}

func listDirs(baseDir, dir string, relative bool, maxDepth, depth int) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	dirs := []string{}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		entryPath := filepath.Join(dir, entry.Name())
		dirPath := entryPath
		if relative {
			dirPath, err = filepath.Rel(baseDir, dirPath)
			if err != nil {
				return nil, err
			}
		}
		dirs = append(dirs, dirPath)

		if maxDepth > 0 && depth >= maxDepth {
			continue
		}

		subDirs, err := listDirs(baseDir, entryPath, relative, maxDepth, depth+1)
		if err != nil {
			return nil, err
		}

		for _, subDir := range subDirs {
			dirs = append(dirs, subDir)
		}
	}
	return dirs, nil
}

func ListFiles(dir string, relative bool, maxDepth int) ([]string, error) {
	dir, err := ExpandUser(dir)
	if err != nil {
		return nil, err
	}
	return listFiles(dir, dir, relative, maxDepth, 1)
}

func listFiles(baseDir, dir string, relative bool, maxDepth, depth int) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	files := []string{}

	for _, entry := range entries {
		entryPath := filepath.Join(dir, entry.Name())

		if entry.IsDir() {
			if maxDepth > 0 && depth >= maxDepth {
				continue
			}
			subFiles, err := listFiles(baseDir, entryPath, relative, maxDepth, depth+1)
			if err != nil {
				return nil, err
			}
			for _, file := range subFiles {
				files = append(files, file)
			}
		} else {
			if relative {
				entryPath, err = filepath.Rel(baseDir, entryPath)
				if err != nil {
					return nil, err
				}
			}
			files = append(files, entryPath)
		}
	}
	return files, nil
}

func FindLatestFile(dir string, relative bool, maxDepth int) (string, error) {
	files, err := ListFiles(dir, relative, maxDepth)
	if err != nil {
		return "", err
	}

	latestFile := ""
	latestTime := time.Time{}

	for _, file := range files {
		fullPath := file
		if relative {
			fullPath = filepath.Join(dir, file)
		}

		info, err := os.Stat(fullPath)
		if err != nil {
			return "", err
		}

		if info.ModTime().After(latestTime) {
			latestTime = info.ModTime()
			latestFile = file
		}
	}

	return latestFile, nil
}
