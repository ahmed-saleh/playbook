package app_test

import (
	"net/http/httptest"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ahmed-saleh/playbook/config"
	"github.com/ahmed-saleh/playbook/models"
	"github.com/ahmed-saleh/playbook/pkg/app"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "App Suite")
}

var _ = Describe("Server", func() {

	Describe("API responders", func() {
		Context("Response should be retunred correctly", func() {

			It("should respond without any error", func() {
				c, _ := gin.CreateTestContext(httptest.NewRecorder())
				res := app.Gin{C: c}
				res.Response(400, 400, "error validation")
			})
		})
	})
})

var _ = Describe("JWT", func() {
	defer GinkgoRecover()

	BeforeEach(func() {
		config.Setup("../../ini/test.ini")
	})

	Context("Creatng JWT", func() {
		It("should create correctly", func() {
			u := &models.User{}
			u.Uuid = uuid.New().String()

			token, err := app.CreateJwtAccessToken(u.Uuid)
			Expect(err).NotTo(HaveOccurred())
			Expect(token).NotTo(BeEmpty())
			c := &app.Claims{}
			t, e := jwt.ParseWithClaims(token, c, func(token *jwt.Token) (interface{}, error) {
				return "Invalid Key", nil
			})

			Expect(e).To(HaveOccurred())
			Expect(t.Valid).To(BeFalse())

			t, e = jwt.ParseWithClaims(token, c, func(token *jwt.Token) (interface{}, error) {
				return config.AppSetting.JwtSecret, nil
			})

			Expect(e).To(HaveOccurred())
			Expect(t.Valid).To(BeTrue())
		})
	})
})
