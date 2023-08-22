package datastructures

import (
	"encoding/xml"
	"time"
)

type Repo_debtXML struct { //nolint:revive, stylecheck, nolintlint
	XMLName  xml.Name `xml:"Repo_debtXML" json:"-"` //nolint:revive, stylecheck, nolintlint
	XMLNs    string   `xml:"xmlns,attr" json:"-"`
	FromDate string   `xml:"fromDate"`
	ToDate   string   `xml:"ToDate"`
}

func (data *Repo_debtXML) Init() {
	data.XMLNs = cbrNamespace
}

func (data *Repo_debtXML) Validate() error {
	fromDateDate, err := time.Parse(inputDTLayout, data.FromDate)
	if err != nil {
		return ErrBadRawData
	}
	toDateDate, err := time.Parse(inputDTLayout, data.ToDate)
	if err != nil {
		return ErrBadRawData
	}
	if fromDateDate.After(toDateDate) {
		return ErrBadInputDateData
	}
	return nil
}

type Repo_debtXMLResult struct { //nolint:revive, stylecheck, nolintlint
	// Repo_debt node
	RD []Repo_debtXMLResultElem `xml:"RD" json:"RD"`
}

type Repo_debtXMLResultElem struct {
	Date     time.Time `xml:"Date" json:"Date"`
	Debt     string    `xml:"debt" json:"debt"`
	Debt_auc string    `xml:"debt_auc" json:"debt_auc"` //nolint:revive, stylecheck, nolintlint
	Debt_fix string    `xml:"debt_fix" json:"debt_fix"` //nolint:revive, stylecheck, nolintlint
}
