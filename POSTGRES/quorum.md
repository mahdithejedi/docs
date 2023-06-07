[//]: # (__QUORUM__)
[//]: # (__DISTRIBUTES_SYSTEMS__)
[//]: # (__DISTRIBUTESSYSTEMS__)
[//]: # (__REPLICATION__)
[//]: # (__NETWORK__)
[//]: # (__NETWORKING__)
[//]: # (__PROTOCOL__)


**For reading about _PAXOS_ refer to [POSTGRESQL section](../SOFTWARE/PAXOS.md)**
___
## Explanation

A quorum in a distributed system is the minimal number of replicas on which a distributed operation (commit/abort) must
be completed before proclaiming the operation’s success.

**Selection quorum** -> We should select more than half of the replicas in the cluster. If there are  _R_ replicas in the cluster, we should choose _(R/2)+1_
as a quorum number.

**Quorum Rule about write and Read** -> A quorum can be achieved only when the nodes adhere to the methodology  _w+r>n_ where _w_ is the minimum nodes for write operations, _r_ is the minimum nodes for read operations, and _n_ is the total number of nodes in a cluster.
___
## What happens when there is “Split” or “Cluster Partitioning”
Its a tie situation, this can be resolve by giving the the voting members some way to perform a tiebreak, may be by giving a member a “super vote” or by disabling one member vote.
___
## Node handling

**Consider that there are 7 nodes. What will happen when Node6 goes down and come back?**
In a quorum-based replication setup, when a node goes down, the remaining nodes continue to operate and replicate
changes among themselves. When the failed node (in this case, Node 6) comes back online, it needs to catch up with the
changes made by the other nodes during its downtime before it can start accepting new write requests.

The process of syncing the failed node with the changes made by other nodes is called resynchronization or catch-up
replication. The exact process can vary depending on the replication setup, but typically, the resynchronization process
involves the following steps:

1. The failed node (Node 6) connects to one of the other nodes in the quorum and requests a list of changes made since
   its last known state.

2. The other node sends a list of changes, typically in the form of WAL (Write-Ahead Log) records.

3. The failed node applies the changes to its own copy of the database, bringing it up to date with the other nodes.

4. Once the failed node has caught up with the other nodes, it can start accepting new write requests and participating
   in the quorum again.

It's worth noting that the resynchronization process can take some time and may involve a significant amount of data
transfer, depending on the size and complexity of the database and the amount of changes made during the failed node's
downtime. However, once the failed node has caught up with the other nodes, the quorum-based replication can resume
normal operation, ensuring that all nodes have consistent and up-to-date copies of the database.
___
### Sources

[alibaba cloud](https://www.alibabacloud.com/blog/postgresql-multi-node-quorum-based-zero-data-loss-and-ha-failover-switchover-solution_597673)
<br />


___
[Postgresql Documentation](https://www.postgresql.org/docs/current/warm-standby.html#synchronous-replication)
  + [Run time configuration replication](https://www.postgresql.org/docs/current/runtime-config-replication.html)
  + [Async commit](https://www.postgresql.org/docs/current/wal-async-commit.html)
___
  
<br />

[Medium](https://blog.softwaremill.com/quorum-replication-on-postgresql-7dbf2f340cd)
<br />
[What is quorum in distributed systems? <small> educative.io </small>](https://www.educative.io/answers/what-is-quorum-in-distributed-systems)
<br />
    [Medium Quorum in Distributed Systemds](https://medium.com/@sunny_81705/quorum-in-distributed-systems-37cbe17aae88)
___
#### Witness Node

[Witness node ins Quorum Replication](https://subscription.packtpub.com/book/data/9781838984854/1/ch01lvl1sec07/considering-quorum)
<br />
___
### Quorum vs Paxos?
Quorum and Paxos are two related concepts in distributed computing, but they have some important differences.

Quorum is a concept that refers to the minimum number of nodes in a distributed system that must agree on a decision or action for it to be considered valid. In the context of database replication, quorum is often used to ensure that a majority of nodes agree on the state of the database at all times. Quorum-based replication systems typically use a consensus protocol to ensure that a majority of nodes agree on the state of the database.

Paxos, on the other hand, is a specific consensus protocol that is commonly used in distributed systems to ensure that a majority of nodes agree on a decision or action. The Paxos protocol is a leader-based protocol, meaning that it relies on a designated leader node to propose a decision or action to the other nodes in the system. The other nodes then vote on the proposal, and if a majority of nodes agree, the proposal is accepted.

One key difference between quorum and Paxos is that quorum is a general concept that can be implemented using a variety of consensus protocols, while Paxos is a specific consensus protocol. Another difference is that quorum is typically used in the context of database replication, while Paxos can be used in a wide range of distributed computing applications.

In summary, quorum and Paxos are related concepts in distributed computing that both involve ensuring that a majority of nodes agree on a decision or action. However, quorum is a general concept that can be implemented using different consensus protocols, while Paxos is a specific consensus protocol that is commonly used in distributed systems.

___