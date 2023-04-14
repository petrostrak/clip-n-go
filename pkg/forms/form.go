package forms

import (
	"net/url"
	"strings"
)

type Form struct {
	url.Values
	Errors errors
}

// New initializes a custom Form struct.
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required checks that specific fields in the form data
// are present and not blank. If a field fails this check,
// add the appropriate message to the form errors.
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}
