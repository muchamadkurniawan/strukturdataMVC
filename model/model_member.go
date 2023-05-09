package model

import (
	"strukturdataMVC/Database"
	"strukturdataMVC/entity"
)

func ModelInsertMember(container entity.Member) {
	newGerbong := entity.LinkedlistMember{}
	newGerbong.Data = container
	temp := &Database.DBMember
	if temp.Next == nil {
		temp.Next = &newGerbong
	} else {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &newGerbong
	}
}

func ModelViewAllMember() []entity.Member {
	temp := Database.DBMember.Next
	members := []entity.Member{}
	for temp != nil {
		members = append(members, temp.Data)
		temp = temp.Next
	}
	return members
}
