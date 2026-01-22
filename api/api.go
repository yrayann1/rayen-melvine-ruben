package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"groupie-tracker/models"
)

const BaseURL = "https://groupietrackers.herokuapp.com/api"

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *Client) fetch(endpoint string, target interface{}) error {
	resp, err := c.httpClient.Get(BaseURL + endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("erreur API: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

func (c *Client) GetArtists() ([]models.Artist, error) {
	var artists []models.Artist
	err := c.fetch("/artists", &artists)
	return artists, err
}

func (c *Client) GetRelations() (models.RelationList, error) {
	var relations models.RelationList
	err := c.fetch("/relation", &relations)
	return relations, err
}
