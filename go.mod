module github.com/elga-io/borzoi

go 1.14

require (
	github.com/dgraph-io/badger/v3 v3.2103.2
	github.com/elga-io/canideos v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/rs/zerolog v1.26.1 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/crypto v0.0.0-20220315160706-3147a52a75dd
	gorm.io/gorm v1.23.3
)

replace github.com/elga-io/canideos => ../canideos
