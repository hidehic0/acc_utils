package utils

import (
	"log"
	"os"
	"path"
	"path/filepath"

	config_type "hidehic0/acc_utils/type"

	"github.com/BurntSushi/toml"
)

func getConfigPath() string {
	// 環境変数から設定ディレクトリを取得
	configDir, ok := os.LookupEnv("XDG_CONFIG_HOME")
	if !ok {
		configDir = filepath.Join(os.Getenv("HOME"), ".config")
	}

	return path.Join(configDir, "acc_utils/config.toml")
}

func GetConfig() config_type.Config {
	var res config_type.Config

	_, err := toml.DecodeFile(getConfigPath(), &res)

	if err != nil {
		log.Fatal(err)
		os.Exit(256)
	}

	return res
}
