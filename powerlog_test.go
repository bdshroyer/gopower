package gopower_test

import (
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bdshroyer/gopower"
)

var _ = Describe("Powerlog", func() {
	Describe("NewPowerLogReader", func() {
		Context("PowerLog binary exists", func() {
			It("Creates a new reader with a command", func() {
				pl, err := gopower.NewPowerLogReader()

				Expect(err).NotTo(HaveOccurred())

				workDir, err := os.Getwd()
				Expect(err).NotTo(HaveOccurred())

				Expect(filepath.Base(pl.Cmd.String())).To(Equal("PowerLog -duration 1"))
				Expect(pl.FilePath).To(Equal(filepath.Join(workDir, "PowerLog.csv")))
			})
		})
	})
})
