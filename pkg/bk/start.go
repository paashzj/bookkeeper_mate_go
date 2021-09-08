package bk

import (
	"bookkeeper_mate_go/pkg/config"
	"bookkeeper_mate_go/pkg/path"
	"github.com/paashzj/gutil"
	"os"
)

func Config() error {
	if !config.ClusterEnable {
		return ConfigStandalone()
	} else {
		return ConfigCluster()
	}
}

func ConfigCluster() error {
	configProp, err := gutil.ConfigPropFromFile(path.BkOriginalConfig)
	if err != nil {
		return err
	}
	configProp.Set("advertisedAddress", os.Getenv("HOSTNAME")+".bookkeeper")
	configProp.Set("useHostNameAsBookieID", "true")
	configProp.Set("httpServerEnabled", "true")
	configProp.Set("zkServers", config.ZkAddress)
	configProp.Set("metadataServiceUri", "zk+hierarchical://"+config.MetaDataServiceUri+"/ledgers")
	return configProp.Write(path.BkConfig)
}

func ConfigStandalone() error {
	configProp, err := gutil.ConfigPropFromFile(path.BkOriginalStandaloneConfig)
	if err != nil {
		return err
	}
	configProp.Set("advertisedAddress", config.AdvertisedAddress)
	return configProp.Write(path.BkStandaloneConfig)
}
