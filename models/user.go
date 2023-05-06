package models

type User struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	Nama     string `json:"nama"`
	Kelas    string `json:"kelas"`
	Semester int    `json:"semester"`
	Prodi    string `json:"prodi"`
}
