package github

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type Release struct {
	HTTPClient *http.Client
	Owner      string
	Repository string
}

func (r *Release) GetLatestTag() (string, error) {
	url := fmt.Sprintf("https://github.com/%s/%s/releases/latest", r.Owner, r.Repository)
	res, err := r.httpClient(false).Get(url)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(io.Discard, res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}
	if res.StatusCode/100 != 3 {
		return "", errors.New(res.Status)
	}
	loc := res.Header.Get("Location")
	parts := strings.Split(loc, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1], nil
	}
	return "", errors.New("unexpected Location HTTP header: " + loc)
}

func (r *Release) GetLatestAssets() ([]string, error) {
	url := fmt.Sprintf("https://github.com/%s/%s/releases/latest", r.Owner, r.Repository)
	res, err := r.httpClient(true).Get(url)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	if res.StatusCode/100 != 2 {
		return nil, errors.New(res.Status)
	}
	matches := regexp.MustCompile(`(?i)href="(.+?)"`).FindAllSubmatch(b, -1)
	var out []string
	for _, match := range matches {
		if len(match) != 2 {
			continue
		}
		href := string(match[1])
		if !strings.Contains(href, "/releases/download/") {
			continue
		}
		if strings.HasPrefix(href, "https") {
			out = append(out, href)
		} else {
			out = append(out, "https://github.com"+href)
		}
	}
	return out, nil
}

func (r *Release) httpClient(followRedirect bool) *http.Client {
	c := r.HTTPClient
	if c == nil {
		c = http.DefaultClient
	}
	if followRedirect {
		return c
	}
	c2 := &http.Client{}
	*c2 = *c
	c2.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	return c2
}
