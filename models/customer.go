package models

type Customer struct {
	Id      int    `gorm:"id"`
	Name    string `gorm:"name"`
	Age     int    `gorm:"age"`
	Email   string `gorm:"email"`
	Address string `gorm:"address"`
}

func (c *Customer) TableName() string {
	return "customer"
}
