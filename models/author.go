package models

type CreateAuthorReq struct {
	Name  string `json:"name"  validate:"required,min=1,max=50"`
	Email string `json:"email"  validate:"required,email"`
}

type ListAuthorResponse struct {
	Q []struct {
		Uid   string `json:"uid"`
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"q"`
}

type UpdateAuthorReq struct {
	Name  string `json:"name"  validate:"min=1,max=50"`
	Email string `json:"email"  validate:"email"`
}

type Author struct {
	Name  string      `json:"name"`
	Email string      `json:"email"`
	Books interface{} `json:"books"`
}
