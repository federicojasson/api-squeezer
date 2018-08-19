package validation

type NoOpValidator struct {
}

func NewNoOpValidator() *NoOpValidator {
	return &NoOpValidator{}
}

func (v *NoOpValidator) Validate(data []byte) error {
	return nil
}
