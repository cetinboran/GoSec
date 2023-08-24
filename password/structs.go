package password

import cla "github.com/cetinboran/goarg/CLA"

// Password Id yazmadım otomatik eklenicek DB ye zaten bu struct kolaylık için
type Password struct {
	UserId   float64
	Title    string
	Url      string
	Password string
}

func PasswordInit() *Password {
	return &Password{}
}

func CheckInputs(args []cla.Input)

func (p *Password) TakeInputs()
