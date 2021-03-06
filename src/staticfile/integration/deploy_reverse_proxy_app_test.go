package integration_test

import (
	"github.com/cloudfoundry/libbuildpack/cutlass"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("deploy a staticfile app", func() {
	var app *cutlass.App
	AfterEach(func() {
		if app != nil {
			app.Destroy()
		}
		app = nil
	})

	BeforeEach(func() {
		app = cutlass.New(Fixtures("reverse_proxy"))
		app.Buildpacks = []string{"staticfile_buildpack"}
		PushAppAndConfirm(app)
	})

	It("proxies", func() {
		Expect(app.GetBody("/intl/en/policies")).To(ContainSubstring("Google Privacy Policy"))

		By("hides the nginx.conf file", func() {
			Expect(app.GetBody("/nginx.conf")).To(ContainSubstring("<title>404 Not Found</title>"))
		})
	})
})
