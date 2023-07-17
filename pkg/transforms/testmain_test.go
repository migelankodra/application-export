//
// Copyright (c) 2021 Intel Corporation
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
//

package transforms

import (
	"os"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/interfaces/mocks"
	commonDtos "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/stretchr/testify/mock"

	"github.com/migelankodra/application-export/internal/appfunction"
	"github.com/migelankodra/application-export/internal/bootstrap/container"
	"github.com/migelankodra/application-export/internal/common"

	bootstrapContainer "github.com/edgexfoundry/go-mod-bootstrap/v2/bootstrap/container"
	"github.com/edgexfoundry/go-mod-bootstrap/v2/di"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"
)

var lc logger.LoggingClient
var dic *di.Container
var ctx *appfunction.Context
var mockEventClient *mocks.EventClient

func TestMain(m *testing.M) {
	lc = logger.NewMockClient()

	mockEventClient = &mocks.EventClient{}
	mockEventClient.On("Add", mock.Anything, mock.Anything).Return(commonDtos.BaseWithIdResponse{}, nil)

	dic = di.NewContainer(di.ServiceConstructorMap{
		container.ConfigurationName: func(get di.Get) interface{} {
			return &common.ConfigurationStruct{}
		},
		container.EventClientName: func(get di.Get) interface{} {
			return mockEventClient
		},
		bootstrapContainer.LoggingClientInterfaceName: func(get di.Get) interface{} {
			return lc
		},
	})

	ctx = appfunction.NewContext("123", dic, "")

	os.Exit(m.Run())
}
