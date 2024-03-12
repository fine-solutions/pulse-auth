package utils

import (
	"github.com/google/uuid"
	"strings"
)

const serviceNamePrefix = "pulse"

func GenerateUUID() string {
	return serviceNamePrefix + strings.Replace(uuid.New().String(), "-", "", -1)
}
