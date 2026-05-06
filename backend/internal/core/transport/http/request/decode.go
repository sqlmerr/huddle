package core_http_request

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	core_errors "github.com/sqlmerr/huddle/backend/internal/core/errors"
)

var requestValidator = validator.New()

func DecodeAndValidateRequest(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return fmt.Errorf("decode json: %w: %w", core_errors.ErrInvalidArgument, err)
	}

	if err := requestValidator.Struct(dest); err != nil {
		return fmt.Errorf("request validation: %w: %w", core_errors.ErrInvalidArgument, err)
	}

	return nil
}
