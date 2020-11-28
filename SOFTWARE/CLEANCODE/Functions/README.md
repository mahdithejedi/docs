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


