package main

import "fmt"

type CalcRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type CalcResponse struct {
	Action string  `json:"action"`
	Result float64 `json:"result"`
}

func (r *CalcRequest) Add() float64      { return float64(r.A + r.B) }
func (r *CalcRequest) Subtract() float64 { return float64(r.A - r.B) }
func (r *CalcRequest) Multiply() float64 { return float64(r.A * r.B) }
func (r *CalcRequest) Divide() (float64, error) {
	if r.B == 0 {
		return 0, fmt.Errorf("деление на ноль")
	}
	return float64(r.A) / float64(r.B), nil
}