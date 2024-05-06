package app

import (
	"os"

	"github.com/robertkrimen/otto"
	"github.com/sirupsen/logrus"
	"yorick/core"
)

func ExecRunScript(scriptFile, outputDir string) error {
	logrus.Infof("Script file: %s", scriptFile)
	logrus.Infof("Output directory: %s", outputDir)

	content, err := os.ReadFile(scriptFile)
	if err != nil {
		return err
	}
	script := string(content)

	so := core.NewScriptObject(outputDir)

	vm := otto.New()
	err = vm.Set("Y", so)
	if err != nil {
		return err
	}

	_, err = vm.Run(script)
	if err != nil {
		return err
	}

	logrus.Info("All done")
	return nil
}
