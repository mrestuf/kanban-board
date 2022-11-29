package params

type CreateCategory struct {
	Type string `json:"type" validate:"required"`
}

type UpdateCategory struct {
	Type string `json:"type" validate:"required"`
}



