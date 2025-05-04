package user

type createReq struct {
	Name       string  `json:"name" validate:"required"`
	Surname    string  `json:"surname" validate:"required"`
	Patronymic *string `json:"patronymic"`
}
