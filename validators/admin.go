package validators

import "github.com/go-playground/validator/v10"

type AdminLoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Validate attempts to validate the RegisterInput's values.
func (li *AdminLoginInput) Validate() error {
	if err := validate.Struct(li); err != nil {
		// this check ensures there wasn't an error
		// with the validation process itself
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		return err
	}
	return nil
}
