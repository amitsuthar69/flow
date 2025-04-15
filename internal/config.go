package internal

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Root string `toml:"root"`

	Debounce int `toml:"debounce"`

	Build struct {
		Bin          string   `toml:"bin"`
		Cmd          string   `toml:"cmd"`
		IncludeExt   []string `toml:"include_ext"`
		ExcludeDir   []string `toml:"exclude_dir"`
		ExcludeRegex []string `toml:"exclude_regex"`
	} `toml:"build"`
}

const PATH = "./.flow.toml"

func ParseTomlConfig() Config {
	var config Config
	_, err := toml.DecodeFile(PATH, &config)
	if err != nil {
		Log("error", fmt.Sprintf("ERROR: failed to parse %v\n", err))
		return Config{}
	}
	return config
}
