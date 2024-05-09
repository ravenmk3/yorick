package core

import (
	"path/filepath"

	"yorick/utils"
)

const (
	RegKeySystemEnv = `HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\Session Manager\Environment`
	RegKeyUserEnv   = `HKEY_CURRENT_USER\Environment`
)

func (s *ScriptObject) PutHostsFile() {
	srcFile := utils.HostsFilePath
	dstFile := "hosts"
	s.logger.Infof("PutHostsFile: %s => %s", srcFile, dstFile)
	dstFile = filepath.Join(s.currentDir, dstFile)
	err := utils.SafeCopyFile(srcFile, dstFile)
	if err != nil {
		s.logger.Fatal(err)
	}
}

func (s *ScriptObject) ExportRegSystemEnv(dstFile string) {
	s.ExportRegistry(RegKeySystemEnv, dstFile)
}

func (s *ScriptObject) ExportRegUserEnv(dstFile string) {
	s.ExportRegistry(RegKeyUserEnv, dstFile)
}
