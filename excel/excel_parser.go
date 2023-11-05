package excel

import (
	"github.com/xuri/excelize/v2"
	"invoice-creator/model"
	"log"
)

const (
	file           = "resources/invoice_data.xlsx"
	dataSheet      = "Data"
	companiesSheet = "Companies"
	invoicesSheet  = "Invoices"
)

func ParseExcel() (error, model.Company, []model.Company, []model.Invoice) {
	xlsx, err := excelize.OpenFile(file)
	if err != nil {
		return err, model.Company{}, nil, nil
	}
	defer func() {
		// Close the spreadsheet.
		if err := xlsx.Close(); err != nil {
			log.Printf("Error: %v", err)
		}
	}()

	shipFromCompany := loadCompanies(xlsx, dataSheet)[0]
	billToCompanies := loadCompanies(xlsx, companiesSheet)
	invoices := loadInvoices(xlsx, invoicesSheet)

	return nil, shipFromCompany, billToCompanies, invoices
}

func loadCompanies(xlsx *excelize.File, sheet string) []model.Company {
	rows, err := xlsx.GetRows(sheet)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil
	}
	var companies []model.Company
	for _, row := range rows[1:] {
		companies = append(companies, model.Company{
			ShortName:    row[0],
			Name:         row[1],
			AddressLine1: row[2],
			AddressLine2: row[3],
			TIN:          row[4],
		})
	}

	return companies
}

func loadInvoices(xlsx *excelize.File, sheet string) []model.Invoice {
	rows, err := xlsx.GetRows(sheet)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil
	}
	var invoices []model.Invoice
	for _, row := range rows[1:] {
		invoices = append(invoices, model.Invoice{
			ShortName:   row[0],
			Number:      row[1],
			Date:        row[2],
			GrossAmount: row[3],
			Vat:         row[4],
			Retention:   row[5],
			Total:       row[6],
			Created:     row[7] == "yes",
		})
	}

	return invoices
}
