package lesson_twenty_nine

import (
	"testing"
)

func TestPrintMail(t *testing.T) {
	PrintMail()
}

func TestIsEmail(t *testing.T) {
	_, err := IsEmail("hello")
	if err == nil {
		t.Error("hello is not an email")
	}
	_, err = IsEmail("brandon@gmail.com")
	if err == nil {
		t.Error("brandon@gmail.com is an email")
	}
	_, err = IsEmail("brandon@gmail")
	if err != nil {
		t.Error("brandon@gmail is not an email")
	}
}
