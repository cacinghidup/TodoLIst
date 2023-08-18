package models

type SubList struct {
	Id        int          `json:"id" gorm:"primaryKey;column:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Title     string       `json:"title" gorm:"size:100;column:title;check:title ~ '^[a-zA-Z0-9]+$'"`
	Deskripsi string       `json:"deskripsi" gorm:"size:1000"`
	Upload    string       `json:"upload"`
	ListId    int          `json:"list_id"`
	List      ListResponse `gnorm:"foreignKey:ListId"`
}
