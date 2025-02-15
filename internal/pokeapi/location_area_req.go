package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	data, ok := c.cache.Get(fullURL)
	if ok {
		locationAreasResp := LocationAreasResp{}
		if err := json.Unmarshal(data, &locationAreasResp); err != nil {
			return LocationAreasResp{}, err
		}

		return locationAreasResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("error status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	if err = json.Unmarshal(data, &locationAreasResp); err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(fullURL, data)

	return locationAreasResp, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	data, ok := c.cache.Get(fullURL)
	if ok {
		locationArea := LocationArea{}
		if err := json.Unmarshal(data, &locationArea); err != nil {
			return LocationArea{}, err
		}

		return locationArea, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("error status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	if err = json.Unmarshal(data, &locationArea); err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, data)

	return locationArea, nil
}
