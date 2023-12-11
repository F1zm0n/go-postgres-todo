package main

import (
	"errors"
	"net/http"
	"strings"
)

func GetIdKey(h http.Header) (string, error) {
	val := h.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "Id" {
		return "", errors.New("malformed first part of api header")
	}
	return vals[1], nil
}
