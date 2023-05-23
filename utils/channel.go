package utils

import "strconv"

func Channel(channel string, id uint) string {
	return channel + strconv.FormatUint(uint64(id), 10)
}