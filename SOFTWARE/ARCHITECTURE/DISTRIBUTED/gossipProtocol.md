# What is Gossip Protocol?

Gossip protocol is a peer-to-peer communication protocol that allows state sharing in distributed systems. It's based on
the idea of nodes randomly exchanging messages with each other, similar to rumors spreading in a social network

![Gossip protocol](https://media.licdn.com/dms/image/C5612AQGRg4ztrilMrQ/article-inline_image-shrink_1000_1488/0/1520140680132?e=1708560000&v=beta&t=rzenelWSDzwa2XHJKpkVQ-j_mc57Ko0IAmQXXC2Pw3Q)

The gossip protocol is also known as the _epidemic protocol_ because the transmission of the messages is similar to the
way how epidemics spread

## Types Of Gossip Protocol

* anti-entropy model
* rumor-mongering model
* aggregation model

### Anti-entropy model

The anti-entropy algorithm was introduced to reduce the entropy between replicas of a stateful service such as the
database. The replicated data is compared and the difference between replicas are patched. The node with the newest
message shares it with other nodes in every gossip round.

_The anti-entropy model usually transfers the whole dataset resulting in unnecessary bandwidth usage. The techniques
such as checksum, recent update list, and Merkle tree can be used to identify the differences between nodes to avoid
transmission of the entire dataset and reduce network bandwidth usage._ The anti-entropy gossip protocol will send an
unbounded number of messages without termination.

### Rumor-Mongering Gossip Protocol

The rumor-mongering protocol is also known as the dissemination protocol. The rumor-mongering cycle occurs relatively
more frequently than anti-entropy cycles and floods the network with the worst-case load [10]. The rumor-mongering model
utilizes fewer resources such as network bandwidth as only the latest updates are transferred across nodes [8].

A message will be marked as removed after a few rounds of communication to limit the number of messages. There is
usually a high probability that a message will reach all the nodes.

### Aggregation Gossip Protocol

The aggregation gossip protocol computes a system-wide aggregate by sampling information across every node and combining
the values to generate a system-wide value

## Gossip Protocol Performance

The number of nodes that will receive the message from a particular node is known as the fanout. The count of gossip
rounds required to spread a message across the entire cluster is known as the cycle

_cycles necessary to spread a message across the `cluster = O(log n)` to the base of fanout, where n = total number of
nodes_

For instance, it takes approximately 15 gossip rounds to propagate a message across 25,000 nodes. The gossip interval
can be set to a value as low as 10 ms to propagate a message across a big data center in roughly 3 seconds. The
propagation of a message in the gossip protocol should automatically age out to reduce the unnecessary load

## Sources

[highscalability.com](http://highscalability.com/blog/2023/7/16/gossip-protocol-explained.html)
<br />
[wikipedia](https://en.wikipedia.org/wiki/Gossip_protocol)
<br />
[medium](https://medium.com/nerd-for-tech/gossip-protocol-in-distributed-systems-e2b0665c7135)