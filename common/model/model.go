package model

type (
	MainResponse struct {
		Code int
		Data any
	}

	ErrResponse struct {
		Code  int
		Cause string
	}
)

func (e *ErrResponse) Error() string {
	return e.Cause
}
