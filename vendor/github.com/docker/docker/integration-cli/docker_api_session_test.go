package main

import (
	"net/http"

	"github.com/docker/docker/integration-cli/checker"
	"github.com/docker/docker/integration-cli/request"
	"github.com/go-check/check"
)

func (s *DockerSuite) TestSessionCreate(c *check.C) ***REMOVED***
	testRequires(c, ExperimentalDaemon)

	res, body, err := request.Post("/session", func(r *http.Request) error ***REMOVED***
		r.Header.Set("X-Docker-Expose-Session-Uuid", "testsessioncreate") // so we don't block default name if something else is using it
		r.Header.Set("Upgrade", "h2c")
		return nil
	***REMOVED***)
	c.Assert(err, checker.IsNil)
	c.Assert(res.StatusCode, checker.Equals, http.StatusSwitchingProtocols)
	c.Assert(res.Header.Get("Upgrade"), checker.Equals, "h2c")
	c.Assert(body.Close(), checker.IsNil)
***REMOVED***

func (s *DockerSuite) TestSessionCreateWithBadUpgrade(c *check.C) ***REMOVED***
	testRequires(c, ExperimentalDaemon)

	res, body, err := request.Post("/session")
	c.Assert(err, checker.IsNil)
	c.Assert(res.StatusCode, checker.Equals, http.StatusBadRequest)
	buf, err := request.ReadBody(body)
	c.Assert(err, checker.IsNil)

	out := string(buf)
	c.Assert(out, checker.Contains, "no upgrade")

	res, body, err = request.Post("/session", func(r *http.Request) error ***REMOVED***
		r.Header.Set("Upgrade", "foo")
		return nil
	***REMOVED***)
	c.Assert(err, checker.IsNil)
	c.Assert(res.StatusCode, checker.Equals, http.StatusBadRequest)
	buf, err = request.ReadBody(body)
	c.Assert(err, checker.IsNil)

	out = string(buf)
	c.Assert(out, checker.Contains, "not supported")
***REMOVED***
