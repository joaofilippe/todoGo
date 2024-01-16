package helpers

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

var a int

type CheckUUIDSuite struct {
	suite.Suite
	uuids []uuid.UUID
}

func (s *CheckUUIDSuite) SetupSuite() {
	fmt.Println("Setup Suite >>>>")
	a = 1
	fmt.Println("a:", a)

	s.uuids = []uuid.UUID{
		uuid.New(),
		uuid.Nil,
		uuid.UUID{},
	}
}

func (s *CheckUUIDSuite) TearDownSuite() {
	fmt.Println("a:", a)	
	fmt.Println("<<<< TearDownSuite")
}

func (s *CheckUUIDSuite) SetupTest() {
	fmt.Println("SetupTest ~~~~")
	a ++
	fmt.Println("a:", a)

	s.uuids = []uuid.UUID{
		uuid.New(),
		uuid.Nil,
		uuid.UUID{},
	}
}

func (s *CheckUUIDSuite) TearDownTest() {
	fmt.Println("a:", a)	
	fmt.Println("~~~~ TearDownTest")
}


func (suite *CheckUUIDSuite) Test_IsEmpty() {
	fmt.Println("Test_IsEmpty")
	suite.True(IsEmpty(suite.uuids[1]))
	suite.True(IsEmpty(suite.uuids[2]))
	suite.False(IsEmpty(suite.uuids[0]))
}

func Test_CheckUUID(t *testing.T) {
	suite.Run(t, new(CheckUUIDSuite))
}
