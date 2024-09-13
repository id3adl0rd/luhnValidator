package validator

import "testing"

func TestIsLuhnValid(t *testing.T) {
	type args struct {
		card int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Success",
			args: args{
				card: 79927398713,
			},
			want: true,
		},
		{
			name: "Fail",
			args: args{
				card: 1111,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLuhnValid(tt.args.card); got != tt.want {
				t.Errorf("IsLuhnValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkLuhn(t *testing.T) {
	type args struct {
		card int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Success",
			args: args{
				card: 79927398713,
			},
			want: 2,
		},
		{
			name: "Fail",
			args: args{
				card: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkLuhn(tt.args.card); got != tt.want {
				t.Errorf("checkLuhn() = %v, want %v", got, tt.want)
			}
		})
	}
}
