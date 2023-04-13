package client

import "fmt"

type Spec struct {
	OpenAPIURL      string `json:"open_api_url"`
	OpenApiLocation string `json:"open_api_location"`
}

func (*Spec) SetDefaults() {
}

func (s *Spec) Validate() error {
	if s.OpenAPIURL != "" && s.OpenApiLocation != "" {
		return fmt.Errorf("only one of open_api_url or open_api_location can be set")
	}

	return nil
}
