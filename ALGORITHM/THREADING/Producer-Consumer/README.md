# What is Producer-Consumer Algorithm?
In the first version of the problem, there are two cyclic processes, a producer and a consumer, which share a common, fixed-size buffer used as a queue. The producer repeatedly generates data and writes it into the buffer. The consumer repeatedly reads the data in the buffer, removing it in the course of reading it, and using that data in some way. In the first version of the problem, with an unbounded buffer, **the problem is how to design the producer and consumer code so that, in their exchange of data, no data is lost or duplicated, data is read by the consumer in the order it is written by the producer, and both processes make as much progress as possible**
[wikipedia](https://en.wikipedia.org/wiki/Producer%E2%80%93consumer_problem)
<br />
[Producer Consumer Problem using Semaphores | Set 1(geekforgeeks)](https://www.geeksforgeeks.org/producer-consumer-problem-using-semaphores-set-1/)


