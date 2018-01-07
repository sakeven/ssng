package utils

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestAll(t *testing.T) {
	suite.Run(t, new(CacheSuite))
	suite.Run(t, new(AddressSuite))
}
