package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var testBin string
var err error

func TestFlowDrills(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FlowDrills Suite")
}

var _ = BeforeSuite(func() {
	testBin, err = gexec.Build("github.com/natalieparellano/flow-drills")
	Expect(err).ToNot(HaveOccurred())
})
