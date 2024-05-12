package utils

import (
	"gopkg.in/ini.v1"
)

type IniDict map[string]map[string]string

func ReadIniAsDict(file string) (IniDict, error) {
	iniFile, err := ini.Load(file)
	if err != nil {
		return nil, err
	}

	dict := IniDict{}

	for _, section := range iniFile.Sections() {
		secDict := map[string]string{}
		dict[section.Name()] = secDict
		for _, key := range section.Keys() {
			secDict[key.Name()] = key.Value()
		}
	}

	return dict, nil
}
