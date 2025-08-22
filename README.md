# backend-patterns

## [Challenge 1: Handling Slow or Unavailable External Services](https://github.com/nima-abdpoor/backend-patterns/tree/Handling-Slow-or-Unavailable-External-Services)

### Overview

This challenge addresses the scenario where the service depends on an external API (e.g., for sending SMS), which can become slow or temporarily unavailable. The goal is to design a resilient system that:

* Remains responsive under high load or temporary external failures.
* Prevents overloading the third-party API with repeated retries.
* Provides proper feedback to users when the service is unavailable.

---

### Approach

We implemented a Circuit Breaker pattern to protect the service from failures and slowdowns of the external SMS API.

## [Challenge 2: Multi-Step Operation Across Independent Services](https://github.com/nima-abdpoor/backend-patterns/tree/Multi-Step-Operation-Across-Independent-Services)

### Overview

This challenge addresses a scenario where multiple decoupled services must work together to complete a single operation, such as a checkout process:

1. Create an order
2. Deduct inventory
3. Process payment

The services are fully independent, and failures in any step must trigger rollback of previous actions to maintain consistency.

---

### Approach

We implemented the Saga Pattern (Orchestration-based) to manage the multistep workflow.


## [Challenge 3: Triggering Side Effects Without Tight Coupling](https://github.com/nima-abdpoor/backend-patterns/tree/Triggering-Side-Effects-Without-Tight-Coupling)

### Scenario
In distributed systems, different services need to **communicate asynchronously** without being tightly coupled.  
An **Event Bus** helps us achieve this by allowing services to **publish events** and other services to **subscribe** to those events.

This challenge is about building a **lightweight in-memory Event Bus** in Go.

---

## Approach

We design the solution around an **interface**:

```go
type EventBus interface {
    Publish(eventName string, payload interface{})
    Subscribe(eventName string, handler func(payload interface{}))
}
```

## [Challenge 4: Keeping Distributed Data in Sync Over Time](https://github.com/nima-abdpoor/backend-patterns/tree/Keeping-Distributed-Data-in-Sync-Over-Time)

### Scenario
User data exists across multiple independent services.  
When a user updates their profile, the changes should **eventually propagate** to all relevant services, even if some are temporarily down.

The goal is to ensure that updates are not lost, and they are retried until successfully delivered.

---

### Approach

To solve this, we use a **simple distributed task queue** pattern.

- When a user profile update occurs, an **update task** is enqueued.
- A **worker** processes tasks from the queue and tries to deliver them to the destination service.
- If the destination service is **temporarily unavailable**, the task is retried until it succeeds.
- This guarantees **eventual consistency** across services.

This design simulates how systems like **Kafka, RabbitMQ, or SQS** work, but implemented in a minimal in-memory way for demonstration.

---

### Flow

1. User updates profile in Service A.
2. The update event is enqueued into the task queue.
3. A worker continuously processes tasks:
    - Calls the target service API.
    - If it fails (service down), the task is re-enqueued with a delay.
4. Once delivery succeeds, the task is marked as done.

---