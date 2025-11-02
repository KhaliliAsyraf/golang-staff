package feature

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"goravel/tests"
)

type EmployeeTestSuite struct {
	suite.Suite
	tests.TestCase
}

func TestEmployeeTestSuite(t *testing.T) {
	suite.Run(t, new(EmployeeTestSuite))
}

// SetupTest will run before each test in the suite.
func (s *EmployeeTestSuite) SetupTest() {
}

// TearDownTest will run after each test in the suite.
func (s *EmployeeTestSuite) TearDownTest() {
}

func (s *EmployeeTestSuite) TestIndex() {
	// TODO
}
