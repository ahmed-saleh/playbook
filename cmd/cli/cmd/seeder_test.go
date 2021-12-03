package cmd_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	c "github.com/ahmed-saleh/playbook/cmd/cli/cmd"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Seeder Cmd Suite")
}

var _ = Describe("Seed Registering", func() {

	BeforeEach(func() {
		// c.Setup("../ini/test.ini")
	})

	Describe("Register Seeder", func() {
		defer GinkgoRecover()

		Context("Get the registered seeders", func() {
			It("it will get the correct seeders", func() {
				res := c.GetRegisteredSeeders()
				Expect(res).Should(Equal([]string{"admin_seeder"})) //only one seeder is registered
			})

			// It("it will return err", func() {

			// })
		})
	})
})
