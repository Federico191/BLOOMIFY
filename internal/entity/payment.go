package entity

type Payment struct {
	PaymentId   uint   `gorm:"primaryKey"`
	PaymentType string `gorm:"size:10"`
	Booking     []Booking
}
