package entities

import (
	"errors"
)

type Brand struct {
	ID        string
	Name      string
	CreatedAt string
	UpdatedAt string
}

func NewBrand(payload map[string]string) (*Brand, error) {
	if err := verifyPayload(payload); err != nil {
		return nil, err
	}

	return &Brand{
		ID:        payload["id"],
		Name:      payload["name"],
		CreatedAt: payload["created_at"],
		UpdatedAt: payload["updated_at"],
	}, nil
}

func verifyPayload(payload map[string]string) error {
	id, idOk := payload["id"]
	name, nameOk := payload["name"]
	createdAt, createdAtOk := payload["created_at"]
	updatedAt, updatedAtOk := payload["updated_at"]

	if !idOk || !nameOk || !createdAtOk || !updatedAtOk {
		return errors.New("BRAND.NOT_CONTAIN_NEEDED_PROPERTY")
	}

	if id == "" || name == "" || createdAt == "" || updatedAt == "" {
		return errors.New("BRAND.NOT_CONTAIN_NEEDED_PROPERTY")
	}

	return nil
}
