package db

import (
	"strconv"
	"strings"
)

const (
	head = "Gift"
)

func buildBucketKey(mode string, id uint32, args ...string) string {
	key := head + ":" + mode + ":{" + strconv.FormatUint(uint64(id), 10) + "}"
	if len(args) > 0 {
		key += ":" + strings.Join(args, ":")
	}
	return key
}

func buildUserKey(mode string, uid uint64, args ...string) string {
	key := head + ":" + mode + ":{" + strconv.FormatUint(uid, 10) + "}"
	if len(args) > 0 {
		key += ":" + strings.Join(args, ":")
	}
	return key
}

func buildKey(mode string, args ...string) string {
	key := head + ":" + mode
	if len(args) > 0 {
		key += ":" + strings.Join(args, ":")
	}
	return key
}
