package helpers

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

type UrlParts struct {
	Host string
	Path string
}

func parseUrl(rawUrl string) (UrlParts, error) {
	re := regexp.MustCompile(`^https?:\/\/https?:\/\/`)
	if ok := re.Match([]byte(rawUrl)); ok {
		return UrlParts{}, errors.New(fmt.Sprintf("URL %s is not valid", rawUrl))
	}

	u, err := url.Parse(rawUrl)
	if err != nil {
		return UrlParts{}, err
	}

	urlParts := UrlParts{
		Host: u.Host,
		Path: u.Path,
	}

	return urlParts, nil
}

func GetDomainFromUrl(rawUrl string) (string, error) {
	urlParts, err := parseUrl(rawUrl)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(`^www\.`)
	domain := re.ReplaceAllString(urlParts.Host, "")

	return domain, nil
}

func GetPathFromUrl(rawUrl string) (string, error) {
	urlParts, err := parseUrl(rawUrl)
	if err != nil {
		return "", err
	}

	if urlParts.Path == "/" {
		return urlParts.Path, nil
	} else {
		return strings.TrimSuffix(urlParts.Path, "/"), nil
	}
}
