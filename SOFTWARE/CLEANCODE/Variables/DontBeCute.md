## Don't be cute
*	do not use cultute dependent names like : eatMyShorts for abort
## use one word per concept
*	Pick one word for one abstract concept and stick with it, for instance you can not have `fetch`,`retrive` and `get` as equivalent method for diffrent classes
*	The function names have to stand alone, and they have to be consistent in order for you to pick the correct method without any additional exploration


for consider `UserManager` and `DeviceHandler`, can not both of them be either  _Manager_ or _Handler_? are the really _Handler_ ? the name lead you to expect _Manager_ and _Handler_ are two diffrent concept and they have diffrent in type as well as class


**A consistent lexicon is a great boon to the programmers who must use your code.**

## Dont pun
Using the same term for two different ideas is essentially a pun. If you follow the “one word per concept” rule, you could end up with many classes that have, for example, an _add_ method; in one class add is creating a new value by concatenating two exsiting values, and another class with put a single value to an array, should it consider to be _add_ or _append_ or _insert_ ?

## Use Solution Domain Names
you are programmer and who read the code is programmer too, why not using name of algorithm and patterns? like _AccountVisitor_?


## Add meaningful context 
There are a few names which are meaningful in and of themselves—most are not. Instead, you need to place names in context for your reader by enclosing them in well-named classes, functions, or namespaces


## Don’t Add Gratuitous Context
Shorter names are generally better than longer ones, so long as they are clear. Add no more context to a name than is necessary.
