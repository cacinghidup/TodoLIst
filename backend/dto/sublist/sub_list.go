package sublistdto

type CreateSubList struct {
	Title     string `json:"title" validate:"required"`
	Deskripsi string `json:"deskripsi" validate:"required"`
	Upload    string `json:"file"`
	ListId    int    `json:"list_id" validate:"required"`
}

type UpdateSubList struct {
	Title     string `json:"title"`
	Deskripsi string `json:"deskripsi"`
	Upload    string `json:"file"`
	ListId    int    `json:"list_id"`
}

type ConvertResponse struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Deskripsi string `json:"deskripsi"`
	Upload    string `json:"upload"`
	ListId    int    `json:"list_id"`
}
