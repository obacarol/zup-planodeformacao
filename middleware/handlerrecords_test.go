package middleware

import (
	"planodeformacao-upgrade/models"
	"testing"
)

func TestIsValidID(t *testing.T) {
	type args struct {
		idClient int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "123",
			args: args{
				idClient: 123,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidID(tt.args.idClient); got != tt.want {
				t.Errorf("IsValidID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidTransaction(t *testing.T) {
	type args struct {
		tt models.Transaction
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Pay in",
			args: args{
				tt: "Pay in",
			},
			want: true,
		},
		{
			name: "Payin",
			args: args{
				tt: "Payin",
			},
			want: false,
		},
		{
			name: "Pay In",
			args: args{
				tt: "Pay In",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidTransaction(tt.args.tt); got != tt.want {
				t.Errorf("isValidTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
