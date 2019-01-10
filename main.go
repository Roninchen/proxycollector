package main

import (
	"time"

	"proxycollector/collector"
	"proxycollector/scheduler"
	"proxycollector/server"
	"proxycollector/storage"
	"proxycollector/verifier"

	"github.com/cihub/seelog"
)

func main() {
	// Load log.
	scheduler.SetLogger("logConfig.xml")
	defer seelog.Flush()

	// Load database.
	database, err := storage.NewStorage()
	defer database.Close()
	if err != nil {
		seelog.Critical(err)
		panic(err)
	}

	seelog.Infof("database initialize finish.")

	// Start server
	go server.NewServer(database)

	// Verify storage every 5min.
	verifyTicker := time.NewTicker(time.Minute * 5)
	go func() {
		for _ = range verifyTicker.C {
			verifier.VerifyAndDelete(database)
			seelog.Debug("verify database.")
		}
	}()

	configs := collector.NewCollectorConfig("collectorConfig.xml")
	scheduler.Run(configs, database)
}
