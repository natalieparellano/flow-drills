package moves_test

import (
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/natalieparellano/flow-drills/moves"
)

var _ = Describe("validate", func() {
	var input []byte
	Context("when given valid yaml", func() {
		Context("when read from a file", func() {
			It("returns an array of moves", func() {
				input, err := ioutil.ReadFile("moves.yml")
				Expect(err).ToNot(HaveOccurred())

				_, err = moves.Validate(input)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		It("returns an array of moves", func() {
			input = []byte(`---
- name: flower
  properties:
  - type: timing
    values: null
  - type: direction
    values: null
  variations: [stall, turn]
`)
			res, err := moves.Validate(input)
			Expect(err).ToNot(HaveOccurred())

			move := res[0]
			Expect(move.Name).To(Equal("flower"))

			timing := move.Properties[0]
			Expect(timing.Values).To(Equal(moves.ValidValues("timing")))

			variations := move.Variations
			Expect(variations).To(Equal([]string{"stall", "turn"}))
		})

		It("overrides default property values with valid inputs", func() {
			input = []byte(`---
- name: flower
  properties:
  - type: timing
    values: ["split"]
  - type: direction
    values: null
`)
			res, err := moves.Validate(input)
			Expect(err).ToNot(HaveOccurred())

			move := res[0]
			Expect(move.Name).To(Equal("flower"))

			timing := move.Properties[0]
			Expect(timing.Values).To(Equal([]string{"split"}))
		})
	})

	Context("when given invalid yaml", func() {
		Context("when property type is invalid", func() {
			It("returns an error", func() {
				input = []byte(`---
- name: flower
  properties:
  - type: some-invalid-property
    values: null
  - type: direction
    values: null
`)
				_, err := moves.Validate(input)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when property value is invalid", func() {
			It("returns an error", func() {
				input = []byte(`---
- name: flower
  properties:
  - type: timing
    values:
		- some-invalid-value
  - type: direction
    values: null
`)
				_, err := moves.Validate(input)
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
