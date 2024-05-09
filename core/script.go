package core

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"yorick/utils"
)

type ScriptObject struct {
	logger     *logrus.Logger
	outputDir  string
	currentDir string
	taskBegun  bool
	taskName   string
}

func NewScriptObject(outputDir string) *ScriptObject {
	return &ScriptObject{
		logger:     logrus.StandardLogger(),
		outputDir:  outputDir,
		currentDir: outputDir,
	}
}

func (s *ScriptObject) Format(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}

func (s *ScriptObject) LogInfo(format string, args ...any) {
	logrus.Infof(format, args...)
}

func (s *ScriptObject) LogWarn(format string, args ...any) {
	logrus.Warnf(format, args...)
}

func (s *ScriptObject) LogError(format string, args ...any) {
	logrus.Errorf(format, args...)
}

func (s *ScriptObject) GetEnv(key string) string {
	return os.Getenv(key)
}

func (s *ScriptObject) TaskBegin(name string, dir string) {
	s.taskBegun = true
	s.taskName = name
	s.logger.Infof("Task begin: %s", s.taskName)
	s.DestDir(dir)
}

func (s *ScriptObject) TaskEnd() {
	s.taskBegun = false
	s.taskName = ""
}

func (s *ScriptObject) DestDir(dir string) {
	s.currentDir = filepath.Join(s.outputDir, dir)
	s.logger.Infof("Destination directory: %s", s.currentDir)
}

func (s *ScriptObject) IsDir(name string) bool {
	isDir, err := utils.IsDir(name)
	if err != nil {
		s.logger.Warnf("IsDir: %s", err.Error())
		return false
	}
	return isDir
}

func (s *ScriptObject) IsFile(name string) bool {
	isFile, err := utils.IsFile(name)
	if err != nil {
		s.logger.Warnf("IsFile: %s", err.Error())
		return false
	}
	return isFile
}

func (s *ScriptObject) CopyFile(srcFile, dstFile string) {
	s.logger.Infof("CopyFile: %s => %s", srcFile, dstFile)

	srcFile, err := utils.ResolvePath(srcFile)
	if err != nil {
		s.logger.Fatal(err)
	}

	isFile, err := utils.IsFile(srcFile)
	if err != nil {
		s.logger.Fatal(err)
	}
	if !isFile {
		s.logger.Errorf("Invalid source file: %s", srcFile)
		return
	}

	dstFile = filepath.Join(s.currentDir, dstFile)
	err = utils.SafeCopyFile(srcFile, dstFile)
	if err != nil {
		s.logger.Fatal(err)
	}
}

func (s *ScriptObject) CopyDir(srcDir, dstDir string) {
	s.logger.Infof("CopyDir: %s => %s", srcDir, dstDir)

	srcDir, err := utils.ResolvePath(srcDir)
	if err != nil {
		s.logger.Fatal(err)
	}

	isDir, err := utils.IsDir(srcDir)
	if err != nil {
		s.logger.Fatal(err)
	}
	if !isDir {
		s.logger.Errorf("Invalid source dirctory: %s", srcDir)
		return
	}

	dstDir = filepath.Join(s.currentDir, dstDir)
	err = utils.SafeCopyDir(srcDir, dstDir)
	if err != nil {
		s.logger.Fatal(err)
	}
}

func (s *ScriptObject) ExportRegistry(key, dstFile string) {
	s.logger.Infof("ExportRegistry: %s => %s", key, dstFile)

	dstFile = filepath.Join(s.currentDir, dstFile)
	err := utils.MakeParentDir(dstFile)
	if err != nil {
		s.logger.Fatal(err)
	}

	cmd := exec.Command("reg", "export", key, dstFile, "/y")
	err = cmd.Run()
	if err != nil {
		s.logger.Fatal(err)
	}
}
