package models

type Category struct {
	CategoryID uint 		`json:"catid"`
	CategoryName string	`json:"categoryname"`
	Description string	`json:"description"`
}

type Product struct {
	ProductID uint					`json:"prodid"`
	ProductName string			`json:"productname"`
	QuantityPerUnit string	`json:"qtyperunit"`
	UnitPrice float64				`json:"unitprice"`
	ReorderLevel int				`json:"rorlevel"`
}

type Customer struct {
	CustomerID string				`json:"custid"`
	CompanyName string			`json:"companyname"`
	ContactTitle string			`json:"title"`
	ContactName string			`json:"contactname"`
	City string							`json:"city"`
	Country string					`json:"country"`
}

type Order struct {
	OrderID int							`json:"orderid"`
	OrderDate string				`json:"orderdate"`
	RequiredDate string			`json:"reqdate"`
	ShippedDate string			`json:"shipdate"`
}

type Orderdetail struct {
	ProductName string			`json:"prodname"`
	UnitPrice float64				`json:"unitprice"`
	Quantity float64				`json:"quantity"`
}
