package validation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/LuigiAzevedo/simplebank/util"
)

func TestValidateString(t *testing.T) {

	type test struct {
		name      string
		str       string
		minLength int
		maxLength int
		want      error
	}

	tests := []test{
		{name: "OK", str: util.RandomString(10), minLength: 2, maxLength: 10, want: nil},
		{name: "Special characters", str: "!@#", minLength: 0, maxLength: 3, want: nil},
		{name: "Bigger than max", str: util.RandomString(11), minLength: 2, maxLength: 10, want: fmt.Errorf("must contain from 2-10 characters")},
		{name: "Smaller than min", str: util.RandomString(1), minLength: 2, maxLength: 10, want: fmt.Errorf("must contain from 2-10 characters")},
		{name: "Empty string", str: util.RandomString(0), minLength: 0, maxLength: 10, want: nil},
	}
	for _, tc := range tests {
		got := ValidateString(tc.str, tc.minLength, tc.maxLength)
		require.Equal(t, tc.want, got)

	}
}

func TestValidateUsername(t *testing.T) {

	type test struct {
		name     string
		username string
		want     error
	}

	tests := []test{
		{name: "OK", username: util.RandomString(10), want: nil},
		{name: "Bigger than max", username: util.RandomString(101), want: fmt.Errorf("must contain from 3-100 characters")},
		{name: "Smaller than min / Empty", username: util.RandomString(0), want: fmt.Errorf("must contain from 3-100 characters")},
		{name: "Uppercase", username: "AbC", want: fmt.Errorf("must contain only lowercase letters, digits or underscores")},
		{name: "Special Characters", username: "!@#$", want: fmt.Errorf("must contain only lowercase letters, digits or underscores")},
		// {name: "Empty string", str: util.RandomString(0), minLength: 0, maxLength: 10, want: nil},
	}
	for _, tc := range tests {
		got := ValidateUsername(tc.username)
		require.Equal(t, tc.want, got)

	}
}

func TestValidateFullName(t *testing.T) {

	type test struct {
		name     string
		fullName string
		want     error
	}

	tests := []test{
		{name: "OK", fullName: util.RandomString(20), want: nil},
		{name: "Bigger than max", fullName: util.RandomString(201), want: fmt.Errorf("must contain from 3-200 characters")},
		{name: "Smaller than min / Empty", fullName: util.RandomString(0), want: fmt.Errorf("must contain from 3-200 characters")},
		{name: "Special Characters", fullName: "!@#$", want: fmt.Errorf("must contain only letters and spaces")},
	}
	for _, tc := range tests {
		got := ValidateFullName(tc.fullName)
		require.Equal(t, tc.want, got)

	}
}

func TestValidatePassword(t *testing.T) {

	type test struct {
		name     string
		password string
		want     error
	}

	tests := []test{
		{name: "OK", password: util.RandomString(20), want: nil},
		{name: "Bigger than max", password: util.RandomString(201), want: fmt.Errorf("must contain from 6-100 characters")},
		{name: "Smaller than min / Empty", password: util.RandomString(0), want: fmt.Errorf("must contain from 6-100 characters")},
		{name: "Special Characters", password: "!@#$%¨", want: nil},
	}
	for _, tc := range tests {
		got := ValidatePassword(tc.password)
		require.Equal(t, tc.want, got)

	}
}

func TestValidateEmail(t *testing.T) {

	type test struct {
		name  string
		email string
		want  error
	}

	tests := []test{
		{name: "OK", email: "name@email.com", want: nil},
		{name: "Bigger than max", email: util.RandomString(201), want: fmt.Errorf("must contain from 3-200 characters")},
		{name: "Smaller than min / Empty", email: util.RandomString(0), want: fmt.Errorf("must contain from 3-200 characters")},
		{name: "Special Characters", email: "!@#$%¨", want: nil},
		{name: "Invalid Format", email: util.RandomString(10), want: fmt.Errorf("invalid email address")},
	}
	for _, tc := range tests {
		got := ValidateEmail(tc.email)
		require.Equal(t, tc.want, got)

	}
}
