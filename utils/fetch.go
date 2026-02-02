package utils

import (
	"io"
	"net/http"
	"sync"

	"github.com/heronhoga/auto-biodata-api/models"
)

func Fetch(service, url string, ch chan models.ApiResult, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		ch <- models.ApiResult{Service: service, Err: err}
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- models.ApiResult{Service: service, Err: err}
		return
	}

	ch <- models.ApiResult{Service: service, Body: body}
}