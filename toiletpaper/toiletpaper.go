package toiletpaper

import (
	"fmt"

	"github.com/google/uuid"
)

type NameBrand struct {
	Name string
}

func (nm NameBrand) Wipe() {
	fmt.Printf(`kinda gross, but it's name brand: %s`, nm.Name)
}

func GetRandomBrand() NameBrand {
	if uuid.New()[0] == 0 {
		return Charmin()

	}
	return Walmart()
}
