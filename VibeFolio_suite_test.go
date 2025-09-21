package VibeFolio_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestVibeFolio(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "VibeFolio Suite")
}
