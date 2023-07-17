// TODO: Update the Attrbuition.txt file when adding/removing dependencies

module new-app-service

go 1.15

require (
	github.com/edgexfoundry/go-mod-core-contracts/v2 v2.1.0
	github.com/google/uuid v1.3.0
	github.com/migelankodra/application-export v1.0.0
	github.com/stretchr/testify v1.8.4
)

replace github.com/migelankodra/application-export => ../
