package Models

type PriRequest struct {
	Precio int `json:"precio" binding:"required"`
}
