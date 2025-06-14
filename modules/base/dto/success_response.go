package dto

type CreatedEntityResponse struct {
	Id any `json:"id" validate:"required"`
}

func CreatedOK(id any) CreatedEntityResponse {
	return CreatedEntityResponse{Id: id}
}
