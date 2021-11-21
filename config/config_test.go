package config_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	c "github.com/ahmed-saleh/playbook/config"
)

func TestPlaybook(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}

var _ = Describe("Config", func() {

	BeforeEach(func() {
		c.Setup("../ini/test.ini")
	})

	Describe("Varifying DB setup", func() {
		Context("Database Name", func() {
			It("should be ", func() {
				Expect(c.MysqlSettings.Name).To(Equal("playbook_test_db"))
			})
		})
	})

})
