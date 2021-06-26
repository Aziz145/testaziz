package model


type Users struct {
	Id			 int 	    `gorm:"primaryKey;index;autoIncrement:true;NOT NULL;column:id"`
	Username     string     `json:"username" gorm:"type:varchar(100);UNIQUE" binding:"required"`
	Password     string     `json:"password" gorm:"type:varchar(100)" binding:"required"  `
	Name         string      `json:"name" gorm:"type:varchar(100)"binding:"required"  `
}
func (u *Users) TableName() string {
	// custom table name, this is default
	return "users"
}
