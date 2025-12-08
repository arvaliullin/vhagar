module github.com/arvaliullin/vhagar

go 1.21.10

require (
	gopkg.in/yaml.v3 v3.0.0
	github.com/streadway/amqp v1.0.0
	github.com/lib/pq v1.10.9
	github.com/golang/mock v1.6.0
)

replace gopkg.in/yaml.v3 => /usr/share/gocode/src/gopkg.in/yaml.v3
replace github.com/streadway/amqp => /usr/share/gocode/src/github.com/streadway/amqp
replace github.com/lib/pq => /usr/share/gocode/src/github.com/lib/pq
replace github.com/golang/mock => /usr/share/gocode/src/github.com/golang/mock
