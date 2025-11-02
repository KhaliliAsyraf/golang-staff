package feature

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"goravel/tests"
)

type DepartmentTestSuite struct {
	suite.Suite
	tests.TestCase
}

func TestDepartmentTestSuite(t *testing.T) {
	suite.Run(t, new(DepartmentTestSuite))
}

// SetupTest will run before each test in the suite.
func (s *DepartmentTestSuite) SetupTest() {
}

// TearDownTest will run after each test in the suite.
func (s *DepartmentTestSuite) TearDownTest() {
}

func (s *DepartmentTestSuite) TestIndex() {
	// TODO
	response, err := s.Http(s.T()).WithHeader("Content-Type", "application/json").Post("/users", nil)
	s.Nil(err)

	response.AssertStatus(201)
}
