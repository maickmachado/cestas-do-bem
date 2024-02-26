package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Form creates a custom form struct
type Form struct {
	url.Values
	Errors errors
}

// New initializes a new Form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required checks required fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field is required")
		}
	}
}

// TODO: create a function wich check if is a number

// MinLength check the minimum length of string
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("this field must be at least %d character long", length))
		return false
	}
	return true
}

// Has checks if form fields is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "This field is empty")
		return false
	}
	return true
}

// Valid returns true if there is no error
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
