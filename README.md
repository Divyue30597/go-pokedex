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

## Some more feature to work on

- [ ] Update the CLI to support the "up" arrow to cycle through previous commands
- [ ] Simulate battles between pokemon
- [ ] Add more unit tests
- [ ] Refactor your code to organize it better and make it more testable
- [ ] Keep pokemon in a "party" and allow them to level up
- [ ] Allow for pokemon that are caught to evolve after a set amount of time
- [ ] Persist a user's Pokedex to disk so they can save progress between sessions
- [ ] Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you are given choices of areas and just type "left" or "right"
- [ ] Random encounters with wild pokemon
- [ ] Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon
