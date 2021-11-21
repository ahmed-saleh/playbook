package models_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	c "github.com/ahmed-saleh/playbook/config"
	"github.com/ahmed-saleh/playbook/models"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Models Suite")
}

var _ = Describe("User", func() {

	validData := map[string]interface{}{
		"email":        "test@email",
		"display_name": "mr. John Doe",
	}
	inValidData := map[string]interface{}{
		"display_name": "mr. John Doe",
	}

	BeforeEach(func() {
		c.Setup("../ini/test.ini")
		models.Setup(c.MysqlSettings)
	})

	Describe("Adding User", func() {
		Context("User data must be valid", func() {
			It("it will return nil", func() {
				res := models.AddUser(validData)
				Expect(res).To(BeNil())
			})

			It("it will return err", func() {
				res := models.AddUser(inValidData)
				Expect(res).Should(HaveOccurred())
				Expect(res).To(ContainSubstring("email"))
			})
		})
	})
})
