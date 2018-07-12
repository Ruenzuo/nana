package emulator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/Ruenzuo/nana/emulator"
)

var _ = Describe("Emulator", func() {
	var (
		emulator    Emulator
		memoryValue uint8
	)

	BeforeEach(func() {
		emulator = *NewEmulator()
	})

	Describe("verifying memory access", func() {
		Context("when writting to ROM", func() {
			BeforeEach(func() {
				memoryValue = emulator.ReadMemory(0x0001)
				emulator.WriteMemory(0x0001, 0xFF)
			})

			It("should not change the memory at that address", func() {
				Expect(emulator.ROM[0x0001]).To(Equal(memoryValue))
			})
		})

		Context("when writting to restricted section", func() {
			BeforeEach(func() {
				memoryValue = emulator.ReadMemory(0xFEA0)
				emulator.WriteMemory(0xFEA0, 0xFF)
			})

			It("should not change the memory at that address", func() {
				Expect(emulator.ROM[0xFEA0]).To(Equal(memoryValue))
			})
		})

		Context("when writting to echo", func() {
			BeforeEach(func() {
				emulator.WriteMemory(0xE000, 0xFF)
			})

			It("should write to echo address", func() {
				Expect(emulator.ROM[0xE000]).To(Equal(uint8(0xFF)))
				Expect(emulator.ROM[0xC000]).To(Equal(uint8(0xFF)))
			})
		})

		Context("when writing to unrestricted section", func() {
			BeforeEach(func() {
				emulator.WriteMemory(0xFFFF, 0xFF)
			})

			It("should write to address", func() {
				Expect(emulator.ROM[0xFFFF]).To(Equal(uint8(0xFF)))
			})
		})
	})
})