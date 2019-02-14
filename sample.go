package bazelsample

import "github.com/pkg/errors"

// Sample sample method.
func Sample(flag bool) (string, error) {
	if !flag {
		return "", errors.New("error")
	}

	return "Success", nil
}
