package entity

type Entry struct {
	EntryID        uint   `json:"entryID" gorm:"primary_key"`
	CoinName       string `json:"coinName"`
	OrderID        string `json:"orderID"`
	Amount         uint   `json:"amount"`
	Price          uint   `json:"price"`
	TransactionFee uint   `json:"transactionFee"`
	PortfolioID    uint   `json:"portfolioId"`
}

func (f *Entry) SetPortfolioId(portfolioId uint) {
	f.PortfolioID = portfolioId
}
