package cmd

import "fmt"

// Option is CLI option
type Option struct {
	AuthToken  string
	AuthSecret string
}

// Validate returns validation errors of CLI option
func (o *Option) Validate() []error {
	var ret []error
	if o.AuthToken == "" {
		ret = append(ret, fmt.Errorf("[%s] is required", "auth-token"))
	}
	if o.AuthSecret == "" {
		ret = append(ret, fmt.Errorf("[%s] is required", "auth-secret"))
	}
	return ret
}
