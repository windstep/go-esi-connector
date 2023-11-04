package go_esi_connector

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

// Just as FYI, we still need to implement some rules on CACHING!
// https://docs.esi.evetech.net/docs/esi_introduction.html

const maxAttempts = 3

type Client struct {
	baseUrl              string
	client               *http.Client
	identificationHeader string
}

type Page struct {
	Current int32
	Total   int32
}

type RateLimit struct {
	remainErrors int
	resetSeconds int
}

func NewClient(baseApiEndpoint string, identificationHeader string, httpClient *http.Client) *Client {
	return &Client{
		baseUrl:              baseApiEndpoint,
		client:               httpClient,
		identificationHeader: identificationHeader,
	}
}

func attachRequiredHeaders(request *http.Request, client Client) *http.Request {
	request.Header.Add("Accept", "application/json")
	// This header **should be** present in most scenarios
	// https://docs.esi.evetech.net/docs/esi_introduction.html
	// Rules section
	request.Header.Add("User-Agent", client.identificationHeader)
	return request
}

func attachAuthHeader(request *http.Request, token string) *http.Request {
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	return request
}

func getRequestsLimitFromHeaders(responseHeaders http.Header) *RateLimit {
	remain, _ := strconv.Atoi(responseHeaders.Get("X-ESI-Error-Limit-Remain"))
	reset, _ := strconv.Atoi(responseHeaders.Get("X-ESI-Error-Limit-Reset"))

	return &RateLimit{
		remainErrors: remain,
		resetSeconds: reset,
	}
}

func getPageFromHeaders(current int32, responseHeaders http.Header) *Page {
	total, _ := strconv.Atoi(responseHeaders.Get("X-Pages"))
	return &Page{
		Current: current,
		Total:   int32(total),
	}
}

func (esi Client) get(endpoint string) ([]byte, http.Header, error) {
	request, err := http.NewRequest(http.MethodGet, esi.baseUrl+endpoint, nil)
	if err != nil {
		slog.Error("cannnot make a request, %v", err)
		return nil, nil, err
	}

	request = attachRequiredHeaders(request, esi)
	return esi.doRequest(request)
}

func (esi Client) getAuth(endpoint string, token string) ([]byte, http.Header, error) {
	request, err := http.NewRequest(http.MethodGet, esi.baseUrl+endpoint, nil)
	if err != nil {
		slog.Error("cannnot make a request, %v", err)
		return nil, nil, err
	}

	request = attachRequiredHeaders(request, esi)
	request = attachAuthHeader(request, token)
	return esi.doRequest(request)
}

func (esi Client) post(endpoint string, content []byte) ([]byte, http.Header, error) {
	request, err := http.NewRequest(http.MethodGet, esi.baseUrl+endpoint, bytes.NewBuffer(content))
	if err != nil {
		slog.Error("cannnot make a request, %v", err)
		return nil, nil, err
	}

	request = attachRequiredHeaders(request, esi)
	return esi.doRequest(request)
}

func (esi Client) postAuth(endpoint string, content []byte, token string) ([]byte, http.Header, error) {
	request, err := http.NewRequest(http.MethodGet, esi.baseUrl+endpoint, bytes.NewBuffer(content))
	if err != nil {
		slog.Error("cannnot make a request, %v", err)
		return nil, nil, err
	}

	request = attachRequiredHeaders(request, esi)
	request = attachAuthHeader(request, token)
	return esi.doRequest(request)
}

func (esi Client) deleteAuth(endpoint string, token string) ([]byte, http.Header, error) {
	request, err := http.NewRequest(http.MethodGet, esi.baseUrl+endpoint, nil)
	if err != nil {
		slog.Error("cannnot make a request, %v", err)
		return nil, nil, err
	}

	request = attachRequiredHeaders(request, esi)
	request = attachAuthHeader(request, token)
	return esi.doRequest(request)
}

func (esi Client) putAuth(endpoint string, content []byte, token string) ([]byte, http.Header, error) {
	request, err := http.NewRequest(http.MethodGet, esi.baseUrl+endpoint, bytes.NewBuffer(content))
	if err != nil {
		slog.Error("cannnot make a request, %v", err)
		return nil, nil, err
	}

	request = attachRequiredHeaders(request, esi)
	request = attachAuthHeader(request, token)
	return esi.doRequest(request)
}

func (esi Client) doRequest(request *http.Request) ([]byte, http.Header, error) {
	for i := 0; i < maxAttempts; i++ {
		// Let's predefine the delay we hit in each iteration
		delay := 1 * time.Second
		response, err := esi.client.Do(request)

		// Unknown error when trying to make a request. Log it and let's try again
		if err != nil {
			slog.Error("esi request failed, error was %v", err)
			continue
		}

		// This is an authorization error - no need to proceed requests couple more time
		if response.StatusCode == 401 || response.StatusCode == 403 {
			slog.Error("user authorization failed")
			break
		}

		limits := getRequestsLimitFromHeaders(response.Header)
		slog.Debug("errors left %d", limits.remainErrors)

		// Trying not to spam ESI Endpoints if we hit "Error limiting"
		// https://developers.eveonline.com/blog/article/error-limiting-imminent
		if response.StatusCode == 420 || response.StatusCode == 520 {
			slog.Error("rate limit exceed, timeout for %d seconds", limits.resetSeconds)
			delay = time.Second * time.Duration(limits.resetSeconds)
			time.Sleep(delay)
			continue
		}

		// Just in case the body contains something useful even for error detection
		message, err := io.ReadAll(response.Body)
		defer response.Body.Close()

		// Unknown error occurred, we should wait a second and try to repeat the request
		// Message will always be present as string, but there might be an "empty string" if some errors occurred
		if response.StatusCode < 200 || response.StatusCode > 299 {
			slog.Error("faced an unknown error with status code %d and message %s", response.StatusCode, string(message))
			time.Sleep(delay)
			continue
		}

		// In case there are errors while decoding - our final clients should really know that the response is somewhat OK-ish string
		if err != nil {
			slog.Error("could not decode esi answer %v", err)
			continue
		}

		// I think we have done all bad scenarios and now our response is ok. Let's return
		slog.Debug("decoded esi answer %v", string(message))
		return message, response.Header, err
	}

	return nil, nil, errors.New(fmt.Sprintf("max attempts of %d exceeded, request failed, giving up", maxAttempts))
}
