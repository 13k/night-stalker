package models

var HeroModel Model = (*Hero)(nil)

type HeroID uint64

// Hero ...
type Hero struct {
	ID               HeroID `gorm:"column:id;primary_key"`
	Name             string `gorm:"column:name;size:255;unique_index;not null"`
	LocalizedName    string `gorm:"column:localized_name;size:255;not null"`
	ImageFullURL     string `gorm:"column:image_full_url"`
	ImageLargeURL    string `gorm:"column:image_large_url"`
	ImageSmallURL    string `gorm:"column:image_small_url"`
	ImagePortraitURL string `gorm:"column:image_portrait_url"`
	Timestamps
}

func (*Hero) TableName() string {
	return "heroes"
}
