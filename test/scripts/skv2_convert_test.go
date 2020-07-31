package sync_scripts_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-apis/hack/convert"
	"io/ioutil"
	"os"
	"os/exec"
)

var _ = Describe("Convert to Skv2", func() {

	It("works", func() {
		protoFile := "base_skv1.proto"
		testProtoFile := "skv1_to_skv2_test.proto"

		cpCmd := exec.Command("cp", protoFile, testProtoFile)
		err := cpCmd.Run()

		fileInfo, err := os.Stat(testProtoFile)
		convert.ConvertToSkv2ProtoFile(testProtoFile, fileInfo , err)

		expectedProtoFile := "expected_skv2_result.proto"
		one, err := ioutil.ReadFile(testProtoFile)
		Expect(err).NotTo(HaveOccurred())
		two, err := ioutil.ReadFile(expectedProtoFile)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(one)).To(Equal(string(two)))

		err = os.Remove(testProtoFile)
		Expect(err).NotTo(HaveOccurred())
	})

})
