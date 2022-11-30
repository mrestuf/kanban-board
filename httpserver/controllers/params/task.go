package params

type CreateTask struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	CategoryId  int    `json:"category_id" validate:"required"`
}

type UpdateTask struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateTaskStatus struct {
	Status *bool `json:"status" validate:"required"`
}

type UpdateTaskCategory struct {
	CategoryId int `json:"category_id" validate:"required"`
}
