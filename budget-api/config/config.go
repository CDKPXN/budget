package config

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

var (
	// Conf 是整个包返回的内容，即一个toml.Tree
	Conf = New()
)

// New 返回一个toml.Tree的单例实例
func New() *toml.Tree {
	config, err := toml.LoadFile("./config/config.toml")

	if err != nil {
		fmt.Println("TomlError ", err.Error())
	}

	return config
}
