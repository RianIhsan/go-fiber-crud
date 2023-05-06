package models

type UserReq struct {
	Nama     string `json:"nama"`
	Kelas    string `json:"kelas"`
	Semester int    `json:"semester"`
	Prodi    string `json:"prodi"`
}
