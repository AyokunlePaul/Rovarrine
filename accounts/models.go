package accounts

import (
	"errors"
	"strings"
)

type Request = map[string]interface{}

func IsValidRequestData(r Request) error {
	if _, ok := r["initial_credit"].(float64); !ok {
		return errors.New("initial credit is required")
	}
	if customerId, ok := r["customer_id"].(string); !ok {
		return errors.New("customer id is required")
	} else if strings.TrimSpace(customerId) == "" {
		return errors.New("customer id cannot be empty")
	}

	return nil
}
