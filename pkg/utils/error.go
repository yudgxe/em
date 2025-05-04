package utils

import "fmt"

func WrapError(err error, op string, opt ...string) error {
	if err != nil {
		if len(opt) > 0 {
			return fmt.Errorf("%s: %s: %w", op, BuildString(": ", opt...), err)
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
