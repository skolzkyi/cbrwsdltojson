package datastructures

import (
	"errors"
)

const (
	cbrNamespace  = "http://web.cbr.ru/"
	inputDTLayout = "2006-01-02"
)

var (
	ErrBadInputDateData = errors.New("fromDate after toDate")
	ErrBadRawData       = errors.New("parse raw date error")
)
