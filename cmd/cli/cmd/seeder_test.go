package cmd_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	c "github.com/ahmed-saleh/playbook/cmd/cli/cmd"
	"github.com/ahmed-saleh/playbook/models"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Seeder Cmd Suite")
}

var _ = Describe("Seed Registering", func() {

	BeforeEach(func() {
	})

	Describe("Register Seeder", func() {
		defer GinkgoRecover()

		sqlDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		Expect(err).To(BeNil())
		var error error

		models.DB, error = gorm.Open(mysql.New(mysql.Config{
			Conn:                      sqlDB,
			SkipInitializeWithVersion: true,
		}), &gorm.Config{})

		Expect(error).To(BeNil())

		Context("Get the registered seeders", func() {
			It("it will get the correct seeders", func() {
				res := c.GetRegisteredSeeders()
				Expect(res).Should(Equal([]string{"admin_seeder"})) //only one seeder is registered
			})

			It("it will return without an err", func() {
				res, err := c.GetSeeder("admin_seeder")
				Expect(err).To(BeNil())

				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`uuid`,`display_name`,`email`,`password`) VALUES (?,?,?,?,?,?,?)").
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()

				err = res.Seed()
				Expect(err).NotTo(HaveOccurred())

				err = mock.ExpectationsWereMet()
				Expect(err).NotTo(HaveOccurred())
			})

		})
	})
})
