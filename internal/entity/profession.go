package entity

type Profession struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:100;not null"`
	Description string `gorm:"type:text;not null"`
	Doctors     []Doctor
}
