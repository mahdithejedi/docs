# Functions!

## SMALL !!!
The first rule of functions is that they should be small. The second rule of functions is that they should be smaller than that
<br />
Each function should tell a story about what is happening in the function and each function should led you to the next compelling order
<br />

## Block and Indenting
This implies that in statements such as _if_, _else_, _while_ etc, should be one line which is a function call:

*	this keep enclosing functions small
*	adding documetation value because the function called within the block can have a nicely descriptive name


functions should not be long enough to hold nested structure the indent level of a function should not be greater than one or two. This, of course, makes the functions easier to read and understand


## Do one thing

**FUNCTIONS SHOULD DO ONE THING . THEY SHOULD DO IT WELL .THEY SHOULD DO IT ONLY**
<br />
If a function does only those steps that are one level below the stated name of the function, then the function is doing one thing. After all, the reason we write functions is to decompose a larger concept (in other words, the name of the function) into a set of steps at the next level of abstraction in other way if your function is too large it's doing more than one thing so you should beak it into smaller functions which do one thing.
<br />
So, another way to know that a function is doing more than “one thing” is if you can extract another function from it with a name that is not merely a restatement of its implementation . like if statement , sometimes if statement can no be extracted to another function


## Reading Code from Top to Bottom: The Stepdown Rule
We want the code to read like a top-down narrative, we want every function to be followed by those at next level of abstraction so we can read the program descending,
<br />


## Use Descriptive Names
Don’t be afraid to make a name long. A long descriptive name is better than a short enigmatic name. A long descriptive name is better than a long descriptive comment. Use a naming convention that allows multiple words to be easily read in the function names, and then make use of those multiple words to give the function a name that says what it does.Don’t be afraid to spend time choosing a name. Indeed, you should try several different names and read the code with each in place. Modern IDEs like Eclipse or IntelliJ makeit trivial to change names. Use one of those IDEs and experiment with different names until you find one that is as descriptive as you can make it.
<br />
Choosing descriptive names will clarify the design of the module in your mind and help you to improve it. It is not at all uncommon that hunting for a good name results in a favorable restructuring of the code.
<br />
Be consistent in your names. **Use the same phrases, nouns, and verbs in the function  names you choose for your modules**. Consider, for example, the names _includeSetupAndTeardownPages_ , _includeSetupPages_ , _includeSuiteSetupPage_ , and _includeSetupPage_ . The similar phraseology in those names allows the sequence to tell a story. Indeed, if I showed you just the sequence above, you’d ask yourself: “What happened to _includeTeardownPages_ , _includeSuiteTeardownPage_ , and _includeTeardownPage_?” How’s that for being “. . . pretty much what you expected.”


## Function Arguments
The ideal number of arguments for a function is zero (niladic). Next comes one (monadic), followed closely by two (dyadic). Three arguments (triadic) should be avoided where possible. More than three (polyadic) requires very special justification—and then shouldn’t be used anyway
<br />
there is _ reasons for this:
*	it's hard for testing functions with lots of arguments
*	differentiating the diffrent function with similar input would be hard

as we know arguments are hard, output argumets are harder than that, When we read a function, we are used to the idea of information going in to the function through arguments and out through the return value. We don’t usually expect information to be going out through the arguments. So output arguments often cause us to do a double-take	

## Common Monadic Form
There is two reasons to pass a single argument into a function:
*	you may ask a question about an argument ``bool FileIsExists`
*	transform input to another type and return it in output for eg `openfile` convert a str to a file

remmember there is some Event function with does not return anything for eg: `void passwordAttemptFailedNtimes(int attempts)`


## Flag arguments
Flag arguments are ugly beacuse they do more than one thing, do one thing when they are true and another when they are false, so break them down in two seprate function

## Verbs and Keywords
We can put input type in Function name in the type of verb/noun for example instead of `write(name)` we can use `writeFields(name)` which make it clear that name if a Field or instead of write `assertEquals` we cab have `assertExpectedEqualsActual(expected, actual)`.


