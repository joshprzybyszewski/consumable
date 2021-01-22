package papers

import (
	"github.com/joshprzybyszewski/consumers/lawoffice"
)

func LegalSize() (float64, int) {
	return lawoffice.LegalPaperWidth(), lawoffice.LegalPaperHeight()
}
