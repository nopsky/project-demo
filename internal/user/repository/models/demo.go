package models

// Users [...]
type Users struct {
	ID        int64  `gorm:"primary_key;column:id;type:bigint(20);not null"`
	Username  string `gorm:"column:username;type:varchar(50);not null"`
	Passsword string `gorm:"column:passsword;type:varchar(255);not null"`
	CompanyID int64  `gorm:"column:company_id;type:bigint(20);not null"`
	Nickname  string `gorm:"column:nickname;type:varchar(50);not null"`
	Role      string `gorm:"column:role;type:varchar(50);not null"`
}
