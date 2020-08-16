# What is command design pattern?

this pattern will encapsulate any request to a standalone obj, which can be trasfare to diffrent object

This trasformation( which handle by Command) lets you to do more operation like:
	* logging a request
	* delay or queue a request
	* reversable operation(like undo, rollback, etc)
	* and many more operation

## How to implement?

1. Create Receiver class(or server side components)

2. Create Command and make receivers just accept values from Command

3. Create concrete command for diffrent sender

4. Create sender(/client) and connect every sender with it's specific concrete command


