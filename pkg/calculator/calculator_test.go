package calculator

import (
	"testing"
)

func TestValidParentheses(t *testing.T) {
	tests := []struct {
		expression string
		wantErr    bool
	}{
		{"(2+2)", false},
		{"(2+2)", false},
		{"(2+2))", true},
		{"((2+2))", false},
		{"((2+2)", true},
		{"2+2)", true},
	}

	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			err := ValidParentheses(tt.expression)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidParentheses() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}


func TestValidExpression(t *testing.T) {
	tests := []struct {
		expression string
		wantErr    bool
	}{
		{"2-+2", true},
		{"a2+2", true},
		{"2+2*2/3-4+5-6+7-8+9-10+11-", true},
		{"2+2*2/3-4+5-6+7-8+9-10+11-12", false},
		{"2+2*2/3-4+5-6+7-8+9-10+11-12+", true},
		{"2+2*2/3-4+5-6+7-8+9-10+11-12+13", false},
	}

	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			err := ValidExpression(tt.expression)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidExpression() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCalc(t *testing.T) {
	tests := []struct {
		expression string
		want       float64
		wantErr    bool
	}{
		{"5/0", 0, true},
		{"2-+2", 0, true},
		{"a2+2", 0, true},
		{"2+2", 4, false},
		{"2+2*2", 6, false},
		{"2+2*2/2", 4, false},
		{"2+2*2/2-1", 3, false},
		{"(2+2)*2", 8, false},
		{"(2+2)*2/2", 4, false},
		{"2+(2*2)/2", 4, false},
		{"2+(2*2)/2-1", 3, false},
		{"2+2*2/2-1+1", 4, false},
		{"2+2*2/2-1+1-1", 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			got, err := Calc(tt.expression)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
