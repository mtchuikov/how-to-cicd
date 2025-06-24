package main

import (
	"encoding/json"
	"fmt"
)

type UserV1 struct {
	Username  string
	FirstName string
	LastName  string
}

func (m *UserV1) ToJson() ([]byte, error) {
	result, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("failed to encode user: %w", err)
	}

	return result, nil
}
