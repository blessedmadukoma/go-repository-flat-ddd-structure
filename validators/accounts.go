package validators

type ListAccountInput struct {
	PageNumber int `json:"page_number" binding:"required"` // limit
	PageSize   int `json:"page_size" binding:"required"`   // offset
}
