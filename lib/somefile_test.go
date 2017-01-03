package lib_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/zet4/go-travis-docker-test/lib"
)

var _ = Describe("Lib", func() {
	Describe("Counter stuff", func() {
		Context("Counter is zero", func() {
			It("should return 0 after first call", func() {
				val := lib.IncrementAndReturn()
				Expect(val).To(Equal(uint64(0)))
			})
			It("should return 1 after second call", func() {
				val := lib.IncrementAndReturn()
				Expect(val).To(Equal(uint64(1)))
			})
			It("should return 0 after reset and third call", func() {
				lib.Reset()
				val := lib.IncrementAndReturn()
				Expect(val).To(Equal(uint64(0)))
			})
		})
	})
})
