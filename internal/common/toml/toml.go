package toml

import (
	"io"

	"github.com/BurntSushi/toml"
)

func Decode(data string, v interface{}) (toml.MetaData, error) {
	return toml.Decode(data, v)
}

func DecodeFile(path string, v interface{}) (toml.MetaData, error) {
	return toml.DecodeFile(path, v)
}

func DecodeReader(r io.Reader, v interface{}) (toml.MetaData, error) {
	return toml.DecodeReader(r, v)
}
