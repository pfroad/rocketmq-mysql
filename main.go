package main

import (
	"context"
	"os"

	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
)

func main() {
	cfg := replication.BinlogSyncerConfig{
		ServerID: 1001,
		Flavor:   "mysql",
		Host:     "10.35.22.61",
		Port:     3306,
		User:     "apuser",
		Password: "airparking",
	}
	syncer := replication.NewBinlogSyncer(cfg)

	streamer, _ := syncer.StartSync(mysql.Position{"mysql-bin.000026", 334759527})

	for {
		ev, _ := streamer.GetEvent(context.Background())
		ev.Dump(os.Stdout)
	}
}
