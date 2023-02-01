package accounts

import (
	"errors"
	"strings"
)

type Request struct {
	Data map[string]interface{}
}

func (r *Request) IsValidRequestData() error {
	if customerId, ok := r.Data["customer_id"].(string); !ok {
		return errors.New("customer id is required")
	} else if strings.TrimSpace(customerId) == "" {
		return errors.New("customer id cannot be empty")
	}
	if _, ok := r.Data["initial_credit"].(int64); !ok {
		return errors.New("initial credit is required")
	}
	return nil
}
