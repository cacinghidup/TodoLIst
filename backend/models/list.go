package models

type List struct {
	Id        int    `json:"id" gorm:"primaryKey;column:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Title     string `json:"title" gorm:"size:100;column:title;check:title ~ '^[a-zA-Z0-9]+$'"`
	Deskripsi string `json:"deskripsi" gorm:"size:1000"`
	Upload    string `json:"upload"`
	SubList   []SubList
}

type ListResponse struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func (ListResponse) TableName() string {
	return "lists"
}
