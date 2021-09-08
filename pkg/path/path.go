package path

import (
	"os"
	"path/filepath"
)

// zookeeper
var (
	BkHome                     = os.Getenv("BOOKKEEPER_HOME")
	BkConfigDir                = filepath.FromSlash(BkHome + "/conf")
	BkConfig                   = filepath.FromSlash(BkConfigDir + "/bk_server.conf")
	BkStandaloneConfig         = filepath.FromSlash(BkConfigDir + "/standalone.conf")
	BkOriginalConfig           = filepath.FromSlash(BkConfigDir + "/bk_server_original.conf")
	BkOriginalStandaloneConfig = filepath.FromSlash(BkConfigDir + "/standalone_original.conf")
)

// mate
var (
	BkMatePath              = filepath.FromSlash(BkHome + "/mate")
	BkScripts               = filepath.FromSlash(BkMatePath + "/scripts")
	BkInitScript            = filepath.FromSlash(BkScripts + "/init-bookkeeper.sh")
	BkStartScript           = filepath.FromSlash(BkScripts + "/start-bookkeeper.sh")
	BkStartStandaloneScript = filepath.FromSlash(BkScripts + "/start-bookkeeper-standalone.sh")
)
