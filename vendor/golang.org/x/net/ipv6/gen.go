// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

//go:generate go run gen.go

// This program generates system adaptation constants and types,
// internet protocol constants and tables by reading template files
// and IANA protocol registries.
package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func main() ***REMOVED***
	if err := genzsys(); err != nil ***REMOVED***
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	***REMOVED***
	if err := geniana(); err != nil ***REMOVED***
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	***REMOVED***
***REMOVED***

func genzsys() error ***REMOVED***
	defs := "defs_" + runtime.GOOS + ".go"
	f, err := os.Open(defs)
	if err != nil ***REMOVED***
		if os.IsNotExist(err) ***REMOVED***
			return nil
		***REMOVED***
		return err
	***REMOVED***
	f.Close()
	cmd := exec.Command("go", "tool", "cgo", "-godefs", defs)
	b, err := cmd.Output()
	if err != nil ***REMOVED***
		return err
	***REMOVED***
	b, err = format.Source(b)
	if err != nil ***REMOVED***
		return err
	***REMOVED***
	zsys := "zsys_" + runtime.GOOS + ".go"
	switch runtime.GOOS ***REMOVED***
	case "freebsd", "linux":
		zsys = "zsys_" + runtime.GOOS + "_" + runtime.GOARCH + ".go"
	***REMOVED***
	if err := ioutil.WriteFile(zsys, b, 0644); err != nil ***REMOVED***
		return err
	***REMOVED***
	return nil
***REMOVED***

var registries = []struct ***REMOVED***
	url   string
	parse func(io.Writer, io.Reader) error
***REMOVED******REMOVED***
	***REMOVED***
		"https://www.iana.org/assignments/icmpv6-parameters/icmpv6-parameters.xml",
		parseICMPv6Parameters,
	***REMOVED***,
***REMOVED***

func geniana() error ***REMOVED***
	var bb bytes.Buffer
	fmt.Fprintf(&bb, "// go generate gen.go\n")
	fmt.Fprintf(&bb, "// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT\n\n")
	fmt.Fprintf(&bb, "package ipv6\n\n")
	for _, r := range registries ***REMOVED***
		resp, err := http.Get(r.url)
		if err != nil ***REMOVED***
			return err
		***REMOVED***
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK ***REMOVED***
			return fmt.Errorf("got HTTP status code %v for %v\n", resp.StatusCode, r.url)
		***REMOVED***
		if err := r.parse(&bb, resp.Body); err != nil ***REMOVED***
			return err
		***REMOVED***
		fmt.Fprintf(&bb, "\n")
	***REMOVED***
	b, err := format.Source(bb.Bytes())
	if err != nil ***REMOVED***
		return err
	***REMOVED***
	if err := ioutil.WriteFile("iana.go", b, 0644); err != nil ***REMOVED***
		return err
	***REMOVED***
	return nil
***REMOVED***

func parseICMPv6Parameters(w io.Writer, r io.Reader) error ***REMOVED***
	dec := xml.NewDecoder(r)
	var icp icmpv6Parameters
	if err := dec.Decode(&icp); err != nil ***REMOVED***
		return err
	***REMOVED***
	prs := icp.escape()
	fmt.Fprintf(w, "// %s, Updated: %s\n", icp.Title, icp.Updated)
	fmt.Fprintf(w, "const (\n")
	for _, pr := range prs ***REMOVED***
		if pr.Name == "" ***REMOVED***
			continue
		***REMOVED***
		fmt.Fprintf(w, "ICMPType%s ICMPType = %d", pr.Name, pr.Value)
		fmt.Fprintf(w, "// %s\n", pr.OrigName)
	***REMOVED***
	fmt.Fprintf(w, ")\n\n")
	fmt.Fprintf(w, "// %s, Updated: %s\n", icp.Title, icp.Updated)
	fmt.Fprintf(w, "var icmpTypes = map[ICMPType]string***REMOVED***\n")
	for _, pr := range prs ***REMOVED***
		if pr.Name == "" ***REMOVED***
			continue
		***REMOVED***
		fmt.Fprintf(w, "%d: %q,\n", pr.Value, strings.ToLower(pr.OrigName))
	***REMOVED***
	fmt.Fprintf(w, "***REMOVED***\n")
	return nil
***REMOVED***

type icmpv6Parameters struct ***REMOVED***
	XMLName    xml.Name `xml:"registry"`
	Title      string   `xml:"title"`
	Updated    string   `xml:"updated"`
	Registries []struct ***REMOVED***
		Title   string `xml:"title"`
		Records []struct ***REMOVED***
			Value string `xml:"value"`
			Name  string `xml:"name"`
		***REMOVED*** `xml:"record"`
	***REMOVED*** `xml:"registry"`
***REMOVED***

type canonICMPv6ParamRecord struct ***REMOVED***
	OrigName string
	Name     string
	Value    int
***REMOVED***

func (icp *icmpv6Parameters) escape() []canonICMPv6ParamRecord ***REMOVED***
	id := -1
	for i, r := range icp.Registries ***REMOVED***
		if strings.Contains(r.Title, "Type") || strings.Contains(r.Title, "type") ***REMOVED***
			id = i
			break
		***REMOVED***
	***REMOVED***
	if id < 0 ***REMOVED***
		return nil
	***REMOVED***
	prs := make([]canonICMPv6ParamRecord, len(icp.Registries[id].Records))
	sr := strings.NewReplacer(
		"Messages", "",
		"Message", "",
		"ICMP", "",
		"+", "P",
		"-", "",
		"/", "",
		".", "",
		" ", "",
	)
	for i, pr := range icp.Registries[id].Records ***REMOVED***
		if strings.Contains(pr.Name, "Reserved") ||
			strings.Contains(pr.Name, "Unassigned") ||
			strings.Contains(pr.Name, "Deprecated") ||
			strings.Contains(pr.Name, "Experiment") ||
			strings.Contains(pr.Name, "experiment") ***REMOVED***
			continue
		***REMOVED***
		ss := strings.Split(pr.Name, "\n")
		if len(ss) > 1 ***REMOVED***
			prs[i].Name = strings.Join(ss, " ")
		***REMOVED*** else ***REMOVED***
			prs[i].Name = ss[0]
		***REMOVED***
		s := strings.TrimSpace(prs[i].Name)
		prs[i].OrigName = s
		prs[i].Name = sr.Replace(s)
		prs[i].Value, _ = strconv.Atoi(pr.Value)
	***REMOVED***
	return prs
***REMOVED***