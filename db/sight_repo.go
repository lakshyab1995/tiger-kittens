package db

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SightingRepository interface {
	Create(sighting *Sighting) error
	GetLastSighting(tigerId int) (*Sighting, error)
	GetAllSightings(tigerId int, first *int, after *string) ([]*Sighting, *string, error)
}

type sightingrepository struct {
	db *gorm.DB
}

func NewSightingRepository(db *gorm.DB) SightingRepository {
	return &sightingrepository{db: db}
}

func (r *sightingrepository) Create(sighting *Sighting) error {
	return r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&sighting).Error
}

func (r *sightingrepository) GetLastSighting(tigerId int) (*Sighting, error) {
	var sighting Sighting
	err := r.db.Where("tiger_id = ?", tigerId).Order("timestamp DESC").First(&sighting).Error
	if err != nil {
		return nil, err
	}
	return &sighting, nil
}

func (r *sightingrepository) GetAllSightings(tigerId int, first *int, after *string) ([]*Sighting, *string, error) {
	var sightings []*Sighting
	query := r.db.Model(&Sighting{}).Order("timestamp DESC").Where("tiger_id = ?", tigerId)

	if after != nil {
		afterID, err := DecodeCursor(*after)
		if err != nil {
			return nil, nil, err
		}
		query = query.Where("id = ?", afterID)
	}
	if first != nil {
		query = query.Limit(*first + 1)
	}
	if err := query.Find(&sightings).Error; err != nil {
		return nil, nil, err
	}
	var nextCursor *string
	if first != nil && len(sightings) > *first {
		sightings = sightings[:*first]
		nextCursorVal := EncodeCursor(int(sightings[*first-1].ID))
		nextCursor = &nextCursorVal
	}
	return sightings, nextCursor, nil
}
