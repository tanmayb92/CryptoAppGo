package entity

type Portfolio struct {
	PortfolioID   uint    `json:"portfolioID" gorm:"primary_key"`
	PortfolioName string  `json:"portfolioName"`
	CreationDate  string  `json:"creationDate"`
	UserID        uint    `json:"userId"`
	Enteries      []Entry `json:"enteries" gorm:"foreignkey:EntryID"`
}

func (f *Portfolio) SetUserId(userID uint) {
	f.UserID = userID
}
