package main

import (
	"fmt"
	"strukturdataMVC/controller"
	"strukturdataMVC/entity"
	"strukturdataMVC/model"
	"testing"
)

func TestInsert(t *testing.T) {
	fmt.Print("testing insert")
	member1 := entity.Member{
		Id:    1,
		Nama:  "A",
		Point: 10,
	}

	member2 := entity.Member{
		Id:    2,
		Nama:  "B",
		Point: 11,
	}
	model.ModelInsertMember(member1)
	model.ModelInsertMember(member2)
}

func TestModelViewAll(t *testing.T) {
	fmt.Print("testing insert")
	member1 := entity.Member{
		Id:    1,
		Nama:  "A",
		Point: 10,
	}

	member2 := entity.Member{
		Id:    2,
		Nama:  "B",
		Point: 11,
	}
	model.ModelInsertMember(member1)
	model.ModelInsertMember(member2)
	data := model.ModelViewAllMember()
	fmt.Println(data)
}

func TestFindAllController(t *testing.T) {
	member1 := entity.Member{
		Id:    1,
		Nama:  "A",
		Point: 10,
	}

	member2 := entity.Member{
		Id:    2,
		Nama:  "B",
		Point: 11,
	}
	model.ModelInsertMember(member1)
	model.ModelInsertMember(member2)
	allMember := controller.ControllerFindAllMember()
	fmt.Println(allMember)
}
