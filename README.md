# learn-pub-sub-starter (Peril)

This is the starter code used in Boot.dev's [Learn Pub/Sub](https://learn.boot.dev/learn-pub-sub) course.

## Running RabbitMQ locally

The repository ships with a helper script that brings up a RabbitMQ broker with the `rabbitmq_stomp` plugin enabled (the automated tests check for it).

```bash
./rabbit.sh start   # builds the custom image if needed and starts the container
./rabbit.sh logs    # tails the broker logs
./rabbit.sh stop    # stops the container
```

The container exposes the default AMQP port (`5672`) plus the management UI/API (`15672`), so the server in `cmd/server` can connect without additional configuration.
