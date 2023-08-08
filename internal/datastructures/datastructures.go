package datastructures

import (
	"errors"
	//"time"
)

const (
	cbrNamespace  = "http://web.cbr.ru/"
	inputDTLayout = "2006-01-02"
)

var (
	ErrBadInputDateData = errors.New("fromDate after toDate")
	ErrBadRawData       = errors.New("parse raw date error")
)

/*
type RequestSeld struct {
	Seld bool
}

type RequestGetCursDynamic struct {
	FromDate   string
	ToDate     string
	ValutaCode string
}

func (data *RequestGetCursDynamic) Validate() error {
	fromDateDate, err := time.Parse(inputDTLayout, data.FromDate)
	if err != nil {
		return err
	}
	toDateDate, err := time.Parse(inputDTLayout, data.ToDate)
	if err != nil {
		return err
	}
	if fromDateDate.After(toDateDate) {
		return ErrBadInputDateData
	}
	return nil
}*/
