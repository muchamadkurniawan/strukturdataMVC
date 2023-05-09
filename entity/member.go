package entity

type Member struct {
	Id    int
	Nama  string
	Point float32
}

type LinkedlistMember struct {
	Data Member
	Next *LinkedlistMember
}
