package finder_test

import (
	"app/internal/common/finder"
	"fmt"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/suite"
)

type RedisFinderTestSuite struct {
	suite.Suite

	rds *miniredis.Miniredis
}

func (rft *RedisFinderTestSuite) Test_Normal() {
	fds := make([]finder.IFinder, 0)
	expectedAddrs := make([]string, 0)
	N := 5
	for i := 0; i < N; i++ {
		expectedAddrs = append(expectedAddrs, fmt.Sprintf("127.0.0.1:%d", 1000+i))
		fds = append(fds, &finder.RedisFinder{})
		fds[i].SetUp(
			expectedAddrs[i],
			&redis.Options{
				Addr: rft.rds.Addr(),
			})
		fds[i].Start()
	}

	for i := 0; i < N; i++ {
		addrs, err := fds[i].GetAllAddrs()
		rft.Require().Nil(err)
		rft.T().Logf("addrs:%v", addrs)
		rft.Require().ElementsMatch(expectedAddrs, addrs)
	}
}

func TestRedisFinder(t *testing.T) {
	rft := &RedisFinderTestSuite{
		rds: miniredis.RunT(t),
	}
	suite.Run(t, rft)
}
