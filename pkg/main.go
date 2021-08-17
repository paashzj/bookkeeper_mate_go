package main

import (
	"bookkeeper_mate_go/pkg/bk"
	"bookkeeper_mate_go/pkg/config"
	"bookkeeper_mate_go/pkg/path"
	"bookkeeper_mate_go/pkg/util"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
)

func main() {
	util.Logger().Debug("this is a debug msg")
	util.Logger().Info("this is a info msg")
	util.Logger().Error("this is a error msg")
	err := bk.Config()
	if err != nil {
		if config.ClusterInit {
			fmt.Println(fmt.Sprintf("generate config failed %v", err))
		} else {
			util.Logger().Error("generate config failed ", zap.Error(err))
		}
	}
	if config.ClusterInit {
		err := util.CallScript(path.BkInitScript)
		if err != nil {
			fmt.Println(fmt.Sprintf("bk server init failed %v", err))
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	}
	if config.ClusterEnable {
		err := util.CallScript(path.BkStartScript)
		if err != nil {
			util.Logger().Error("start bk server failed ", zap.Error(err))
			os.Exit(1)
		}
	} else {
		err := util.CallScript(path.BkStartStandaloneScript)
		if err != nil {
			util.Logger().Error("start bk server failed ", zap.Error(err))
			os.Exit(1)
		}
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-interrupt:
			return
		}
	}
}
