package db

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TigerRepository interface {
	Create(tiger *Tiger) error
	GetIdByName(name string) int
	List(first *int, after *string) ([]*Tiger, *string, error)
}

type tigerrepository struct {
	db        *gorm.DB
	sightRepo SightingRepository
}

func NewTigerRepository(db *gorm.DB, sightRepo SightingRepository) TigerRepository {
	return &tigerrepository{
		db:        db,
		sightRepo: sightRepo,
	}
}

func (r *tigerrepository) Create(tiger *Tiger) error {
	if err := r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&tiger).Error; err != nil {
		return err
	}

	tiger.ID = r.GetIdByName(tiger.Name)
	return nil
}

func (r *tigerrepository) GetIdByName(name string) int {
	var tigerId int
	r.db.Model(&Tiger{}).Select("id").Where("name = ?", name).First(&tigerId)
	return tigerId
}

func (r *tigerrepository) List(first *int, after *string) ([]*Tiger, *string, error) {
	var tigers []*Tiger
	query := r.db.Model(&Tiger{})

	if after != nil {
		afterID, err := DecodeCursor(*after)
		if err != nil {
			return nil, nil, err
		}
		query = query.Where("id > ?", afterID)
	}

	if first != nil {
		query = query.Limit(*first + 1)
	}

	if err := query.Find(&tigers).Error; err != nil {
		return nil, nil, err
	}

	var nextCursor *string
	if first != nil && len(tigers) > *first {
		tigers = tigers[:*first]
		nextCursorVal := EncodeCursor(tigers[*first-1].ID)
		nextCursor = &nextCursorVal
	}

	return tigers, nextCursor, nil
}
