package utils

import (
	"gopkg.in/ini.v1"
)

type IniDict map[string]map[string]string

func ReadIniAsDict(file string) (IniDict, error) {
	file, err := ExpandUser(file)
	if err != nil {
		return nil, err
	}

	iniFile, err := ini.Load(file)
	if err != nil {
		return nil, err
	}

	dict := IniDict{}

	for _, section := range iniFile.Sections() {
		name := section.Name()
		if name == ini.DefaultSection {
			continue
		}
		secDict := map[string]string{}
		dict[name] = secDict
		for _, key := range section.Keys() {
			secDict[key.Name()] = key.Value()
		}
	}

	return dict, nil
}
