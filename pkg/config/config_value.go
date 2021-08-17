package config

import "github.com/paashzj/gutil"

// bk config
var (
	ClusterEnable     bool
	ClusterInit       bool
	AdvertisedAddress string
	ZkAddress         string
)

func init() {
	ClusterEnable = gutil.GetEnvBool("CLUSTER_ENABLE", false)
	ClusterInit = gutil.GetEnvBool("CLUSTER_INIT", false)
	AdvertisedAddress = gutil.GetEnvStr("BOOKKEEPER_ADVERTISED_ADDRESS", "localhost")
	ZkAddress = gutil.GetEnvStr("ZK_ADDR", "")
}
