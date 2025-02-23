package exceptions

type DomainErrorTranslator struct {
	directories map[string]error
}

func NewDomainErrorTranslator() *DomainErrorTranslator {
	return &DomainErrorTranslator{
		directories: map[string]error{
			"BRAND.NOT_CONTAIN_NEEDED_PROPERTY": NewInvariantError("tidak dapat membuat user baru karena properti yang dibutuhkan tidak ada"),
		},
	}
}

func (det *DomainErrorTranslator) Translate(err error) error {
	if translatedErr, exists := det.directories[err.Error()]; exists {
		return translatedErr
	}
	return err
}
