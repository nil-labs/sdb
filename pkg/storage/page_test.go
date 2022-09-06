package storage_test

import (
	"time"

	"github.com/nil-labs/sdb/pkg/storage"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gmeasure"
)

var _ = Describe("Pages I/O", func() {

	page := storage.NewPage()
	page.Data()[1] = byte(1) // change a byte
	It("should be a able to write Pages", func() {
		Expect(mngr.WritePage(page)).NotTo(HaveOccurred())
	})

	page2 := storage.NewPage()
	It("should be able to read back Pages", func() {
		Expect(mngr.ReadPage(page2)).NotTo(HaveOccurred())
		Expect(page2.Data()[1]).To(Equal(byte(1)))
	})

	It("writing pages should take less than 900 µs", Serial, Label("measurement"), func() {
		experiment := gmeasure.NewExperiment("Writing Sequential Pages")
		experiment.Sample(func(idx int) {
			experiment.MeasureDuration("writing", func() {
				mngr.WritePage(page)
			})
		}, gmeasure.SamplingConfig{N: 500, Duration: time.Minute, NumParallel: 1})
		AddReportEntry(experiment.Name, experiment)
		writingStats := experiment.GetStats("writing")
		medianDuration := writingStats.DurationFor(gmeasure.StatMedian)
		Expect(medianDuration).To(BeNumerically("~", time.Microsecond, 900*time.Microsecond))
	})
})
