package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/liyue201/gin-rice-demo/rfs"
)

func main() {
	r := gin.Default()
	fs, err := rfs.New("statics")
	if err != nil {
		fmt.Printf("no box\n")
		return
	}
	r.StaticFS("/", fs)
	http.ListenAndServe("0.0.0.0:6662", r)
}
