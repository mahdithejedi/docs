# [What are meta classes?](https://stackoverflow.com/a/6581949/9651641)

typiclly everything in python is a object
No matter is it a `int` or `str` or even *`class`*

So when everything is a object there should be a object which created this objects like `class`

in the class the object that create class is **`metaclass`**

**metaclass** is a classs' of a class

by default `type` create classes in python

## How metaclass inherit or works

*   **First search for `metaclass` in class creation**

*   **Then search for `metaclass` in module**
  <small> you can set `__metaclass__` var inside module</small>

*   **Then search for `metaclass` in the father of the class**

*   **If non of above happend it user `type` to create class**

