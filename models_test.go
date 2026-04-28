package main

import "testing"

func TestCalcRequest_Add(t *testing.T) {
	type fields struct {
		A int
		B int
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
		{
			name:   "Сложение положительных чисел",
			fields: fields{A: 10, B: 5},
			want:   15.0,
		},
		{
			name:   "Сложение с нулем",
			fields: fields{A: 0, B: 10},
			want:   10.0,
		},
		{
			name:   "Сложение отрицательных чисел",
			fields: fields{A: -5, B: -5},
			want:   -10.0,
		},
		{
			name:   "Сложение положительного и отрицательного",
			fields: fields{A: 10, B: -3},
			want:   7.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CalcRequest{
				A: tt.fields.A,
				B: tt.fields.B,
			}
			if got := r.Add(); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
