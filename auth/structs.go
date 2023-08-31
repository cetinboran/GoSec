package auth

import cla "github.com/cetinboran/goarg/CLA"

type Auth struct {
	Password string
	Errors   map[string]*cla.OptionError
}

func AuthInit(errors map[string]*cla.OptionError) *Auth {
	return &Auth{Errors: errors}
}

func (a *Auth) TakeInputs(args []cla.Input) {
	for _, i2 := range args {
		if i2.Argument == "p" || i2.Argument == "pass" {
			a.Password = i2.Value
		}
	}
}

func (a *Auth) HandleInputs() {
	if a.Password != ""{
		Check(a)
	}
}
