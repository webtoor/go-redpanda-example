# GO Redpanda Example

This is repository how to use [Redpanda](https://redpanda.com)
 in golang, with [Sarama](https://github.com/Shopify/sarama)

## How to run

```bash
# clone this repo
git clone https://github.com/webtoor/go-redpanda-example.git

cd go-redpanda-example

# running docker compose
docker-compose up -d

# run consumer with command
go run cmd/consumer/main.go

# run producer with command
go run cmd/producer/main.go
```