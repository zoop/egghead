package models

type Users struct {
	BaseModel
	UID       string   `gorm:"type:varchar(255);not null;uniqueIndex:idx_users_product" json:"uid"`
	Name      string   `gorm:"varchar(255);not null" json:"name"`
	ProductID int      `gorm:"not null;uniqueIndex:idx_users_product" json:"-"`
	Balance   float64  `gorm:"default:0" json:"balance"`
	Archieved bool     `gorm:"default:false;index" json:"-"`
	Metadata  Metadata `gorm:"type:jsonb;" json:"metadata" sql:"default: 'null'::jsonb"`
}
