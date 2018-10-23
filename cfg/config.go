package cfg

import (
	`errors`
	`gopkg.in/ini.v1`
	`os`
)

func Load(filePath string, cfg interface{}) error {
	var (
		conf *ini.File
		err  error
	)
	if _, err := os.Stat(filePath); err != nil && os.IsExist(err) {
		return errors.New("file not found")
	}
	if conf, err = ini.Load(filePath); err != nil {
		return errors.New("Loading cfg ini error : " + err.Error())
	}

	if err = conf.MapTo(cfg); err != nil {
		return errors.New("map to struct error : " + err.Error())
	}
	return nil
}
