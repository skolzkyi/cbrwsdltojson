package datastructures

import (
	"errors"
	"time"
)

var (
	ErrBadInputDateData = errors.New("fromDate after toDate")
)

type RequestOnDate struct {
	OnDate string
}

func (data *RequestOnDate) Validate(inputDTLayout string) error {
	_, err := time.Parse(inputDTLayout, data.OnDate)
	if err != nil {
		return err
	}
	return nil
}

type RequestBetweenDate struct {
	FromDate string
	ToDate   string
}

func (data *RequestBetweenDate) Validate(inputDTLayout string) error {
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

type ResponseValuteCursDynamic struct {
	OnDate           time.Time
	ValuteCursOnDate []ResponseValuteCursDynamicElem
}

type ResponseValuteCursDynamicElem struct {
	CursDate string
	Vcode    string
	Vcurs    string
	Vnom     int32
}
