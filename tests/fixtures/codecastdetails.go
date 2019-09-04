package fixtures

type CodecastDetails struct {
}

func (cd *CodecastDetails) RequestCodecast(permalink string) bool {
	return false
}

func (cd *CodecastDetails) CodecastDetailsOfferPurchaseOf(licenseType string) bool {
	return false
}

func (cd *CodecastDetails) CodecastDetailsTitle() string {
	return "NULL"
}

func (cd *CodecastDetails) CodecastDetailsDate() string {
	return "NULL"
}
