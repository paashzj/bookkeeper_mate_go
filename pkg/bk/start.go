package bk

import (
	"bookkeeper_mate_go/pkg/config"
	"bookkeeper_mate_go/pkg/path"
	"bookkeeper_mate_go/pkg/util"
	"fmt"
	"github.com/paashzj/gutil"
	"io/ioutil"
	"os"
	"strings"
)

func Config() error {
	configProp, err := initFromFile(path.BkOriginalConfig)
	if err != nil {
		return err
	}
	if !config.ClusterEnable {
		configProp.Set("advertisedAddress", config.AdvertisedAddress)
	} else {
		configProp.Set("advertisedAddress", os.Getenv("HOSTNAME")+".bookkeeper")
		configProp.Set("useHostNameAsBookieID", "true")
		configProp.Set("httpServerEnabled", "true")
		configProp.Set("zkServers", config.ZkAddress)
		configProp.Set("metadataServiceUri", "zk+hierarchical://"+config.MetaDataServiceUri+"/ledgers")
	}
	return configProp.Write(path.BkConfig)
}

func initFromFile(file string) (*gutil.ConfigProperties, error) {
	configProp := gutil.ConfigProperties{}
	configProp.Init()
	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(fileBytes), "\n")
	for _, line := range split {
		if strings.HasPrefix(line, "#") {
			continue
		}
		array := strings.Split(line, "=")
		if len(array) != 2 {
			util.Logger().Error(fmt.Sprintf("line error %s", line))
			continue
		}
		configProp.Set(array[0], array[1])
	}
	return &configProp, nil
}
