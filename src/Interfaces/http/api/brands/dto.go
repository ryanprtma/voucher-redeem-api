package brands

type createBrandRequest struct {
	Name string `json:"name" binding:"required"`
}

type brandResponses struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
