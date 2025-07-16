package helpers

import (
	"kutamukti-api/pkg/exceptions"
	"net/http"
	"strings"
)

func StringPointer(b []byte) *string {
	str := string(b)
	return &str
}

func NormalizePhoneNumber(phone string) (string, *exceptions.Exception) {
	phone = strings.TrimSpace(phone)

	if strings.HasPrefix(phone, "0") {
		return "62" + phone[1:], nil
	}

	if strings.HasPrefix(phone, "62") {
		return phone, nil
	}

	return "", exceptions.NewException(http.StatusBadRequest, "invalid phone number format: must start with 0, 62, or +62")
}
