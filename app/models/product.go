package models

type Products struct {
	BaseModel
	UID      string   `gorm:"type:varchar(50);uniqueIndex" json:"uid"`
	Slug     string   `gorm:"type:varchar(50);uniqueIndex; not null" json:"slug"`
	Name     string   `gorm:"type:varchar(50);not null" json:"name"`
	Metadata Metadata `gorm:"type:jsonb" json:"metadata"`
}

func (p *Products) TableName() string {
	return "products"
}
