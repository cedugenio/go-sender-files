package app

import "cedugenio/go-sender-files/properties"

func Start(configPath string) error {
	err := properties.Load(configPath)
	if err != nil {
		return err
	}
	return nil
}
