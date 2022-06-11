package tools

import (
	"github.com/gofrs/uuid"
	"strings"
)

func GetUUID() string {
	v4, _ := uuid.NewV4()
	str := v4.String()
	str = strings.ReplaceAll(str, "-", "")
	return str
}
