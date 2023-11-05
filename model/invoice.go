package model

type Invoice struct {
	ShortName, Number, Date, GrossAmount, Vat, Retention, Total string
	Created                                                     bool
}
