package validators

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Validate attempts to validate the RegisterInput's values.
func (ri *RegisterInput) Validate() error {
	if err := validate.Struct(ri); err != nil {
		// this check ensures there wasn't an error
		// with the validation process itself
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		return err
	}
	return nil
}
