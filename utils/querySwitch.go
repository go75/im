package utils

import (
	"errors"
	"strconv"
)

func QuerySwitch(key string, res any) error {
	switch res.(type) {
	case bool:
		if key == "true" {
			res = true
		} else if key == "false" {
			res = false
		} else {
			return errors.New("type error")
		}
	case uint:
		i, err := strconv.Atoi(key)
		if err != nil {
			return err
		}
		res = uint(i)
	case uint8:
		i, err := strconv.Atoi(key)
		if err != nil {
			return err
		}
		res = uint8(i)
	case uint64:
		i, err := strconv.Atoi(key)
		if err != nil {
			return err
		}
		res = uint64(i)
	default:
		return errors.New("type conversion is not supported")
	}
	return nil
}