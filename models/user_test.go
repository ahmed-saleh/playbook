package models_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

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
		"password":     "testingPassword1",
	}
	inValidData := map[string]interface{}{
		"display_name": "mr. John Doe",
	}

	BeforeEach(func() {
		// c.Setup("../ini/test.ini")
	})

	Describe("Adding User", func() {
		defer GinkgoRecover()

		sqlDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		Expect(err).To(BeNil())
		var error error

		models.DB, error = gorm.Open(mysql.New(mysql.Config{
			Conn:                      sqlDB,
			SkipInitializeWithVersion: true,
		}), &gorm.Config{})

		Expect(error).To(BeNil())

		Context("User data must be valid", func() {
			It("it will commit to db successfully", func() {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`uuid`,`display_name`,`email`,`password`) VALUES (?,?,?,?,?,?,?)").
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()

				models.AddUser(validData)
				err := mock.ExpectationsWereMet()
				Expect(err).NotTo(HaveOccurred())
			})

			It("it will return err", func() {
				res := models.AddUser(inValidData)
				Expect(res).Should(HaveOccurred())
				Expect(res).To(ContainSubstring("email"))
			})
		})
	})
})
