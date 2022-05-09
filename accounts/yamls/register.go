package yamls

import (
	"path/filepath"
	"runtime"
)

var (
	_, filePath, _, _ = runtime.Caller(0)
	Directory         = filepath.Dir(filePath)
)
var (
	ConfsDir    string
	AcServerDir string
	JsStoreDir  string

	OperatorCreds string
	OpJwtPath     string

	SYSAccountCreds string
	SYSAccountJwt   string
	SysCredFile     string

	AdminAccountCreds string
	AdminAccountJwt   string
	AdminCredFile     string

	XAccountCreds string
	XAccountJwt   string
	XCredFile     string

	YAccountCreds string
	YAccountJwt   string
	YCredFile     string

	ServerConfigFile    string
	AccountServerConfig string
)

func UpdateCredentialPaths() {
	OperatorCreds = filepath.Join(ConfDir(), "KO.creds")
	OpJwtPath = filepath.Join(ConfDir(), "KO.jwt")

	SYSAccountCreds = filepath.Join(ConfDir(), "SYS_account.creds")
	SYSAccountJwt = filepath.Join(ConfDir(), "SYS.jwt")
	SysCredFile = filepath.Join(ConfDir(), "sys.creds")

	AdminAccountCreds = filepath.Join(ConfDir(), "Admin_account.creds")
	AdminAccountJwt = filepath.Join(ConfDir(), "Admin.jwt")
	AdminCredFile = filepath.Join(ConfDir(), "admin.creds")

	XAccountCreds = filepath.Join(ConfDir(), "X_account.creds")
	XAccountJwt = filepath.Join(ConfDir(), "X.jwt")
	XCredFile = filepath.Join(ConfDir(), "x.creds")

	YAccountCreds = filepath.Join(ConfDir(), "Y_account.creds")
	YAccountJwt = filepath.Join(ConfDir(), "Y.jwt")
	YCredFile = filepath.Join(ConfDir(), "y.creds")

	ServerConfigFile = filepath.Join(ConfDir(), "server.conf")
	AccountServerConfig = filepath.Join(ConfDir(), "nas.conf")
}

func ConfDir() string {
	if len(ConfsDir) > 0 {
		return ConfsDir
	}
	return filepath.Join(Directory, "nats/confs/ac_store")
}

func AccServerDir() string {
	if len(AcServerDir) > 0 {
		return AcServerDir
	}

	return filepath.Join(Directory, "nats/confs/nas_store")
}

func JSStoreDir() string {
	if len(JsStoreDir) > 0 {
		return JsStoreDir
	}
	return filepath.Join(Directory, "nats/confs/jetstream")
}

func SetConfDir(dir string) {
	ConfsDir = dir
}

func SetAccServerDir(dir string) {
	AcServerDir = dir
}

func SetJSStoreDir(dir string) {
	JsStoreDir = dir
}
