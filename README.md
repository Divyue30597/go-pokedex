# Refactoring plan

- [ ] Add refactoring basically make functions generics. For example `func GenericRequest[T any](c *Client, url string, search string) (T, error)` we need to make it a method receiver for `c *Client`. So that it can be shared with other packages.

```go
package pokeapi

import (
	"encoding/json"
	"io"
	"log"
)

func GenericRequest[T any](c *Client, url string, search string) (T, error) {
	completeUrl := baseUrl + url + search

	var response T
	res, err := c.httpClient.Get(completeUrl)
	if err != nil {
		return response, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and message: %s", res.StatusCode, res.Body)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return response, err
	}

	return response, nil
}

```
