package model

type Gender string

const (
	EnumGenderMale   Gender = "male"
	EnumGenderFemale Gender = "female"
)

func NewGender(gender *string) *Gender {
	if gender == nil {
		return nil
	}

	g := Gender(*gender)
	return &g
}

type User struct {
	Name        string  `json:"name"` 
	Surname     string  `json:"surname"`
	Patronymic  *string `json:"patronymic"`
	Nationality *string `json:"nationality"`
	Gender      *Gender `json:"gender"`
	Age         *int32  `json:"age"`
	Id          int64   `json:"id"`
}
