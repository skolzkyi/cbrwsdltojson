package datastructures

import (
	"errors"
	"time"
)

var (
	ErrBadInputDateData = errors.New("fromDate after toDate")
	ErrBadRawData       = errors.New("parse raw date error")
)

type RequestBetweenDate struct {
	FromDate string
	ToDate   string
}

func (data *RequestBetweenDate) Validate(inputDTLayout string) error {
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

type RequestSeld struct {
	Seld bool
}

type RequestGetCursDynamic struct {
	FromDate   string
	ToDate     string
	ValutaCode string
}

func (data *RequestGetCursDynamic) Validate(inputDTLayout string) error {
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
}
