package entities

import "errors"

type CreateBrand struct {
	Name string
}

func NewCreateBrand(name string) (*CreateBrand, error) {
	if name == "" {
		return nil, errors.New("CREATE_BRAND.NOT_CONTAIN_NEEDED_PROPERTY")
	}

	if len(name) > 50 {
		return nil, errors.New("CREATE_BRAND.BRAND_LIMIT_CHAR")
	}

	return &CreateBrand{
		Name: name,
	}, nil
}
