// Copyright The Shipwright Contributors
//
// SPDX-License-Identifier: Apache-2.0

package clusterbuildstrategy_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBuild(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ClusterBuildStrategy Suite")
}
