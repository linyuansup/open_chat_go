package main

import (
	"opChat/global"
)

type Aaa struct {
	Bbb string
}

func main() {
	e := global.StartServe()
	if e != nil {
		panic(e)
	}
}
