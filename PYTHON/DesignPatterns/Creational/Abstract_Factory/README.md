# What is Abstract Factory?

[Abstract Factory defines an interface for creating all distinct products but leaves the actual product creation to concrete factory classes. Each factory type corresponds to a certain product variety.

The client code calls the creation methods of a factory object instead of creating products directly with a constructor call (new operator). Since a factory corresponds to a single product variant, all its products will be compatible.][1]


# Where to use?

we use when we dont know what is our actual inputs

so we set some behavior so we can use this class for later useage wihtout specify actual inputs and contents


[If an application is to be portable, it needs to encapsulate platform dependencies. These "platforms" might include: windowing system, operating system, database, etc](https://sourcemaking.com/design_patterns/abstract_factory)

## Examples

[faif/python-pattern](https://github.com/faif/python-patterns/blob/master/patterns/creational/abstract_factory.py)
[geeksforgeeks](geeksforgeeks.org/abstract-factory-method-python-design-patterns/)


[1][https://refactoring.guru/design-patterns/abstract-factory/python/example]
