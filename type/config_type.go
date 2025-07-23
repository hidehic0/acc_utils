package types

type FileConfig struct {
	File string `toml:"file"`
	Cmd  string `toml:"cmd"`
}

type Config struct {
	Configs []FileConfig `toml:"configs"`
}
