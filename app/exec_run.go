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

	vm := otto.New()
	fo := core.NewFunctionsObject(vm)
	so := core.NewScriptObject(vm, outputDir)

	err = fo.RegisterFuncs()
	if err != nil {
		return err
	}

	err = so.RegisterFuncs()
	if err != nil {
		return err
	}

	_, err = vm.Run(core.InitScript + script)
	if err != nil {
		return err
	}

	err = so.ExecTasks()
	if err != nil {
		return err
	}

	logrus.Info("All done")
	return nil
}
