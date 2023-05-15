package global

import "os"

func initDefaultAvatar() {
	wd, _ := os.Getwd()
	_, err := os.Stat(wd + "/storage/avatar/e859977fae97b33c7e3e56d46098bd5d.jpg")
	if err != nil {
		panic(err)
	}
}

func initDir() {
	dir, _ := os.Getwd()
	os.MkdirAll(dir+"/log", os.ModePerm)
}
