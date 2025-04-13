package internal

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Root string `toml:"root"`

	Build struct {
		Bin        string   `toml:"bin"`
		Cmd        string   `toml:"cmd"`
		IncludeExt []string `toml:"include_ext"`
		ExcludeDir []string `toml:"exclude_dir"`
	} `toml:"build"`
}

const PATH = "./.flow.toml"

func ParseTomlConfig() Config {
	var config Config
	_, err := toml.DecodeFile(PATH, &config)
	if err != nil {
		fmt.Printf("ERROR: failed to parse %s\n", PATH)
		return Config{}
	}
	return config
}
