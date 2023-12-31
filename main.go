package main

import (
	"errors"
	"fmt"
	"github.com/thoas/go-funk"
	"invoice-creator/excel"
	"invoice-creator/model"
	"invoice-creator/pdf"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: go run main.go excel_file_path")
		return
	}

	err, shipFromCompany, billToCompanies, invoices := excel.ParseExcel(args[1])
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	invoicesToCreate := funk.Filter(invoices, func(invoice model.Invoice) bool {
		return !invoice.Created
	}).([]model.Invoice)

	for _, invoice := range invoicesToCreate {
		err, billToCompany := getInvoiceCompany(invoice, billToCompanies)
		if err != nil {
			log.Printf("Error: %v", err)
		} else {
			pdf.CreateInvoice(invoice, shipFromCompany, billToCompany)
		}
	}
}

func getInvoiceCompany(invoice model.Invoice, billToCompanies []model.Company) (error, model.Company) {
	for _, company := range billToCompanies {
		if company.ShortName == invoice.ShortName {
			return nil, company
		}
	}
	return errors.New(fmt.Sprintf("company %s not found in companies data", invoice.ShortName)), model.Company{}
}
