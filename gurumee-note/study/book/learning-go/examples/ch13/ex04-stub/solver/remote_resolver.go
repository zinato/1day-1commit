package solver

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type RemoteResolver struct {
	MathServerURL string
	Client        *http.Client
}

func (rs RemoteResolver) Resolve(ctx context.Context, expr string) (float64, error) {
	URL := fmt.Sprintf("%v?expression=%v", rs.MathServerURL, url.QueryEscape(expr))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := rs.Client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, errors.New(string(contents))
	}

	result, err := strconv.ParseFloat(string(contents), 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
