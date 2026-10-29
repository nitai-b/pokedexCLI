// Package pokeapi for using PokeAPI
package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nitai-b/pokedexcli/internal/pokecache"
)

type LocationResponse struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationDetailsResponse struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func GetLocationData(url string, cache *pokecache.Cache) (LocationResponse, error) {
	var mapLocation LocationResponse
	data, hasCache := cache.Get(url)

	if hasCache {
		if err := json.Unmarshal(data, &mapLocation); err != nil {
			return mapLocation, err
		}
	} else {
		resp, err := http.Get(url)
		if err != nil {
			return mapLocation, err
		}

		defer func() {
			if err := resp.Body.Close(); err != nil {
				fmt.Println("error closing response body: ", err)
			}
		}()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mapLocation, err
		}
		cache.Add(url, body)

		if err := json.Unmarshal(body, &mapLocation); err != nil {
			return mapLocation, err
		}
	}
	return mapLocation, nil
}

func GetLocationDetails(url string, cache *pokecache.Cache) (LocationDetailsResponse, error) {
	var locationDetails LocationDetailsResponse
	data, hasCache := cache.Get(url)

	if hasCache {
		if err := json.Unmarshal(data, &locationDetails); err != nil {
			return locationDetails, err
		}
	} else {
		resp, err := http.Get(url)
		if err != nil {
			return locationDetails, err
		}

		defer func() {
			if err := resp.Body.Close(); err != nil {
				fmt.Println("error closing response body: ", err)
			}
		}()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return locationDetails, err
		}
		cache.Add(url, body)

		if err := json.Unmarshal(body, &locationDetails); err != nil {
			return locationDetails, err
		}
	}
	return locationDetails, nil
}

func GetPokemonDetails(pokemonName string, cache *pokecache.Cache) (Pokemon, error) {
	var pokemon Pokemon
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	data, hasCache := cache.Get(url)

	if hasCache {
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return pokemon, err
		}
	} else {
		resp, err := http.Get(url)
		if err != nil {
			return pokemon, err
		}

		defer func() {
			if err := resp.Body.Close(); err != nil {
				fmt.Println("error closing response body: ", err)
			}
		}()

		if resp.StatusCode == http.StatusNotFound {
			return pokemon, fmt.Errorf("'%s' is not a real Pokemon", pokemonName)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return pokemon, err
		}
		cache.Add(url, body)

		if err := json.Unmarshal(body, &pokemon); err != nil {
			return pokemon, err
		}
	}

	return pokemon, nil
}
