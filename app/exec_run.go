package app

import (
	"github.com/robertkrimen/otto"
	"github.com/sirupsen/logrus"
)

func ExecRunScript(scriptFile, outputDir string) error {
	logrus.Infof("Script file: %s", scriptFile)
	logrus.Infof("Output directory: %s", outputDir)
	vm := otto.New()
	_, _ = vm.Run(``)
	logrus.Info("All done")
	return nil
}
