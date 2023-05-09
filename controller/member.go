package controller

import (
	"strukturdataMVC/entity"
	"strukturdataMVC/model"
)

func ControllerFindAllMember() []entity.Member {
	return model.ModelViewAllMember()
}
