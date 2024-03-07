package web

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pushmarket/render"
)

func (a *API) Decoder(r *http.Request, v interface{}) error {
	if err := render.DefaultDecoder(r, v); err != nil {
		return err
	}

	if validateable, ok := v.(validation.Validatable); ok {
		if err := validateable.Validate(); err != nil {
			return err
		}
	}

	return nil
}
