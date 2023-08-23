package key

import (
	"fmt"

	cla "github.com/cetinboran/goarg/CLA"
	"github.com/cetinboran/gosec/global"
)

// Start'ın her şeyin başladığı yer olsun.
func Start(args []cla.Input) {
	global.Auth(args)

	fmt.Println("Welcome to the key")

}
