package enricher

import (
	"context"
	"em/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ageURL         string = "https://api.agify.io"
	genderURL      string = "https://api.genderize.io"
	nationalityURL string = "https://api.nationalize.io"
)

var client *http.Client = http.DefaultClient

type Enricher struct {
	//TODO: logger
}

func New() *Enricher {
	return &Enricher{}
}

func (this *Enricher) GetAgeByName(ctx context.Context, name string) (*int32, error) {
	var body agifyResp
	if err := req(&body, ctx, ageURL, name); err != nil {
		return nil, fmt.Errorf("req: %w", err)
	}

	return body.Age, nil
}

func (this *Enricher) GetGenderByName(ctx context.Context, name string) (*string, error) {
	var body genderizeResp
	if err := req(&body, ctx, genderURL, name); err != nil {
		return nil, fmt.Errorf("req: %w", err)
	}

	return body.Gender, nil
}

func (this *Enricher) GetNationalityByName(ctx context.Context, name string) (*string, error) {
	var body nationalizeResp
	if err := req(&body, ctx, nationalityURL, name); err != nil {
		return nil, fmt.Errorf("req: %w", err)
	}

	return body.Name(), nil
}

func req(body any, ctx context.Context, url string, name string) error {
	url, err := utils.AddParamToUrl(url, utils.NewParam("name", name))
	if err != nil {
		return fmt.Errorf("utils.AddParamToUrl: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("http.NewRequestWithContext: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("this.Client.Do: %w", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return err
	}

	return nil
}
