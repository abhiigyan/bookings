package forms

import (
	"net/url"
)

// form creates a custom struct and embeds a url.values obj
type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}
