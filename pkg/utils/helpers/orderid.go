package helpers

import (
	"github.com/google/uuid"
	"strconv"
)

func OrderId(prefix string) string {
	var id = strconv.FormatUint(uint64(uuid.New().ID()), 10)
	return prefix + id
}
