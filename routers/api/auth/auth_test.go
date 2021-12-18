package auth_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ahmed-saleh/playbook/config"
	"github.com/ahmed-saleh/playbook/models"
	"github.com/ahmed-saleh/playbook/routers"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Auth api Suite")
}

type header struct {
	Key   string
	Value string
}

func performRequest(r http.Handler, method, path string, hd header, payload *bytes.Buffer) *httptest.ResponseRecorder {

	req := httptest.NewRequest(method, path, payload)
	req.Header.Add(hd.Key, hd.Value)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

var _ = Describe("Login", func() {
	defer GinkgoRecover()

	BeforeEach(func() {
		//TODO: use filepath package
		config.Setup("../../../ini/test.ini")
		models.SetupSqli()
	})

	r := gin.New()

	BeforeEach(func() {
		routers.InitRouter(r)
	})

	Context("Login attempt", func() {
		It("Login will succeed and JWT will be returned", func() {
			// h := header{}
			// payload := map[string]string{
			// 	"email":    "test@test.com",
			// 	"password": "asdf",
			// }
			// body, _ := json.Marshal(payload)

			// w := performRequest(r, "POST", "/api/auth/login", h, bytes.NewBuffer(body))
			// Expect(w.Body).Should(Equal(http.StatusOK))
			// Expect(w.Code).Should(Equal(http.StatusOK))
		})
	})
})
