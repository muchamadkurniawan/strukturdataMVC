package view

import (
	"fmt"
	"strukturdataMVC/controller"
)

func InsertMember() {
	var id int
	var nama string
	var point float32
	fmt.Print("masukkan ID Baru : ")
	fmt.Scan(&id)
	fmt.Print("masukkan Nama Baru : ")
	fmt.Scan(&nama)
	fmt.Print("masukkan Point Baru : ")
	fmt.Scan(&point)
	controller.ControllerInsertMember(id, nama, point)
}

func ViewMember() {
	allMember := controller.ControllerFindAllMember()
	for _, member := range allMember {
		fmt.Println("id Member :", member.Id)
		fmt.Println("Nama Member :", member.Nama)
		fmt.Println("Point Member :", member.Point)
	}
}
