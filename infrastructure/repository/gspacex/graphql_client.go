package gspacex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GraphqlQuery struct {
	Query     string      `json:"query"`
	Variables interface{} `json:"variables"`
}

type errorsParser struct {
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

type GraphqlClient struct {
	endpoint string
}

func (c *GraphqlClient) load(queryResponse interface{}, queryRequest interface{}) error {
	response, err := c.execute(queryRequest)

	if err != nil {
		return err
	}

	err = c.read(response, &queryResponse)

	if err != nil {
		return err
	}

	return nil
}

func (c *GraphqlClient) execute(query interface{}) ([]byte, error) {
	body, _ := json.Marshal(query)
	resp, err := http.Post(c.endpoint, "application/json", bytes.NewBuffer(body))

	if err != nil {
		return nil, fmt.Errorf("can not complete graphql query")
	}

	response, err := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("can not complete graphql query")
	}

	return response, nil
}

func (c *GraphqlClient) read(response []byte, queryResponse interface{}) error {
	errorsResponse := errorsParser{}
	err := json.Unmarshal(response, queryResponse)
	if err != nil {
		return fmt.Errorf("could not parse query")
	}
	err = json.Unmarshal(response, &errorsResponse)
	if err != nil {
		return fmt.Errorf("could not parse error")
	}
	if invalidGraphQLQuery(errorsResponse) {
		return fmt.Errorf("%v", errorsResponse.Errors)
	}
	return nil
}

func invalidGraphQLQuery(errorsResponse errorsParser) bool {
	return len(errorsResponse.Errors) > 0
}
