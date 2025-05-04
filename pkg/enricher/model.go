package enricher

type agifyResp struct {
	Age *int32 `json:"age"`
}

type genderizeResp struct {
	Gender *string `json:"gender"`
}

type nationalizeResp struct {
	Country []struct {
		Name string `json:"country_id"`
	} `json:"country"`
}

func (this *nationalizeResp) Name() *string {
	if len(this.Country) == 0 {
		return nil
	}

	return &this.Country[0].Name
}
