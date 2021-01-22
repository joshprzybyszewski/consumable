package lawoffice

import "fmt"

func Printer() {
	fmt.Printf(`printing on %v x %v paper`, LegalPaperHeight(), LegalPaperWidth())
}
