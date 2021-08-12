package build

import (
	rice "github.com/GeertJohan/go.rice"
)

func Data(path string) ([]byte, error) {
	return rice.MustFindBox("statics").Bytes(path)
}
