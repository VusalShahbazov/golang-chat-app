package requests

type GetMessages struct {
	Pagination
}

type StoreRoom struct {
	Name string `json:"name" binding:"required"`
}

type SendMessage struct {
	Body string `json:"body" binding:"required"`
}
