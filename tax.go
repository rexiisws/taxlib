package taxlib

import "log"

const ltfLimit float64 = 200000.00
const insurLimit float64 = 100000.00

func (t *TaxImpl) Calculate(p PersonalIncome) *PersonalTax {

	var tax float64
	if p.Ltf > ltfLimit {
		p.Ltf = ltfLimit
	}
	if p.Insurance > insurLimit {
		p.Insurance = insurLimit
	}

	netIncome := p.Income - p.Ltf - p.Insurance - p.Expense

	log.Println(netIncome)

	if netIncome < 150000 {
		tax = netIncome * 0.05
	} else if netIncome <= 300000 {
		tax += (netIncome-150000)*0.05 + 7500
	} else if netIncome <= 500000 {
		tax += (netIncome-300000)*0.1 + 15000
	} else if netIncome <= 750000 {
		tax += (netIncome-500000)*0.15 + 35000
	} else if netIncome <= 1000000 {
		tax += (netIncome-750000)*0.2 + 72500
	} else if netIncome <= 2000000 {
		tax += (netIncome-1000000)*0.25 + 122500
	} else if netIncome <= 5000000 {
		tax += (netIncome-2000000)*0.3 + 372500
	} else {
		tax += (netIncome-5000000)*0.35 + 1272500
	}

	log.Println(tax)

	return &PersonalTax{Name: p.Name, Tax: tax, NetIncome: netIncome}
}

type PersonalIncome struct {
	Name      string
	Income    float64
	Ltf       float64
	Insurance float64
	Expense   float64
}

type PersonalTax struct {
	Name      string
	Tax       float64
	NetIncome float64
}

type Tax interface {
	Calculate(personalIncome PersonalIncome) PersonalTax
}

type TaxImpl struct {
	Datatable string
}
