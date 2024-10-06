package service

import (
	"encoding/json"
	"io"
	"net/http"
	log "rest/logging"
	"rest/models"
)

func GetRecommendations() models.Remm {
	resp, err := http.Get("https://api.jikan.moe/v4/recommendations/anime")
	if err != nil {
		log.Error(err)
		return models.Remm{}
	}
	recomm := models.Remm{}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return models.Remm{}
	}
	json.Unmarshal(bytes, &recomm)
	return recomm

}
