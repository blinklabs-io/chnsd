module github.com/blinklabs-io/chnsd

go 1.19

require (
	github.com/blinklabs-io/snek v0.0.0-00010101000000-000000000000
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/miekg/dns v1.1.55
	go.uber.org/zap v1.24.0
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/blinklabs-io/gouroboros => ../gouroboros

replace github.com/blinklabs-io/snek => ../snek

require (
	github.com/blinklabs-io/gouroboros v0.47.0 // indirect
	github.com/fxamacker/cbor/v2 v2.4.0 // indirect
	github.com/jinzhu/copier v0.3.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/crypto v0.11.0 // indirect
	golang.org/x/mod v0.9.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/tools v0.7.0 // indirect
)
