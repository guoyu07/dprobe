package client

import (
	"net/url"

	"github.com/docker/docker/api/types"

	"golang.org/x/net/context"
)

// NodeRemove removes a Node.
func (cli *Client) NodeRemove(ctx context.Context, nodeID string, options types.NodeRemoveOptions) error ***REMOVED***
	query := url.Values***REMOVED******REMOVED***
	if options.Force ***REMOVED***
		query.Set("force", "1")
	***REMOVED***

	resp, err := cli.delete(ctx, "/nodes/"+nodeID, query, nil)
	ensureReaderClosed(resp)
	return wrapResponseError(err, resp, "node", nodeID)
***REMOVED***
