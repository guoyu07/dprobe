package client

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/internal/testutil"

	"golang.org/x/net/context"
)

func TestContainerLogsNotFoundError(t *testing.T) ***REMOVED***
	client := &Client***REMOVED***
		client: newMockClient(errorMock(http.StatusNotFound, "Not found")),
	***REMOVED***
	_, err := client.ContainerLogs(context.Background(), "container_id", types.ContainerLogsOptions***REMOVED******REMOVED***)
	if !IsErrNotFound(err) ***REMOVED***
		t.Fatalf("expected a not found error, got %v", err)
	***REMOVED***
***REMOVED***

func TestContainerLogsError(t *testing.T) ***REMOVED***
	client := &Client***REMOVED***
		client: newMockClient(errorMock(http.StatusInternalServerError, "Server error")),
	***REMOVED***
	_, err := client.ContainerLogs(context.Background(), "container_id", types.ContainerLogsOptions***REMOVED******REMOVED***)
	if err == nil || err.Error() != "Error response from daemon: Server error" ***REMOVED***
		t.Fatalf("expected a Server Error, got %v", err)
	***REMOVED***
	_, err = client.ContainerLogs(context.Background(), "container_id", types.ContainerLogsOptions***REMOVED***
		Since: "2006-01-02TZ",
	***REMOVED***)
	testutil.ErrorContains(t, err, `parsing time "2006-01-02TZ"`)
	_, err = client.ContainerLogs(context.Background(), "container_id", types.ContainerLogsOptions***REMOVED***
		Until: "2006-01-02TZ",
	***REMOVED***)
	testutil.ErrorContains(t, err, `parsing time "2006-01-02TZ"`)
***REMOVED***

func TestContainerLogs(t *testing.T) ***REMOVED***
	expectedURL := "/containers/container_id/logs"
	cases := []struct ***REMOVED***
		options             types.ContainerLogsOptions
		expectedQueryParams map[string]string
	***REMOVED******REMOVED***
		***REMOVED***
			expectedQueryParams: map[string]string***REMOVED***
				"tail": "",
			***REMOVED***,
		***REMOVED***,
		***REMOVED***
			options: types.ContainerLogsOptions***REMOVED***
				Tail: "any",
			***REMOVED***,
			expectedQueryParams: map[string]string***REMOVED***
				"tail": "any",
			***REMOVED***,
		***REMOVED***,
		***REMOVED***
			options: types.ContainerLogsOptions***REMOVED***
				ShowStdout: true,
				ShowStderr: true,
				Timestamps: true,
				Details:    true,
				Follow:     true,
			***REMOVED***,
			expectedQueryParams: map[string]string***REMOVED***
				"tail":       "",
				"stdout":     "1",
				"stderr":     "1",
				"timestamps": "1",
				"details":    "1",
				"follow":     "1",
			***REMOVED***,
		***REMOVED***,
		***REMOVED***
			options: types.ContainerLogsOptions***REMOVED***
				// An complete invalid date, timestamp or go duration will be
				// passed as is
				Since: "invalid but valid",
			***REMOVED***,
			expectedQueryParams: map[string]string***REMOVED***
				"tail":  "",
				"since": "invalid but valid",
			***REMOVED***,
		***REMOVED***,
		***REMOVED***
			options: types.ContainerLogsOptions***REMOVED***
				// An complete invalid date, timestamp or go duration will be
				// passed as is
				Until: "invalid but valid",
			***REMOVED***,
			expectedQueryParams: map[string]string***REMOVED***
				"tail":  "",
				"until": "invalid but valid",
			***REMOVED***,
		***REMOVED***,
	***REMOVED***
	for _, logCase := range cases ***REMOVED***
		client := &Client***REMOVED***
			client: newMockClient(func(r *http.Request) (*http.Response, error) ***REMOVED***
				if !strings.HasPrefix(r.URL.Path, expectedURL) ***REMOVED***
					return nil, fmt.Errorf("Expected URL '%s', got '%s'", expectedURL, r.URL)
				***REMOVED***
				// Check query parameters
				query := r.URL.Query()
				for key, expected := range logCase.expectedQueryParams ***REMOVED***
					actual := query.Get(key)
					if actual != expected ***REMOVED***
						return nil, fmt.Errorf("%s not set in URL query properly. Expected '%s', got %s", key, expected, actual)
					***REMOVED***
				***REMOVED***
				return &http.Response***REMOVED***
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte("response"))),
				***REMOVED***, nil
			***REMOVED***),
		***REMOVED***
		body, err := client.ContainerLogs(context.Background(), "container_id", logCase.options)
		if err != nil ***REMOVED***
			t.Fatal(err)
		***REMOVED***
		defer body.Close()
		content, err := ioutil.ReadAll(body)
		if err != nil ***REMOVED***
			t.Fatal(err)
		***REMOVED***
		if string(content) != "response" ***REMOVED***
			t.Fatalf("expected response to contain 'response', got %s", string(content))
		***REMOVED***
	***REMOVED***
***REMOVED***

func ExampleClient_ContainerLogs_withTimeout() ***REMOVED***
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, _ := NewEnvClient()
	reader, err := client.ContainerLogs(ctx, "container_id", types.ContainerLogsOptions***REMOVED******REMOVED***)
	if err != nil ***REMOVED***
		log.Fatal(err)
	***REMOVED***

	_, err = io.Copy(os.Stdout, reader)
	if err != nil && err != io.EOF ***REMOVED***
		log.Fatal(err)
	***REMOVED***
***REMOVED***
