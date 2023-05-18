package global

import (
	"errors"
	"fmt"
	"opChat/entity"

	"gorm.io/gorm"
)

func initID() {
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

	g := entity.Group{}
	err = Database.Last(&g).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			panic(err)
		}
		NowGroupID = 600000000
	} else {
		NowGroupID = int32(u.ID)
	}
	Log.Info("startup", fmt.Sprintf("NowGroupID = %d", NowGroupID))

	r := entity.Relation{}
	err = Database.Last(&r).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			panic(err)
		}
		NowRelationID = 0
	} else {
		NowRelationID = int32(r.ID)
	}
	Log.Info("startup", fmt.Sprintf("NowRelationID = %d", NowRelationID))

	m := entity.Message{}
	err = Database.Last(&m).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			panic(err)
		}
		NowMessageID = 0
	} else {
		NowMessageID = int32(m.ID)
	}
	Log.Info("startup", fmt.Sprintf("NowMessageID = %d", NowMessageID))
}