# Golang In-Memory Pub-Sub Package
You can import this package in your project and use the publish-subscribe patterns for different topics
```go
    import (
        "github.com/guptswayam/pubsub"
    );
```

### Creating Broker, Adding Subscribers, Subscribing patterns and Attaching Handler for patterns
```go
    // Creating broker
    broker := pubsub.NewBroker("test-broker")     // "test-broker" is broker name

    // Adding Subscribers
    s1 := broker.AddSubscriber()
    s2 := broker.AddSubscriber()

    // Subscribing pattern
    pattern := "test"
    broker.Subscribe(s1, pattern)
    broker.Subscribe(s2, pattern)

    // Attaching Handler to pattern
    broker.SetHandler(pattern, func(s *pubsub.Subscriber, msg *pubsub.Message) {
        // ...
    });
```

### Publishing Messages
```go
    broker.Publish(pattern, "Hi Subsvribers!")
```

### Example
1. For complete example, please check `pubsub_test.go` file at the root folder of repository


### Resources
1. [Pub-Sub Pattern](https://blog.logrocket.com/building-pub-sub-service-go/)
2. [Registering Go Package](https://go.dev/doc/modules/publishing)