package richerror

import "encoding/json"

type Code int

const (
	CodeInvalid Code = iota + 1
	CodeForbidden
	CodeNotFound
	CodeUnexpected
)

type RichError struct {
	operation string
	message   string
	code      Code
	innerErr  error
	meta      map[string]interface{}
}

func (r RichError) Error() string {
	if r.message != "" {
		return r.message
	}

	er, ok := r.innerErr.(RichError)
	if !ok {
		return r.innerErr.Error()
	}

	return er.Error()
}

func New(operation string) RichError {
	return RichError{operation: operation}
}

func (r RichError) WithOp(operation string) RichError {
	r.operation = operation

	return r
}

func (r RichError) WithMessage(message string) RichError {
	r.message = message

	return r
}

func (r RichError) WithCode(code Code) RichError {
	r.code = code

	return r
}

func (r RichError) WithError(err error) RichError {
	_, ok := err.(RichError)
	if !ok {
		r.innerErr = New(r.operation).
			WithMessage(err.Error()).
			WithCode(r.code)
		return r
	}
	r.innerErr = err

	return r
}

func (r RichError) WithMeta(meta map[string]interface{}) RichError {
	r.meta = meta

	return r
}

func (r RichError) Code() Code {
	if r.code != 0 {
		return r.code
	}

	er, ok := r.innerErr.(RichError)
	if !ok {
		return 0
	}

	return er.Code()

}

func (r RichError) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Operation string
		Code      Code
		Message   string
		InnerErr  error
		Meta      map[string]interface{}
	}{
		Operation: r.operation,
		Code:      r.code,
		Message:   r.message,
		InnerErr:  r.innerErr,
		Meta:      r.meta,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}
