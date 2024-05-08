package validators

type ListAccountInput struct {
	PageNumber int `form:"page_number" binding:"required"` // limit
	PageSize   int `form:"page_size" binding:"required"`   // offset
}
