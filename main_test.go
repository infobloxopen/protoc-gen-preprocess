package main

import (
	"testing"
)

func TestSamePackage(t *testing.T) {
	tests := []struct {
		packageName string
		path        string
		result      bool
	}{
		{"pb", "google.golang.org/protobuf/types/known/timestamppb", false},
		{"pb", "github.com/infobloxopen/atlas-app-toolkit/rpc/resource", false},
		{"pb", "github.com/Infoblox-CTO/atlas.tagging/pkg/pb", true},
	}

	for _, test := range tests {
		if samePackage(test.packageName, test.path) != test.result {
			t.Errorf("test for (%s, %s) is not correct", test.packageName, test.path)
		}
	}
}
