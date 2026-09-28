package models_test

import (
	"code.cloudfoundry.org/bbs/models"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CheckDefinition", func() {
	Describe("Check Validate", func() {
		Context("with an http check", func() {
			It("accepts the HTTP1 (default) http_version", func() {
				check := models.Check{
					HttpCheck: &models.HTTPCheck{Port: 8080, HttpVersion: models.HTTPCheck_HTTP1},
				}
				Expect(check.Validate()).To(Succeed())
			})

			It("accepts the HTTP2 http_version", func() {
				check := models.Check{
					HttpCheck: &models.HTTPCheck{Port: 8080, HttpVersion: models.HTTPCheck_HTTP2},
				}
				Expect(check.Validate()).To(Succeed())
			})

			It("rejects an unknown http_version", func() {
				check := models.Check{
					HttpCheck: &models.HTTPCheck{Port: 8080, HttpVersion: models.HTTPCheck_HTTPVersion(99)},
				}
				err := check.Validate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("http_version"))
			})
		})

		Context("with a tcp check", func() {
			It("succeeds and never inspects http_version", func() {
				check := models.Check{
					TcpCheck: &models.TCPCheck{Port: 8080},
				}
				Expect(check.Validate()).To(Succeed())
			})
		})
	})
})
