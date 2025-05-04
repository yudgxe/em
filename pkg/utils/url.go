package utils

import (
	"fmt"
	neturl "net/url"
)

type Param struct {
	Key   string
	Value string
}

func NewParam(key, value string) Param {
	return Param{Key: key, Value: value}
}

func AddParamToValues(values neturl.Values, params ...Param) (neturl.Values, error) {
	for _, param := range params {
		values.Add(param.Key, param.Value)
	}

	return values, nil
}

func AddParamToUrl(url string, params ...Param) (string, error) {
	u, err := neturl.Parse(url)
	if err != nil {
		return "", fmt.Errorf("url.Parse: %w", err)
	}

	values, err := AddParamToValues(u.Query(), params...)
	if err != nil {
		return "", err
	}

	u.RawQuery = values.Encode()

	return u.String(), nil
}
