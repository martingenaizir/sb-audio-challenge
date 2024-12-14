package modules

import (
	"github.com/martingenaizir/sb-audio-challenge/cmd/constants"
	"github.com/martingenaizir/sb-audio-challenge/cmd/modules/dbclients"
	"github.com/martingenaizir/sb-audio-challenge/cmd/modules/logger"
	"os"
)

func Configure() error {
	setLogLevel()
	return loadDBPools()
}

func setLogLevel() {
	if level, ok := os.LookupEnv(constants.LogLevelKey); ok {
		logger.SetLogLevel(level)
		return
	}

	logger.SetLogLevel(constants.LogLevelDef)
}

func loadDBPools() error {
	// cutting corners here.
	// each pool should have its specific config.

	// TODO use .env
	for _, pool := range constants.DBPools {
		if _, err := dbclients.Build(dbclients.Config{
			Key:   pool,
			Name:  "audio_challenge_db",
			User:  "root",
			Pass:  "",
			Host:  "localhost",
			Port:  "3306",
			Query: "charset=utf8&parseTime=true",
		}); err != nil {
			return err
		}
	}

	return nil
}
