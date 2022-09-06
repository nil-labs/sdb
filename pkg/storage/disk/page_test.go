package disk_test

import (
	"io/ioutil"
	"os"

	"github.com/nil-labs/sdb/pkg/storage/disk"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Page I/O", func() {
	var db *os.File
	var err error
	var mngr *disk.Manager
	BeforeEach(func() {
		db, err = ioutil.TempFile(os.TempDir(), "*.sdb")
		Expect(err).NotTo(HaveOccurred())
		mngr, err = disk.ManagerFromFile(db)
		Expect(err).NotTo(HaveOccurred())
	})
	AfterEach(func() {
		os.Remove(db.Name())
	})
	Describe("Writing", func() {
		It("should flush the content of the page in the underlying file", func() {
			Expect(mngr.WritePage(disk.PageId(0), disk.Page{})).NotTo(HaveOccurred())
		})
	})
})
