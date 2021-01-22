package papers

import (
	"github.com/joshprzybyszewski/consumable/toiletpaper"
)

func TPSize(nb toiletpaper.NameBrand) (float64, int) {
	switch nb {
	case toiletpaper.Chamin():
		return 5, 5
	case toiletpaper.Walmart():
		return 4, 4
	}
	return 0, 0
}
