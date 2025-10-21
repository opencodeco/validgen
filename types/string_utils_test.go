package types

import "testing"

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{
			name:  "valid simple email",
			email: "test@example.com",
			want:  true,
		},
		{
			name:  "valid email with subdomain",
			email: "user@mail.example.com",
			want:  true,
		},
		{
			name:  "valid email with numbers and special chars",
			email: "user123+tag@example.co.uk",
			want:  true,
		},
		{
			name:  "valid email with dots and underscores",
			email: "first.last_name@domain.org",
			want:  true,
		},
		{
			name:  "invalid email without @",
			email: "invalid.email.com",
			want:  false,
		},
		{
			name:  "invalid email without domain",
			email: "user@",
			want:  false,
		},
		{
			name:  "invalid email without local part",
			email: "@domain.com",
			want:  false,
		},
		{
			name:  "invalid email without TLD",
			email: "user@domain",
			want:  false,
		},
		{
			name:  "invalid email with spaces",
			email: "user @domain.com",
			want:  false,
		},
		{
			name:  "invalid email with multiple @",
			email: "user@@domain.com",
			want:  false,
		},
		{
			name:  "invalid email with short TLD",
			email: "user@domain.c",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidEmail(tt.email); got != tt.want {
				t.Errorf("IsValidEmail(%q) = %v, want %v", tt.email, got, tt.want)
			}
		})
	}
}
