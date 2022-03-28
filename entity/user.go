package entity

type User struct {
	UserID     uint        `json:"userID" gorm:"primary_key"`
	FirstName  string      `json:"firstName"`
	LastName   string      `json:"lastName"`
	Age        int         `json:"age"`
	Email      string      `json:"email"`
	Portfolios []Portfolio `json:"portfolios" gorm:"foreignkey:UserID"`
}
