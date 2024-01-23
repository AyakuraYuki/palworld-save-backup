//go:build windows

package main

import (
	"github.com/pelletier/go-toml/v2"
	"github.com/rs/zerolog/log"
	"os"
)

const (
	configPath = "config.toml"
)

func init() {
	if _, err := os.Stat(configPath); err != nil {
		log.Error().Msg("make sure you have placed the config.toml in the same path with palworld-save-backup.exe")
		os.Exit(1)
	}

	bs, err := os.ReadFile(configPath)
	if err != nil {
		log.Err(err)
		os.Exit(1)
	}

	if err = toml.Unmarshal(bs, &config); err != nil {
		log.Error().Msg(`make sure you have filled out the config.toml, an example of config be liked:

# DO NOTICE THAT YOU SHOULD USE "\\" TO SEPARATE PATH
SavePath = "C:\\steamcmd\\steamapps\\common\\PalServer\\Pal\\Saved"
BackupPath = "C:\\palworld-save-backups"
DaysKeep = 5

BackupPath will be created if there is no dir exists`)
		os.Exit(1)
	}

	if config.SavePath == "" || config.BackupPath == "" {
		log.Error().Msg("an empty value detected with SavePath or BackupPath, make sure you have write the path")
		os.Exit(1)
	}

	if config.DaysKeep == 0 {
		config.DaysKeep = 5
		log.Warn().Msg("set to default 5 days because of no DaysKeep value")
	}
}
