package toiletpaper

import "fmt"

type NameBrand struct {
	Name string
}

func (nm NameBrand) Wipe() {
	fmt.Printf(`kinda gross, but it's name brand: %s`, nm.Name)
}

func GetRandomBrand() NameBrand {
	// this isn't random
	return Charmin()
}
