## CRON

```shell
benthos -c ./cron.yaml
```

## Input section

`generate` is an input type which generates messages at specified intervals
`generate.interval` is set to 10 seconds
`generate.mapping` specifies the content and structure of the generated message using a Bloblang mapping.

This proof of concept is simply to show that we can generate new message is generated every 10 seconds.

## Output section

`broker` is an output type which sends messages to multiple outputs based on the specified pattern.
`broker.pattern` is set to `fan_out` which sends messages to all outputs.
`broker.outputs` is an array of outputs to send messages to.

`stdout` is an output type which prints messages to the console.
`http_server` is an output type which serves messages over HTTP.
