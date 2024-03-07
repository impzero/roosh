package roosh

import (
	"testing"
)

func TestStackCalculator_add(t *testing.T) {
	type fields struct {
		Memory *Stack
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"Test 1", fields{&Stack{1, 2, 3, 4, 5, 6}}, false},
		{"Test 2", fields{&Stack{1, 2, 3, 4, 5, 6, 7}}, false},
		{"Test 3", fields{&Stack{1, 2, 3, 4, 5}}, false},
		{"Test 4", fields{&Stack{1, 2, 3, 4}}, false},
		{"Test 5", fields{&Stack{1, 2, 3}}, false},
		{"Test 6", fields{&Stack{1, 2}}, false},
		{"Test 7", fields{&Stack{1}}, true},
		{"Test 8", fields{&Stack{}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &StackCalculator{
				Memory: tt.fields.Memory,
			}
			if _, err := sc.Add(); (err != nil) != tt.wantErr {
				t.Errorf("StackCalculator.add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStackCalculator_sub(t *testing.T) {
	type fields struct {
		Memory *Stack
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"Test 1", fields{&Stack{1, 2, 3, 4, 5, 6}}, false},
		{"Test 2", fields{&Stack{1, 2, 3, 4, 5, 6, 7}}, false},
		{"Test 3", fields{&Stack{1, 2, 3, 4, 5}}, false},
		{"Test 4", fields{&Stack{1, 2, 3, 4}}, false},
		{"Test 5", fields{&Stack{1, 2, 3}}, false},
		{"Test 6", fields{&Stack{1, 2}}, false},
		{"Test 7", fields{&Stack{1}}, true},
		{"Test 8", fields{&Stack{}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &StackCalculator{
				Memory: tt.fields.Memory,
			}
			if _, err := sc.Sub(); (err != nil) != tt.wantErr {
				t.Errorf("StackCalculator.sub() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStackCalculator_mul(t *testing.T) {
	type fields struct {
		Memory *Stack
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"Test 1", fields{&Stack{1, 2, 3, 4, 5, 6}}, false},
		{"Test 2", fields{&Stack{1, 2, 3, 4, 5, 6, 7}}, false},
		{"Test 3", fields{&Stack{1, 2, 3, 4, 5}}, false},
		{"Test 4", fields{&Stack{1, 2, 3, 4}}, false},
		{"Test 5", fields{&Stack{1, 2, 3}}, false},
		{"Test 6", fields{&Stack{1, 2}}, false},
		{"Test 7", fields{&Stack{1}}, true},
		{"Test 8", fields{&Stack{}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &StackCalculator{
				Memory: tt.fields.Memory,
			}
			if _, err := sc.Mul(); (err != nil) != tt.wantErr {
				t.Errorf("StackCalculator.mul() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStackCalculator_div(t *testing.T) {
	type fields struct {
		Memory *Stack
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"Test 1", fields{&Stack{1, 2, 3, 4, 5, 6}}, false},
		{"Test 2", fields{&Stack{1, 2, 3, 4, 5, 6, 7}}, false},
		{"Test 3", fields{&Stack{1, 2, 3, 4, 5}}, false},
		{"Test 4", fields{&Stack{1, 2, 3, 4}}, false},
		{"Test 5", fields{&Stack{1, 2, 3}}, false},
		{"Test 6", fields{&Stack{1, 2}}, false},
		{"Test 7", fields{&Stack{1}}, true},
		{"Test 8", fields{&Stack{}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &StackCalculator{
				Memory: tt.fields.Memory,
			}
			if _, err := sc.Div(); (err != nil) != tt.wantErr {
				t.Errorf("StackCalculator.div() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
