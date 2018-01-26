package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/docker/docker/api/types/swarm"
	"golang.org/x/net/context"
)

// TaskInspectWithRaw returns the task information and its raw representation..
func (cli *Client) TaskInspectWithRaw(ctx context.Context, taskID string) (swarm.Task, []byte, error) ***REMOVED***
	serverResp, err := cli.get(ctx, "/tasks/"+taskID, nil, nil)
	if err != nil ***REMOVED***
		return swarm.Task***REMOVED******REMOVED***, nil, wrapResponseError(err, serverResp, "task", taskID)
	***REMOVED***
	defer ensureReaderClosed(serverResp)

	body, err := ioutil.ReadAll(serverResp.body)
	if err != nil ***REMOVED***
		return swarm.Task***REMOVED******REMOVED***, nil, err
	***REMOVED***

	var response swarm.Task
	rdr := bytes.NewReader(body)
	err = json.NewDecoder(rdr).Decode(&response)
	return response, body, err
***REMOVED***
