package main

import (
	"fmt"
	"strukturdataMVC/controller"
	"strukturdataMVC/entity"
	"strukturdataMVC/model"
	"strukturdataMVC/view"
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

func TestInsertControllerMember(t *testing.T) {
	controller.ControllerInsertMember(1, "kurniawan", 90)
	controller.ControllerInsertMember(2, "aan", 70)
	allMember := controller.ControllerFindAllMember()
	fmt.Println(allMember)
}

func TestInsertView(t *testing.T) {
	view.InsertMember()
}
