Publish Amazon stock ticks to Kafka

```shell
labs-streams publish --to kafka --instrument AMZN
```

Subscribe to Amazon stock ticks via Tyk - WebSocket

```shell
wscat -c "ws://localhost:8080/amzn/instruments/subscribe"
```

Subscribe to Amazon stock ticks via Tyk - SSE

```shell
curl http://localhost:8080/amzn/instruments/stream
```

Publish Google stock ticks to Redis Streams

```shell
labs-streams publish --to redis --instrument GOOGL
```

Subscribe to Google stock ticks via Tyk - WebSocket

```shell
wscat -c "ws://localhost:8080/googl/instruments/subscribe"
```

Subscribe to Google stock ticks via Tyk - SSE

```shell
curl http://localhost:8080/googl/instruments/stream
```

Subscribe to Amazon stock ticks via Tyk - WebSocket

```shell
wscat -c "ws://localhost:8080/amzn/instruments/subscribe"
```

Subscribe to Amazon stock ticks via Tyk - SSE

```shell
curl http://localhost:8080/amzn/instruments/stream
```
