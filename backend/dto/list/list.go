package listdto

type CreateList struct {
	Title     string `json:"title" validate:"required"`
	Deskripsi string `json:"deskripsi" validate:"required"`
	Upload    string `json:"file"`
}

type UpdateList struct {
	Title     string `json:"title"`
	Deskripsi string `json:"deskripsi"`
	Upload    string `json:"file"`
}
