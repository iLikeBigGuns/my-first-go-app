package handlers

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

func TestCalcRequest_Subtract(t *testing.T) {
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
			name:   "Вычитание положительных чисел",
			fields: fields{A: 10, B: 5},
			want:   5.0,
		},
		{
			name:   "Результат должен быть отрицательным",
			fields: fields{A: 5, B: 10},
			want:   -5.0,
		},
		{
			name:   "Вычитание нуля",
			fields: fields{A: 10, B: 0},
			want:   10.0,
		},
		{
			name:   "Вычитание отрицательного числа (минус на минус)",
			fields: fields{A: 10, B: -5},
			want:   15.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CalcRequest{
				A: tt.fields.A,
				B: tt.fields.B,
			}
			if got := r.Subtract(); got != tt.want {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcRequest_Multiply(t *testing.T) {
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
			name:   "Умножение положительных чисел",
			fields: fields{A: 10, B: 5},
			want:   50.0,
		},
		{
			name:   "Умножение на ноль",
			fields: fields{A: 10, B: 0},
			want:   0.0,
		},
		{
			name:   "Умножение на отрицательное число",
			fields: fields{A: 10, B: -5},
			want:   -50.0,
		},
		{
			name:   "Умножение двух отрицательных чисел",
			fields: fields{A: -4, B: -5},
			want:   20.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CalcRequest{
				A: tt.fields.A,
				B: tt.fields.B,
			}
			if got := r.Multiply(); got != tt.want {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcRequest_Divide(t *testing.T) {
	type fields struct {
		A int
		B int
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Деление положительных чисел",
			fields:  fields{A: 10, B: 2},
			want:    5.0,
			wantErr: false,
		},
		{
			name:    "Деление с дробным результатом",
			fields:  fields{A: 5, B: 2},
			want:    2.5,
			wantErr: false,
		},
		{
			name:    "Деление на ноль (ошибка)",
			fields:  fields{A: 10, B: 0},
			want:    0.0,
			wantErr: true, // Здесь мы ждем ошибку
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CalcRequest{
				A: tt.fields.A,
				B: tt.fields.B,
			}
			got, err := r.Divide()
			if (err != nil) != tt.wantErr {
				t.Errorf("Divide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Divide() got = %v, want %v", got, tt.want)
			}
		})
	}
}
