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

	for _, pool := range constants.DBPools {
		if _, err := dbclients.Build(dbclients.Config{
			Key:   pool,
			Name:  os.Getenv(constants.DBNameKey),
			User:  os.Getenv(constants.DBUserKey),
			Pass:  os.Getenv(constants.DBPassKey),
			Host:  os.Getenv(constants.DBHostKey),
			Port:  os.Getenv(constants.DBPortKey),
			Query: os.Getenv(constants.DBQueryKey),
		}); err != nil {
			return err
		}
	}

	return nil
}
