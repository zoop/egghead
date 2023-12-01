package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

// Add the base field required for the models
type BaseModel struct {
	*gorm.Model
	ID        uint           `gorm:"primarykey" json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Metadata field type
type Metadata map[string]interface{}

func (metadata Metadata) Value() (driver.Value, error) {
	return json.Marshal(metadata)
}

func (metadata *Metadata) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &metadata)
}

// list all the models
func GetAllModels() []interface{} {
	return []interface{}{
		&Products{},
		&Users{},
		&TransactionHistory{},
	}
}

// CountAndFind retrieves the count and paginated records for a given model.
func FindAndCount(db *gorm.DB, model interface{}, filter map[string]interface{}, page int, limit int) (interface{}, int64, error) {
	var records interface{}
	var count int64

	// Set default values if not provided
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	// Override default if provided

	dbModel := db.Model(model)

	// Apply the filters
	if filter != nil {
		dbModel = dbModel.Where(filter)
	}

	// Fetch records
	if err := dbModel.Offset((page - 1) * limit).Limit(limit).Find(&records).Error; err != nil {
		return nil, 0, err
	}

	if err := dbModel.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return records, count, nil
}
