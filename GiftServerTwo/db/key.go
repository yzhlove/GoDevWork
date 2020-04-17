package db

import (
	"strconv"
	"strings"
)

const (
	Top = "Gift"
)

func getStrKey(mode string) string {
	return Top + ":" + mode
}

func getSpecKey(mode string, id uint64, top int, args ...string) string {
	key := Top + ":" + mode
	key += ":{" + strconv.FormatUint(id, 10) + "}:"
	key += strconv.Itoa(top)
	if len(args) > 0 {
		key += ":" + strings.Join(args, ":")
	}
	return key
}
