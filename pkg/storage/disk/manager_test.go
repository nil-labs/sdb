package disk_test

import (
	"io/ioutil"
	"os"

	"github.com/nil-labs/sdb/pkg/storage/disk"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Manager", func() {
	Describe("Creating Manager", func() {
		Context("with existing db and log files", func() {
			var db, log *os.File
			var err error
			BeforeEach(func() {
				db, err = ioutil.TempFile(os.TempDir(), "*.sdb")
				Expect(err).NotTo(HaveOccurred())
				log, err = ioutil.TempFile(os.TempDir(), "*.sdb_log")
				Expect(err).NotTo(HaveOccurred())

			})
			AfterEach(func() {
				os.Remove(db.Name())
				os.Remove(log.Name())
			})
			It("should succeed", func() {
				Expect(disk.ManagerFromFiles(db, log)).ToNot(BeNil())
			})
		})
		Context("with not existing log and db files", func() {
			It("should return error", func() {
				_, err := disk.ManagerFromFiles(nil, nil)
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
