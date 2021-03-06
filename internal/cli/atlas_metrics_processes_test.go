// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cli

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongocli/internal/mocks"
)

const oneMinute = "PT1M"

func TestAtlasMetricsProcess_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProcessMeasurementLister(ctrl)

	defer ctrl.Finish()

	expected := &mongodbatlas.ProcessMeasurements{}

	listOpts := &atlasMetricsProcessOpts{
		host:  "hard-00-00.mongodb.net",
		port:  27017,
		store: mockStore,
	}
	listOpts.granularity = oneMinute
	listOpts.period = oneMinute

	opts := listOpts.newProcessMetricsListOptions()
	mockStore.
		EXPECT().ProcessMeasurements(listOpts.projectID, listOpts.host, listOpts.port, opts).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
