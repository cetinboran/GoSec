package load

import cla "github.com/cetinboran/goarg/CLA"

type Load struct {
	UserId int
	Path   string
	Format bool
	Errors map[string]*cla.OptionError
}

func LoadInit(userId int) *Load {
	return &Load{UserId: userId}
}

func (l *Load) TakeInputs(args []cla.Input) {
	for _, i2 := range args {
		if i2.Argument == "p" || i2.Argument == "path" {
			l.Path = i2.Value
		}

		if i2.Argument == "format" {
			l.Format = true
		}
	}
}

func (l *Load) HandleInputs() {
	if l.Format != true {
		Path(l)
	}
}
