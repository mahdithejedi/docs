# What is State Design Pattern?

In state design pattern an object can change it's behaviour when it's internal state changed


Exactly like when you put your phone into Airplain mode.

## When to use it?

*	**Use the State pattern when you have an object that behaves
differently depending on its current state, the number of states
is enormous, and the state-specific code changes frequently.**
<br />
*	**Use State when you have a lot of duplicate code across similar
states and transitions of a condition-based state machine.**
<br />
*	**Use the pattern when you have a class polluted with massive
conditionals that alter how the class behaves according to the
current values of the class’s fields.**
<br />

## How to deploy it?

1. Select a context class or make a new one if state code is distributed across multiple classes
<br />
2. Declare interface class for state(This can lead to mirroring context attributes beacuse we may have state specific atribute)
<br />
3. For every state create concrete class (which driven from interface state class) and move all context code to them
<br />
4. In the context class, add a refrence to state interface
<br />
5. Go over context class and replace deployment with state interface
<br />
6. create an instance to switch class behavoiur


### Resources

https://refactoring.guru/design-patterns/state
<br />


