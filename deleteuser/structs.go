package deleteuser

import cla "github.com/cetinboran/goarg/CLA"

type DeleteUser struct {
	UserId   int
	Password string
}

func DeleteUserInit() *DeleteUser {
	return &DeleteUser{}
}

func (d *DeleteUser) TakeInputs(args []cla.Input) {

	for _, i2 := range args {
		if i2.Argument == "p" || i2.Argument == "pass" {
			CheckValidPassword(i2.Value) // checks the value.
			d.Password = i2.Value
		}
	}

	// Buraya kadar geldiyse zaten password validtir. Id yi direkt atadım.
	d.UserId = FindUserIdByPassword(d.Password)
}

func (d *DeleteUser) HandleInputs() {
	if d.Password != "" {
		UserDelete(d)
	}
}
