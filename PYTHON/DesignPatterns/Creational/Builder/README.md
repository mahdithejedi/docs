#  What is builder and where we use it?


**The director class defines the order in which to execute the building steps, while the builder provides the implementation for those steps.**

Having a director class in your program isn’t strictly necessary. You can always call the building steps in a specific order directly from the client code. However, the director class might be a good place to put various construction routines so you can reuse them across your program.

![Builder structure](https://refactoring.guru/images/patterns/diagrams/builder/structure.png?id=fe9e23559923ea0657aa")

**Use the Builder pattern when you want your code to be able to create different representations of some product (for example, stone and wooden houses).**

Builder is work when we have the objects with same family so instead of make couple same object 
**We decouplicate it**

we make a base class do the same stuff there and extend it in the base classes

what acctually happend here is in Builder design pattern focues is on create complex object step by step  <small>[source](https://stackoverflow.com/a/757773/9651641)</small>

[It’s especially useful when you need to create an object with lots of possible configuration options.](https://refactoring.guru/design-patterns/builder/python/example)

### Examples

[github faif design pattern](https://github.com/faif/python-patterns/blob/master/patterns/creational/builder.py)
[refactoring](https://refactoring.guru/design-patterns/builder/python/example#example-0--main-py)
[geekforgeeks](https://www.geeksforgeeks.org/builder-method-python-design-patterns/)