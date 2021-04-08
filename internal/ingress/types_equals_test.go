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

package ingress

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestEqualConfiguration(t *testing.T) {
	ap, _ := filepath.Abs("../../test/manifests/configuration-a.json")
	a, err := readJSON(ap)
	if err != nil {
		t.Errorf("unexpected error reading JSON file: %v", err)
	}

	bp, _ := filepath.Abs("../../test/manifests/configuration-b.json")
	b, err := readJSON(bp)
	if err != nil {
		t.Errorf("unexpected error reading JSON file: %v", err)
	}

	cp, _ := filepath.Abs("../../test/manifests/configuration-c.json")
	c, err := readJSON(cp)
	if err != nil {
		t.Errorf("unexpected error reading JSON file: %v", err)
	}

	if !a.Equal(b) {
		t.Errorf("expected equal configurations (configuration-a.json and configuration-b.json)")
	}

	if !b.Equal(a) {
		t.Errorf("expected equal configurations (configuration-b.json and configuration-a.json)")
	}

	if a.Equal(c) {
		t.Errorf("expected equal configurations (configuration-a.json and configuration-c.json)")
	}
}

func readJSON(p string) (*Configuration, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}

	var c Configuration

	d := json.NewDecoder(f)
	err = d.Decode(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func TestL4ServiceElementsMatch(t *testing.T) {
	testCases := []struct {
		listA    []L4Service
		listB    []L4Service
		expected bool
	}{
		{nil, nil, true},
		{[]L4Service{{Port: 80}}, nil, false},
		{[]L4Service{{Port: 80}}, []L4Service{{Port: 80}}, true},
		{
			[]L4Service{
				{Port: 80, Endpoints: []Endpoint{{Address: "1.1.1.1"}}}},
			[]L4Service{{Port: 80}},
			false,
		},
		{
			[]L4Service{
				{Port: 80, Endpoints: []Endpoint{{Address: "1.1.1.1"}, {Address: "1.1.1.2"}}}},
			[]L4Service{
				{Port: 80, Endpoints: []Endpoint{{Address: "1.1.1.2"}, {Address: "1.1.1.1"}}}},
			true,
		},
		{
			[]L4Service{
				{Port: 80, Backend: L4Backend{Name: "test", Namespace: "default", Protocol: "TCP", ProxyProtocol: ProxyProtocol{Decode: false, Encode: false}}},
			},
			[]L4Service{
				{Port: 80, Backend: L4Backend{Name: "test", Namespace: "default", Protocol: "TCP", ProxyProtocol: ProxyProtocol{Decode: false, Encode: false}}},
			},
			true,
		},
		{
			[]L4Service{
				{Port: 80, Backend: L4Backend{Name: "test", Namespace: "default", Protocol: "TCP", ProxyProtocol: ProxyProtocol{Decode: false, Encode: false}}},
			},
			[]L4Service{
				{Port: 80, Backend: L4Backend{Name: "test", Namespace: "default", Protocol: "TCP", ProxyProtocol: ProxyProtocol{Decode: false, Encode: true}}},
			},
			false,
		},
	}

	for _, testCase := range testCases {
		result := compareL4Service(testCase.listA, testCase.listB)
		if result != testCase.expected {
			t.Errorf("expected %v but returned %v (%v - %v)", testCase.expected, result, testCase.listA, testCase.listB)
		}
	}
}

func TestIntElementsMatch(t *testing.T) {
	testCases := []struct {
		listA    []int
		listB    []int
		expected bool
	}{
		{nil, nil, true},
		{[]int{}, nil, false},
		{[]int{}, []int{1}, false},
		{[]int{1}, []int{1}, true},
		{[]int{1, 2, 3}, []int{3, 2, 1}, true},
	}

	for _, testCase := range testCases {
		result := compareInts(testCase.listA, testCase.listB)
		if result != testCase.expected {
			t.Errorf("expected %v but returned %v (%v - %v)", testCase.expected, result, testCase.listA, testCase.listB)
		}
	}
}

func TestSSLCertMatch(t *testing.T) {
	now := time.Now()
	cert := &SSLCert{
		UID:        "1",
		Name:       "nameA",
		Namespace:  "namespaceA",
		CASHA:      "CASHA_A",
		CN:         []string{"CommonNameA"},
		CRLSHA:     "CRLSHA_A",
		PemSHA:     "PemSHA_A",
		PemCertKey: "PemKeyA",
		ExpireTime: now,
	}

	testCases := []struct {
		sslCertA *SSLCert
		sslCertB *SSLCert
		expected bool
	}{
		{cert, cert, true},
		{
			cert,
			&SSLCert{
				UID:        "1",
				Name:       "nameA",
				Namespace:  "namespaceA",
				CASHA:      "CASHA_A",
				CN:         []string{"CommonNameA"},
				CRLSHA:     "CRLSHA_A",
				PemSHA:     "PemSHA_A",
				PemCertKey: "PemKeyA",
				ExpireTime: now,
			},
			true,
		},
		{
			cert,
			&SSLCert{
				UID:        "1",
				Name:       "nameA",
				Namespace:  "namespaceA",
				CASHA:      "CASHA_New",
				CN:         []string{"CommonNameA"},
				CRLSHA:     "CRLSHA_A",
				PemSHA:     "PemSHA_A",
				PemCertKey: "PemKeyA",
				ExpireTime: now,
			},
			false,
		},
		{
			cert,
			&SSLCert{
				UID:        "1",
				Name:       "nameA",
				Namespace:  "namespaceA",
				CASHA:      "CASHA_A",
				CN:         []string{"CommonNameA"},
				CRLSHA:     "CRLSHA_NEW",
				PemSHA:     "PemSHA_A",
				PemCertKey: "PemKeyA",
				ExpireTime: now,
			},
			false,
		},
		{
			cert,
			&SSLCert{
				UID:        "1",
				Name:       "nameA",
				Namespace:  "namespaceA",
				CASHA:      "CASHA_A",
				CN:         []string{"CommonNameA"},
				CRLSHA:     "CRLSHA_A",
				PemSHA:     "PemSHA_New",
				PemCertKey: "PemKeyA",
				ExpireTime: now,
			},
			false,
		},
		{
			cert,
			&SSLCert{
				UID:        "1",
				Name:       "nameA",
				Namespace:  "namespaceA",
				CASHA:      "CASHA_A",
				CN:         []string{"CommonNameNew"},
				CRLSHA:     "CRLSHA_A",
				PemSHA:     "PemSHA_A",
				PemCertKey: "PemKeyA",
				ExpireTime: now,
			},
			false,
		},
		{
			cert,
			&SSLCert{
				UID:        "1",
				Name:       "nameA",
				Namespace:  "namespaceA",
				CASHA:      "CASHA_A",
				CN:         []string{"CommonNameA"},
				CRLSHA:     "CRLSHA_A",
				PemSHA:     "PemSHA_A",
				PemCertKey: "PemKeyA",
				ExpireTime: now.Add(time.Minute),
			},
			false,
		},
		{
			cert,
			&SSLCert{
				UID:        "1",
				Name:       "nameA",
				Namespace:  "namespaceA",
				CASHA:      "CASHA_A",
				CN:         []string{"CommonNameA"},
				CRLSHA:     "CRLSHA_A",
				PemSHA:     "PemSHA_A",
				PemCertKey: "PemKeyNew",
				ExpireTime: now,
			},
			false,
		},
		{
			cert,
			&SSLCert{
				UID:        "2",
				Name:       "nameA",
				Namespace:  "namespaceA",
				CASHA:      "CASHA_A",
				CN:         []string{"CommonNameA"},
				CRLSHA:     "CRLSHA_A",
				PemSHA:     "PemSHA_A",
				PemCertKey: "PemKeyA",
				ExpireTime: now,
			},
			false,
		},
	}

	for _, testCase := range testCases {
		result := testCase.sslCertA.Equal(testCase.sslCertB)
		if result != testCase.expected {
			t.Errorf("expected %v but returned %v (%v - %v)", testCase.expected, result, testCase.sslCertA, testCase.sslCertB)
		}
	}
}
