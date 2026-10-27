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
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel int `json:"min_level"`
				MaxLevel int `json:"max_level"`
				Chance   int `json:"chance"`
				Method   struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				ConditionValues []any `json:"condition_values"`
				PokemonDetails  any   `json:"pokemon_details"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
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
