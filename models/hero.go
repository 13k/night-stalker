package models

var HeroModel Model = (*Hero)(nil)

type HeroID uint64

type Hero struct {
	ID            HeroID `gorm:"column:id;primary_key"`
	Name          string `gorm:"column:name;size:255;unique_index;not null"`
	LocalizedName string `gorm:"column:localized_name;size:255;not null"`
	Timestamps
}

func (*Hero) TableName() string {
	return "heroes"
}
