package dbs_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"github.com/thiagozs/test-mock/dbs"
	"github.com/thiagozs/test-mock/mocks"
)

type DispatchServiceTestSuite struct {
	suite.Suite
	dispRepo  *mocks.MockDispatchServices
	underTest dbs.DispatchServices
}

func TestDispatchServiceSuite(t *testing.T) {
	suite.Run(t, new(DispatchServiceTestSuite))
}

func (suite *DispatchServiceTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()

	suite.dispRepo = mocks.NewMockDispatchServices(mockCtrl)
	suite.underTest = dbs.NewRedis()
}

func (suite *DispatchServiceTestSuite) TestPing() {
	suite.dispRepo.EXPECT().Ping().Return(nil)
	err := suite.underTest.Ping()
	suite.NoError(err, "Shouldn't error")
}

func (suite *DispatchServiceTestSuite) TestPingErr() {
	suite.dispRepo.EXPECT().Ping().Return(errors.New("Failconn"))
	err := suite.underTest.Ping()
	fmt.Println(err)
	suite.EqualError(err, "Failconn", "Should be a error")
}
