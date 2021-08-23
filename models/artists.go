package models

import (
	"go.uber.org/zap"
)

// Artist - Artist Model
type Artist struct {
	ArtistID string `json:"artist_id"`
}

// GetArtistByID - Retrieve Artist By ID
func (a *Artist) GetArtistByID(l *zap.SugaredLogger, artistID string) error {
	l.Infof("Getting artist data. ID: %s", artistID)



	return nil
}
