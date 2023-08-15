package grpcrequest

import "github.com/gofrs/uuid"

// GenerateUUID - Generates UUID
func generateUUID() (string, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return u.String(), nil
}
