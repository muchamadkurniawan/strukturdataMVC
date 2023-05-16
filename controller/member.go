package controller

import (
	"fmt"
	"strukturdataMVC/entity"
	"strukturdataMVC/model"
)

func ControllerFindAllMember() []entity.Member {
	return model.ModelViewAllMember()
}

func ControllerInsertMember(id int, nama string, point float32) {
	data := entity.Member{
		Id:    id,
		Nama:  nama,
		Point: point,
	}
	fmt.Println("testing passing data controller ", data)
	model.ModelInsertMember(data)
}
