package main

import (
  "fmt"
  "time"
  "gotest/models"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
  "strings"
)

var apptitle string = "Sample API in GoLang with SQLite3"

var cat models.Category
var cats []models.Category
var prod models.Product
var prods []models.Product
var cust models.Customer
var customers []models.Customer


func repstr(s string, n int) string {
  return strings.Repeat(s,n)
}

func main() {
  totalrecords := 0
  dbcon, err := sql.Open("sqlite3", "northwind.db")
  if err != nil {
    panic(err)
  }
  starttime := time.Now()
  qry := "Select CategoryID, CategoryName, Description From Categories"
	fmt.Println(qry)
	result, _ := dbcon.Query(qry)
	// if err != nil {
	// 	panic(err.Error())
	// }
	for result.Next() {
		err = result.Scan(&cat.CategoryID, &cat.CategoryName, &cat.Description)
		if err != nil {
			panic(err)
		}
    totalrecords += 1
		cats = append(cats, cat)
	}
  if err == nil {
    fmt.Println(len(cats),"Categories fetched")
    fmt.Printf("%4s | %-30s | %-35s\n","ID", "Category Name", "Description")
    fmt.Printf("%4s | %-30s | %-60s\n", repstr("-",4), repstr("-",30), repstr("-",60))
    for _, row := range cats {
      fmt.Printf("%4d | %-30s | %-35s\n",row.CategoryID, row.CategoryName,row.Description)
    }
  }
  // products
  qry = "Select ProductID, ProductName, QuantityPerUnit, UnitPrice, ReorderLevel From Products"

	result, _ = dbcon.Query(qry)
	// if err != nil {
	// 	panic(err)
	// }
	for result.Next() {
		err = result.Scan(&prod.ProductID, &prod.ProductName, &prod.QuantityPerUnit, &prod.UnitPrice, &prod.ReorderLevel)
		if err != nil {
			panic(err)
		}
    totalrecords += 1
		prods = append(prods, prod)
	}
  fmt.Println(repstr("-",100))
  if err == nil {
    fmt.Println(len(prods),"Products fetched")
    fmt.Printf("%4s | %-35s | %-40s | %10s | %10s\n","ID", "Product Name", "Quantity Per Unit", "Unit Price", "RoR Level")
    fmt.Printf("%4s | %-35s | %-40s | %10s | %10s\n", repstr("-",4), repstr("-",30), repstr("-",40), repstr("-",10), repstr("-",10))
    for _, row := range prods {
      fmt.Printf("%4d | %-35s | %-40s | %10.2f | %10d\n",row.ProductID, row.ProductName,row.QuantityPerUnit, row.UnitPrice, row.ReorderLevel)
    }
  }
  // customers
  qry = "Select CustomerID, CompanyName, ContactName, IfNull(Country,'') as Country From Customers"

	result, _ = dbcon.Query(qry)
	for result.Next() {
		err = result.Scan(&cust.CustomerID,&cust.CompanyName,&cust.ContactName,&cust.Country)
		if err != nil {
			panic(err)
		}
    totalrecords += 1
		customers = append(customers, cust)
	}
  fmt.Println(repstr("-",100))
  if err == nil {
    fmt.Println(len(customers),"Customers fetched")
    fmt.Printf("%-10s | %-35s | %-30s | %-20s\n","CustID", "Company Name", "Contact Name", "Country")
    fmt.Printf("%-10s | %-35s | %-30s | %-20s\n", repstr("-",10), repstr("-",35), repstr("-",30), repstr("-",20))
    for _, row := range customers {
      fmt.Printf("%-10s | %-35s | %-30s | %-20s\n",row.CustomerID, row.CompanyName,row.ContactName, row.Country)
    }
  }

  endtime := time.Now()

  fmt.Println("start:", starttime.Format("15:04:05.000 p"))
  fmt.Println("stop:", endtime.Format("15:04:05.000 p"))
  fmt.Println("total records:", totalrecords)
  fmt.Println("time taken:", endtime.Sub(starttime))
}
