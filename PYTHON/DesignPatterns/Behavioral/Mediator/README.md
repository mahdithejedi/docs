# What is Modiator?

Objects in a system communicate through a Mediator instead of directly with each other.
This reduces the dependencies between communicating objects, thereby reducing coupling. [1](https://github.com/faif/python-patterns/blob/master/patterns/behavioral/mediator.py#L2)


### Example

we can have A chatroom which diffrent users can communicate with each other.
<br />
Instead of user call eachother directly they have a **Mediator** which handle communication between deffrent users.[djangosping example](https://www.djangospin.com/design-patterns-python/mediator/)
