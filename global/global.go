package global

import (
	"net/http"
	"time"
)

func Init() {
	initDatabase()
	initIris()
	initDir()
	initDefaultAvatar()
	initLogger()
	initID()
}

func StartServe() error {
	return (&http.Server{
		Addr:         ":" + HttpPort,
		Handler:      App,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}).ListenAndServe()
}
