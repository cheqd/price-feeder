module github.com/ojo-network/price-feeder

go 1.23.8

toolchain go1.23.9

require (
	cosmossdk.io/errors v1.0.2
	cosmossdk.io/math v1.5.3
	github.com/cometbft/cometbft v0.38.12
	github.com/cosmos/cosmos-sdk v0.50.13
	github.com/go-playground/validator/v10 v10.26.0
	github.com/golangci/golangci-lint v1.60.3
	github.com/gorilla/mux v1.8.1
	github.com/gorilla/websocket v1.5.3
	github.com/hashicorp/go-metrics v0.5.4
	github.com/justinas/alice v1.2.0
	github.com/mitchellh/mapstructure v1.5.0
	github.com/ojo-network/ojo v0.4.0-rc1.0.20240912201233-2af60de6026b
	github.com/rs/cors v1.11.1
	github.com/rs/zerolog v1.34.0
	github.com/spf13/cobra v1.9.1
	github.com/spf13/viper v1.19.0
	github.com/stretchr/testify v1.10.0
	golang.org/x/sync v0.14.0
	google.golang.org/grpc v1.71.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/cheqd/cheqd-node v1.4.6-pseudo-version-3.1.5.0.20250508054613-edc1c1cc1b14 // indirect
	github.com/cheqd/cheqd-node/api/v2 v2.3.4 // indirect
)

replace (
	cosmossdk.io/core => cosmossdk.io/core v0.11.1
	github.com/cometbft/cometbft => github.com/cometbft/cometbft v0.38.12
	github.com/confio/ics23/go => github.com/cosmos/cosmos-sdk/ics23/go v0.8.0
	github.com/cosmos/cosmos-sdk => github.com/cheqd/cosmos-sdk v0.50.13-height-mismatch-patched
	// dgrijalva/jwt-go is deprecated and doesn't receive security updates.
	github.com/dgrijalva/jwt-go => github.com/golang-jwt/jwt/v4 v4.4.2
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/osmosis-labs/fee-abstraction/v8 => github.com/cheqd/fee-abstraction/v8 v8.0.3-0.20250415073134-32bfc4118457

	// https://github.com/cheqd/feemarket/tree/cheqd/v0.50.x
	github.com/skip-mev/feemarket => github.com/cheqd/feemarket v1.0.5-0.20250415072337-dc43e8876c3a
)
