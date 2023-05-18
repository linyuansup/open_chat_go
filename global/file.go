package global

import "os"

func initDefaultAvatar() {
	_, err := os.Stat("." + FilePath + "/avatar/e859977fae97b33c7e3e56d46098bd5d.jpg")
	if err != nil {
		panic(err)
	}
}

func initDir() {
	os.MkdirAll("."+LogPath, os.ModePerm)
}
