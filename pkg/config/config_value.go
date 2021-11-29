package config

import "github.com/paashzj/gutil"

// bk config
var (
	ClusterEnable      bool
	ClusterInit        bool
	AdvertisedAddress  string
	ZkAddress          string
	MetaDataServiceUri string
	RemoteMode         bool
	KubernetesCluster  bool
	ZkTlsEnable        bool
	BkTlsEnable        bool
)

func init() {
	ClusterEnable = gutil.GetEnvBool("CLUSTER_ENABLE", false)
	ClusterInit = gutil.GetEnvBool("CLUSTER_INIT", false)
	AdvertisedAddress = gutil.GetEnvStr("BOOKKEEPER_ADVERTISED_ADDRESS", "localhost")
	ZkAddress = gutil.GetEnvStr("ZK_ADDR", "")
	MetaDataServiceUri = gutil.GetEnvStr("META_DATA_SERVICE_URI", "")
	RemoteMode = gutil.GetEnvBool("REMOTE_MODE", true)
	KubernetesCluster = gutil.GetEnvBool("KUBERNETES_CLUSTER", false)
	ZkTlsEnable = gutil.GetEnvBool("ZOOKEEPER_TLS_ENABLE", false)
	BkTlsEnable = gutil.GetEnvBool("BOOKKEEPER_TLS_ENABLE", false)
}
