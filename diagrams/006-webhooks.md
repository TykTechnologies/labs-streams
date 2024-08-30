```mermaid
sequenceDiagram
    participant Developer
    participant Tyk Streams (API Gateway)
    participant Event Broker
    participant Webhook URL

    Developer->>+Tyk Streams (API Gateway): Registers Webhook URL for Event Subscription
    Tyk Streams (API Gateway)->>Tyk Streams (API Gateway): Stores Webhook URL for Future Notifications

    Tyk Streams (API Gateway)->>+Event Broker: Subscribes to Event Topics
    Event Broker-->>Tyk Streams (API Gateway): Sends Events

    alt Event Occurs
        Tyk Streams (API Gateway)->>+Webhook URL: Sends Event Notification
        Webhook URL-->>-Tyk Streams (API Gateway): Acknowledges Receipt (Optional)
    end
```
