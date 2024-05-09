package core

import (
	"fmt"
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

func (s *ScriptObject) CopyFile(srcFile, dstFile string) {
	s.logger.Infof("CopyFile: %s => %s", srcFile, dstFile)
	dstFile = filepath.Join(s.currentDir, dstFile)
	err := utils.SafeCopyFile(srcFile, dstFile)
	if err != nil {
		s.logger.Fatal(err)
	}
}

func (s *ScriptObject) CopyDir(srcDir, dstDir string) {
	s.logger.Infof("CopyDir: %s => %s", srcDir, dstDir)
	dstDir = filepath.Join(s.currentDir, dstDir)
	err := utils.SafeCopyDir(srcDir, dstDir)
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
