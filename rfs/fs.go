package rfs

import (
	rice "github.com/GeertJohan/go.rice"
	"net/http"
)

type RiceFileSystem struct {
	box *rice.Box
}

func New(boxName string) (*RiceFileSystem, error) {
	box, err := rice.FindBox(boxName)
	if err != nil {
		return nil, err
	}
	return &RiceFileSystem{
		box: box,
	}, nil
}

func (fs *RiceFileSystem) Open(name string) (http.File, error) {
	return fs.box.Open(name)
}
