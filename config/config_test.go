package config_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	c "github.com/ahmed-saleh/playbook/config"
)

func TestConfig(t *testing.T) {
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

	Describe("Verifying JWT Setup", func() {
		Context("JWT time", func() {
			It("should be ", func() {
				Expect(c.AppSetting.JwtTime).To(Equal(60)) //test.ini data is 60
			})
		})
	})

})
