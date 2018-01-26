package client

import (
	"io"
	"net/url"
	"time"

	"golang.org/x/net/context"

	"github.com/docker/docker/api/types"
	timetypes "github.com/docker/docker/api/types/time"
)

// ServiceLogs returns the logs generated by a service in an io.ReadCloser.
// It's up to the caller to close the stream.
func (cli *Client) ServiceLogs(ctx context.Context, serviceID string, options types.ContainerLogsOptions) (io.ReadCloser, error) ***REMOVED***
	query := url.Values***REMOVED******REMOVED***
	if options.ShowStdout ***REMOVED***
		query.Set("stdout", "1")
	***REMOVED***

	if options.ShowStderr ***REMOVED***
		query.Set("stderr", "1")
	***REMOVED***

	if options.Since != "" ***REMOVED***
		ts, err := timetypes.GetTimestamp(options.Since, time.Now())
		if err != nil ***REMOVED***
			return nil, err
		***REMOVED***
		query.Set("since", ts)
	***REMOVED***

	if options.Timestamps ***REMOVED***
		query.Set("timestamps", "1")
	***REMOVED***

	if options.Details ***REMOVED***
		query.Set("details", "1")
	***REMOVED***

	if options.Follow ***REMOVED***
		query.Set("follow", "1")
	***REMOVED***
	query.Set("tail", options.Tail)

	resp, err := cli.get(ctx, "/services/"+serviceID+"/logs", query, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return resp.body, nil
***REMOVED***
