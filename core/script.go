package core

import (
	"github.com/sirupsen/logrus"
)

type ScriptObject struct {
	outputDir  string
	currentDir string
}

func NewScriptObject(outputDir string) *ScriptObject {
	return &ScriptObject{
		outputDir:  outputDir,
		currentDir: outputDir,
	}
}

func (s *ScriptObject) LogInfo(msg string) {
	logrus.Infof(msg)
}
