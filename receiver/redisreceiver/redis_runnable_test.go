// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package redisreceiver

import (
	"context"
	"testing"

	"github.com/open-telemetry/opentelemetry-collector/consumer/consumerdata"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestRedisRunnable(t *testing.T) {
	consumer := &fakeMetricsConsumer{}
	logger, _ := zap.NewDevelopment()
	runner := newRedisRunnable(context.Background(), newFakeClient(), "", consumer, logger)
	err := runner.Setup()
	require.Nil(t, err)
	err = runner.Run()
	require.Nil(t, err)
	// + 6 because there are two keyspace entries each of which has three metrics
	require.Equal(t, len(getDefaultRedisMetrics())+6, len(consumer.md.Metrics))
}

type fakeMetricsConsumer struct {
	md consumerdata.MetricsData
}

func (c *fakeMetricsConsumer) ConsumeMetricsData(_ context.Context, md consumerdata.MetricsData) error {
	c.md = md
	return nil
}
