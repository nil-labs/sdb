package storage_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/nil-labs/sdb/pkg/storage"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDisk(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Disk Suite")
}

var db *os.File
var err error
var mngr *storage.Manager

var _ = BeforeSuite(func() {
	db, err = ioutil.TempFile(os.TempDir(), "*.sdb")
	Expect(err).NotTo(HaveOccurred())
	mngr, err = storage.ManagerFromFile(db.Name())
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	Expect(mngr.Close()).NotTo(HaveOccurred())
	Expect(os.Remove(db.Name())).NotTo(HaveOccurred())
})