package disk_test

import (
	"io/ioutil"
	"os"

	"github.com/nil-labs/sdb/pkg/storage/disk"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Manager", func() {
	var db *os.File
	var err error
	BeforeEach(func() {
		db, err = ioutil.TempFile(os.TempDir(), "*.sdb")
		Expect(err).NotTo(HaveOccurred())
	})
	AfterEach(func() {
		os.Remove(db.Name())

	})
	Describe("Creating Manager", func() {
		Context("with existing db and log files", func() {

			It("should succeed", func() {
				Expect(disk.ManagerFromFile(db)).ToNot(BeNil())
			})
		})
		Context("with invalid input", func() {
			It("should return error", func() {
				_, err := disk.ManagerFromFile(nil)
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
