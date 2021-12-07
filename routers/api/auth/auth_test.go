package auth_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ahmed-saleh/playbook/routers"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Auth api Suite")
}

type header struct {
	Key   string
	Value string
}

func performRequest(r http.Handler, method, path string, headers ...header) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

var _ = Describe("Login", func() {
	defer GinkgoRecover()

	r := gin.New()

	BeforeEach(func() {
		routers.InitRouter(r)
	})

	Context("Login attempt", func() {
		It("Login will succeed and JTW will be returned", func() {
			w := performRequest(r, "POST", "/auth/login")
			Expect(w.Code).Should(Equal(http.StatusOK))
			// Expect(w.Body).Should(HaveField("token", w.Body))
		})
	})
})
