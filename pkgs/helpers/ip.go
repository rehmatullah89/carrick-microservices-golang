package helpers

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
)

func GetIP(r *http.Request) (string, error){
	var rawIp string

	ipFromHeader := r.Header.Get("X-Forwarded-For")
	if ipFromHeader == "" {
		ipFromRemoteAddr , _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			return "", err
		}

		rawIp = ipFromRemoteAddr
	} else {
		rawIp = ipFromHeader
	}

	if rawIp != "" {
		rawIpParts := strings.Split(rawIp, ",")
		rawIp = rawIpParts[0]
	}

	userIp := net.ParseIP(rawIp)
	if userIp == nil {
		return "", errors.New(fmt.Sprintf("Is not ip: %v", rawIp))
	}

	return userIp.String(), nil
}
