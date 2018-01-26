// +build !windows

package main

import (
	"fmt"

	"github.com/docker/docker/integration-cli/checker"
	"github.com/docker/docker/integration-cli/request"
	"github.com/go-check/check"
)

// #19100 This is a deprecated feature test, it should be removed in Docker 1.12
func (s *DockerNetworkSuite) TestDeprecatedDockerNetworkStartAPIWithHostconfig(c *check.C) ***REMOVED***
	netName := "test"
	conName := "foo"
	dockerCmd(c, "network", "create", netName)
	dockerCmd(c, "create", "--name", conName, "busybox", "top")

	config := map[string]interface***REMOVED******REMOVED******REMOVED***
		"HostConfig": map[string]interface***REMOVED******REMOVED******REMOVED***
			"NetworkMode": netName,
		***REMOVED***,
	***REMOVED***
	_, _, err := request.Post(formatV123StartAPIURL("/containers/"+conName+"/start"), request.JSONBody(config))
	c.Assert(err, checker.IsNil)
	c.Assert(waitRun(conName), checker.IsNil)
	networks := inspectField(c, conName, "NetworkSettings.Networks")
	c.Assert(networks, checker.Contains, netName, check.Commentf(fmt.Sprintf("Should contain '%s' network", netName)))
	c.Assert(networks, checker.Not(checker.Contains), "bridge", check.Commentf("Should not contain 'bridge' network"))
***REMOVED***
