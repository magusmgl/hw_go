package api

import "3/cli/config"

type API struct {
	Config config.Config
}

func NewAPI(config config.Config) *API {
	return &API{
		Config: config,
	}
}

func StartApi() {}

func StopApi() {}
