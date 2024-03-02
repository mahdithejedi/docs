# What is EventSourcing?

The Event Sourcing pattern defines an approach to handling operations on data that's driven by a sequence of events,
each of which is recorded in an append-only store. Application code sends a series of events that imperatively describe
each action that has occurred on the data to the event store, where they're persisted. Each event represents a set of
changes to the data (such as _AddedItemToOrder_).

## some Points about Event-Sourcing:

* Events are simple objects that describe some action that occurred, together with any associated data that's required
  to describe the action represented by the event.
* The append-only storage of events provides an audit trail that can be used to monitor actions taken against a data
  store. It can regenerate the current state as materialized views or projections by replaying the events at any time,
  and it can assist in testing and debugging the system.

* There's some delay between an application adding events to the event store as the result of handling a request, the
  events being published, and the consumers of the events handling them. During this period, new events that describe
  further changes to entities might have arrived at the event store. The system should be designed to account for
  eventual consistency in these scenarios.

* The event store is the permanent source of information, and so the (struct of) event data should never be updated.

* Multithreaded applications and multiple instances of applications might be storing events in the event store. The
  consistency of events in the event store is vital, as is the order of events that affect a specific entity (the order
  that changes occur to an entity affects its current state). Adding a timestamp to every event can help to avoid
  issues. Another common practice is to annotate each event resulting from a request with an incremental identifier.

* **Event publication might be at least once, and so consumers of the events must be idempotent**. They must not reapply
  the update described in an event if the event is handled more than once. Multiple instances of a consumer can maintain
  and aggregate an entity's property, such as the total number of orders placed. Only one must succeed in incrementing
  the aggregate, when an order-placed event occurs. While this result isn't a key characteristic of event sourcing, it's
  the usual implementation decision.

## When to use it?

* When you want to capture intent, purpose, or reason in the data. For example, changes to a customer entity can be
  captured as a series of specific event types, such as Moved home, Closed account, or Deceased.

* When it's vital to minimize or completely avoid the occurrence of conflicting updates to data.

* When you want to record events that occur, to replay them to restore the state of a system, to roll back changes, or
  to keep a history and audit log

* When you use events.

## When **NOT** to use it?

* Small or simple domains, systems that have little or no business logic, or nondomain systems that naturally work well
  with traditional CRUD data management mechanisms.
* Systems where consistency and real-time updates to the views of the data are required.
* Systems where there's only a low occurrence of conflicting updates to the underlying data. For example, systems that
  predominantly add data rather than updating it.

Every change made is represented as an event, and appended to the event log. An entity’s current state can be created by
replaying all the events in order of occurrence. The system information is sourced from the events. As the event has
already happened, it’s always referred to in the past tense (with a past-participle verb, such as _InvoiceSent_).
<br />
Event Sourcing is primarily for _auditing_
<br />