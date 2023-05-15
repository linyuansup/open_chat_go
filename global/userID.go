package global

import (
	"errors"
	"fmt"
	"opChat/entity"

	"gorm.io/gorm"
)

func initUserID() {
	u := entity.User{}
	err := Database.Last(&u).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			panic(err)
		}
		NowUserID = 100000000
	} else {
		NowUserID = int32(u.ID)
	}
	Log.Info("startup", fmt.Sprintf("NowUserID = %d", NowUserID))
}