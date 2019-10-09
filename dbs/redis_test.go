package dbs_test

import (
	"errors"
	"testing"

	"github.com/thiagozs/test-mock/services"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"github.com/thiagozs/test-mock/mocks"
)

type DispatchServiceTestSuite struct {
	suite.Suite
	dispRepo  *mocks.MockDispatchRepository
	underTest services.DispatchServices
}

func TestDispatchServiceSuite(t *testing.T) {
	suite.Run(t, new(DispatchServiceTestSuite))
}

func (suite *DispatchServiceTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()

	suite.dispRepo = mocks.NewMockDispatchRepository(mockCtrl)
	suite.underTest = services.NewDispatchServices(suite.dispRepo)
}

func (suite *DispatchServiceTestSuite) TestPing() {
	suite.dispRepo.EXPECT().Ping().Return(nil)
	err := suite.underTest.Ping()
	suite.NoError(err, "Shouldn't error")
}

func (suite *DispatchServiceTestSuite) TestPingErr() {
	suite.dispRepo.EXPECT().Ping().Return(errors.New("Failconn"))
	err := suite.underTest.Ping()
	suite.EqualError(err, "Failconn", "Should be a error")
}

func (suite *DispatchServiceTestSuite) TestQueueSet() {
	suite.dispRepo.EXPECT().QueueSet("worker", `{"name":"thiagozs"}`).Return(nil)
	err := suite.underTest.QueueSet("worker", `{"name":"thiagozs"}`)
	suite.NoError(err, "Shouldn't error")
}

func (suite *DispatchServiceTestSuite) TestQueueSetErr() {
	suite.dispRepo.EXPECT().QueueSet("worker", `{"name":"thiagozs"}`).Return(errors.New("Failconn"))
	err := suite.underTest.QueueSet("worker", `{"name":"thiagozs"}`)
	suite.EqualError(err, "Failconn", "Should be a error")
}

func (suite *DispatchServiceTestSuite) TestQueueGet() {
	suite.dispRepo.EXPECT().QueueGet("worker").Return("ok", nil)
	str, err := suite.underTest.QueueGet("worker")
	suite.NoError(err, "Shouldn't error")
	suite.Equal(str, "ok", "Should be ok")
}

func (suite *DispatchServiceTestSuite) TestQueueGetErr() {
	suite.dispRepo.EXPECT().QueueGet("worker").Return("", errors.New("Failconn"))
	str, err := suite.underTest.QueueGet("worker")
	suite.EqualError(err, "Failconn", "Should be a error")
	suite.Equal(str, "", "Should be empty")
}

// QueueRangeList(queue string, start, end int64) ([]string, error)
// QueueTrim(queue string, start, end int64) (string, error)
// QueueGetList(queue string, amount int) ([]string, error)
// QueueSize(queue string) (int, error)
// Incr(key string) (int, error)
// Decr(key string) (int, error)
// Del(key string) (int, error)
// Get(key string) (string, error)
