package utils

import "testing"

func Test_AddParamTo(t *testing.T) {
	for _, test := range []struct {
		name         string
		url          string
		params       []Param
		expecetedUrl string
	}{
		{
			name: "valid_url_with_2_param",
			url:  "https://test.api",
			params: []Param{
				{Key: "key0", Value: "value0"},
				{Key: "key1", Value: "value1"},
			},
			expecetedUrl: "https://test.api?key0=value0&key1=value1",
		},
	} {
		test := test
		t.Run(test.name, func(t *testing.T) {
			url, err := AddParamToUrl(test.url, test.params...)
			if err != nil {
				t.Fatal(err)
			}

			if got, want := url, test.expecetedUrl; got != want {
				t.Fatalf("got: %v, want: %v", got, want)
			}
		})
	}
}
