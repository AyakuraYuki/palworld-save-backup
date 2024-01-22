package main

import (
	"fmt"
	cp "github.com/otiai10/copy"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	timeLayout = `200601021504`
)

// Config present a runtime parameters
type Config struct {
	SavePath   string // the path to Saves of PalServer
	BackupPath string // a path to where you want to keep the backup files
	DaysKeep   int    // how long of days that you want to keep the recently backups
}

var config Config

func main() {
	_ = os.MkdirAll(config.BackupPath, os.ModePerm)

	backupTime := time.Now().Format(timeLayout)
	dirName := fmt.Sprintf("bak_%s", backupTime)
	dirAbsPath := filepath.Join(config.BackupPath, dirName, "Saved")
	_ = os.MkdirAll(dirAbsPath, os.ModePerm)

	// copy `Saved`
	if err := cp.Copy(config.SavePath, dirAbsPath); err != nil {
		log.Err(err)
		os.Exit(1)
	}
	log.Info().Msgf("backup completed, save to %s", dirAbsPath)

	backups, err := os.ReadDir(config.BackupPath)
	if err != nil {
		log.Err(err)
		os.Exit(1)
	}
	if len(backups) == 0 {
		return
	}

	timeBefore := time.Now().Add(-1 * (time.Duration(config.DaysKeep) * 24 * time.Hour))
	log.Info().Msg("backups:")
	for _, backup := range backups {
		fp := filepath.Join(config.BackupPath, backup.Name())
		msgBuilder := strings.Builder{}
		msgBuilder.WriteString(fmt.Sprintf("- %s", fp))

		stat, err0 := os.Stat(fp)
		if err0 != nil {
			log.Err(err0)
			os.Exit(1)
			return
		}
		if stat.ModTime().Before(timeBefore) {
			_ = os.RemoveAll(fp)
			msgBuilder.WriteString(" (outdated, deleted)")
		}

		log.Info().Msg(msgBuilder.String())
	}
}
