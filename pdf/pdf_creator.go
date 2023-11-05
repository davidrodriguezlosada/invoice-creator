package pdf

import (
	"fmt"
	"github.com/signintech/gopdf"
	"invoice-creator/model"
	"log"
	"time"
)

const (
	InvoiceText     = "Factura"
	InvoiceDateText = "Fecha factura"
	BillToText      = "Arrendatario:"
	ShipFromText    = "Arrendador:"
	CompanyTINText  = "NIF"
	DetailText      = "Alquiler local mes de %s %d"
	VatText         = "IVA 21%"
	RetentionText   = "Retención IRPF"
	TotalText       = "NETO A PAGAR"
)

func CreateInvoice(invoice model.Invoice, shipFromCompany model.Company, billToCompany model.Company) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	loadFonts(&pdf)
	setRegularStyle(&pdf)

	addLogo(&pdf)

	// Invoice data
	addInvoiceData(&pdf, invoice, 80.0, 120.0)

	// Customer details
	addCompanyData(&pdf, 80.0, 180.0, BillToText, billToCompany)

	// Company details
	addCompanyData(&pdf, 350.0, 180.0, ShipFromText, shipFromCompany)

	// Invoice details
	invoiceDetailsX := 80.0
	pdf.SetXY(invoiceDetailsX, 320.0)
	parsedDate := getParsedDate(invoice.Date)
	detailText := fmt.Sprintf(DetailText, getMonthName(int(parsedDate.Month())), parsedDate.Year())
	addPriceLine(&pdf, invoiceDetailsX, 400.0, detailText, invoice.GrossAmount)
	addPriceLine(&pdf, invoiceDetailsX, 400.0, VatText, invoice.Vat)
	addPriceLine(&pdf, invoiceDetailsX, 400.0, RetentionText, invoice.Retention)
	addPriceLine(&pdf, invoiceDetailsX+80, 320.0, TotalText, invoice.Total)

	err := pdf.WritePdf(invoice.Number + " - " + invoice.ShortName + ".pdf")
	if err != nil {
		log.Print(err.Error())
	}
}

func loadFonts(pdf *gopdf.GoPdf) {
	err := pdf.AddTTFFont("Arial", "resources/Arial.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	err = pdf.AddTTFFont("Arial Bold", "resources/ArialBold.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
}

func setRegularStyle(pdf *gopdf.GoPdf) {
	err := pdf.SetFont("Arial", "", 10)
	if err != nil {
		log.Print(err.Error())
	}
}

func setBoldStyle(pdf *gopdf.GoPdf) {
	err := pdf.SetFont("Arial Bold", "", 10)
	if err != nil {
		log.Print(err.Error())
	}
}

func addLogo(pdf *gopdf.GoPdf) {
	err := pdf.Image("./resources/logo.png", 70.0, 45.0, &gopdf.Rect{W: 200.0, H: 40.0})
	if err != nil {
		log.Print(err.Error())
	}
}

func addInvoiceData(pdf *gopdf.GoPdf, invoice model.Invoice, invoiceDataX float64, invoiceDataY float64) {
	pdf.SetXY(invoiceDataX, invoiceDataY)
	addLine(pdf, invoiceDataX, InvoiceText+": "+invoice.Number)
	addLine(pdf, invoiceDataX, InvoiceDateText+": "+invoice.Date)
}

func addCompanyData(pdf *gopdf.GoPdf, x float64, y float64, title string, company model.Company) {
	pdf.SetY(y)
	addBoldLine(pdf, x, title)
	addBoldLine(pdf, x, company.Name)
	addLine(pdf, x, CompanyTINText+": "+company.TIN)
	addLine(pdf, x, company.AddressLine1)
	addLine(pdf, x, company.AddressLine2)
}

func addLine(pdf *gopdf.GoPdf, initialX float64, text string) {
	pdf.SetX(initialX)
	err := pdf.Text(text)
	if err != nil {
		log.Print(err.Error())
	}
	pdf.Br(20)
}

func addPriceLine(pdf *gopdf.GoPdf, initialX float64, finalX float64, text string, price string) {
	pdf.SetX(initialX)
	//y := pdf.GetY()
	textWidth := 0.0
	fillString := "_"
	for textWidth < finalX {
		fillString += "_"
		textWidth, _ = pdf.MeasureTextWidth(text + fillString + price + " €")
	}
	err := pdf.Text(text + fillString + price + " €")
	if err != nil {
		log.Print(err.Error())
	}

	pdf.Br(28)
}

func addBoldLine(pdf *gopdf.GoPdf, initialX float64, text string) {
	setBoldStyle(pdf)
	addLine(pdf, initialX, text)
	setRegularStyle(pdf)
}

func getParsedDate(date string) time.Time {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Print(err.Error())
	}
	return parsedDate
}

func getMonthName(mes int) string {
	monthNames := map[int]string{
		1:  "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Septiembre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre",
	}

	return monthNames[mes]
}
