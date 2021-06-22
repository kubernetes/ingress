/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"os"
	"testing"
)

// resetForTesting clears all flag state and sets the usage function as directed.
// After calling resetForTesting, parse errors in flag handling will not
// exit the program.
// Extracted from https://github.com/golang/go/blob/master/src/flag/export_test.go
func resetForTesting(usage func()) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flag.Usage = usage
}

func TestNoMandatoryFlag(t *testing.T) {
	_, _, err := parseFlags()
	if err != nil {
		t.Fatalf("Expected no error but got: %s", err)
	}
}

func TestDefaults(t *testing.T) {
	resetForTesting(func() { t.Fatal("Parsing failed") })

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd",
		"--default-backend-service", "namespace/test",
		"--http-port", "0",
		"--https-port", "0",
	}

	showVersion, conf, err := parseFlags()
	if err != nil {
		t.Fatalf("Unexpected error parsing default flags: %v", err)
	}

	if showVersion {
		t.Fatal("Expected flag \"show-version\" to be false")
	}

	if conf == nil {
		t.Fatal("Expected a controller Configuration")
	}
}

func TestSetupSSLProxy(t *testing.T) {
	// TODO
}

func TestFlagConflict(t *testing.T) {
	resetForTesting(func() { t.Fatal("Parsing failed") })

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd", "--publish-service", "namespace/test", "--http-port", "0", "--https-port", "0", "--publish-status-address", "1.1.1.1"}

	_, _, err := parseFlags()
	if err == nil {
		t.Fatalf("Expected an error parsing flags but none returned")
	}
}

func TestMaxmindEdition(t *testing.T) {
	resetForTesting(func() { t.Fatal("Parsing failed") })

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd", "--publish-service", "namespace/test", "--http-port", "0", "--https-port", "0", "--maxmind-license-key", "0000000", "--maxmind-edition-ids", "GeoLite2-City, TestCheck"}

	_, _, err := parseFlags()
	if err == nil {
		t.Fatalf("Expected an error parsing flags but none returned")
	}
}

func TestMaxmindMirror(t *testing.T) {
	resetForTesting(func() { t.Fatal("Parsing failed") })

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd", "--publish-service", "namespace/test", "--http-port", "0", "--https-port", "0", "--maxmind-mirror", "http://geoip.local", "--maxmind-license-key", "0000000", "--maxmind-edition-ids", "GeoLite2-City, TestCheck"}

	_, _, err := parseFlags()
	if err == nil {
		t.Fatalf("Expected an error parsing flags but none returned")
	}
}

func TestMetricBuckets(t *testing.T) {
	resetForTesting(func() { t.Fatal("Parsing failed") })

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd", "--publish-service", "namespace/test", "--metrics-buckets", "1234567"}

	_, _, err := parseFlags()
	if err == nil {
		t.Fatalf("Expected an error parsing flags but none returned")
	}
}
