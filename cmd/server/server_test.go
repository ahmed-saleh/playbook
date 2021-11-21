package server_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ahmed-saleh/playbook/cmd/server"
	c "github.com/ahmed-saleh/playbook/config"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Server Suite")
}

var _ = Describe("Server", func() {

	BeforeEach(func() {
		c.Setup("../../ini/test.ini")
	})

	Describe("Server testing", func() {
		Context("Server IP", func() {
			//TODO: add server mocks and testing
			It("should be listening on 9911", func() {
				s := server.Build(c.ServerSettings)
				Expect(s.Addr).To(Equal(":9911"))
			})
		})
	})

})
