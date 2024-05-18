package validator

import "testing"

func TestStdlib_Validate(t *testing.T) {
	t.Parallel()
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should parse correct email",
			args: args{
				email: "email@gmail.com",
			},
			want: true,
		},
		{
			name: "Should parse correct email",
			args: args{
				email: "hrvadleo@gmail.com",
			},
			want: true,
		},
		{
			name: "Should parse correct email",
			args: args{
				email: "test@gmail.com",
			},
			want: true,
		},
		{
			name: "Should not parse incorrect email",
			args: args{
				email: "test@gmail.",
			},
			want: false,
		},
		{
			name: "Should not parse incorrect email",
			args: args{
				email: "test@.com",
			},
			want: false,
		},
		{
			name: "Should not parse incorrect email",
			args: args{
				email: "@gmail.com",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := Stdlib{}
			if got := r.Validate(tt.args.email); got != tt.want {
				t.Errorf("Stdlib.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
