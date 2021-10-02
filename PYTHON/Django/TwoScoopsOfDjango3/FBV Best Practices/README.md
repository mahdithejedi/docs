# 9 Best Practices For FVBs

## 1 advantage of FVB
FVBs are a great way tj write less code, prevent confusing CBV inheritance and doing simple views without knowing inheritance and keep it  DRY
<br />
*you can use FVB fo write custom 403, 404 etc error handles*

## 2 Passing HttpRequest Object
the only thins you need for FBV is `HttpRequest` as argument and it's over!. you can write your views with this for example
```python
from django.core.exceptions import PermissionDenied
from django.http import HttpRequest

def user_is_teacher(request: HttpRequest) -> HttpRequest:
    if request.user.can_teach or request.user.is_teacher():
        return request
    raise PermissionDenied
```
or you can add more stuff to `HttpRequest`
```python
from django.core.exceptions import PermissionDenied
from django.http import HttpRequest

def user_is_teacher(request: HttpRequest) -> HttpRequest:
    if request.user.can_teach or request.user.is_teacher():
        request.can_see_teacher_stuff = True
        # inside view you can use this! for example
        # {% if request.user.can_see_teacher_stuff %}
        # {% comment %} view scores of students {% endcomment %}
        # {% endif %}
        return request
    raise PermissionDenied
```

## 3 Decorators Are Sweet
decorators are django *syntactic suger* use if more!

```python
# decorators.py
import functools
def can_see_teacher_stuff(view_func):
    functools.wraps(view_func)
    def new_view(*args, **kwargs):
        request = utils.user_is_teacher(request)
        response = view_func(request, *args, **kwargs)
        return response
    return new_view
```

```python
# view.py
from .decorators import can_see_teacher_stuff

@can_see_teacher_stuff
def view_scores(request):
    pass
```

### 3.1 Be Conservative With Decorators
be careful not to mess up with to many decorators in your view
<br />
try to keep decorators as much less as you can
<br />
watch 
[this video](https://archive.org/details/pyvideo_398___how-to-write-obfuscated-python)
to learn more about this

## 4 additional FVB view
[more on FVB views](https://spookylukey.github.io/django-views-the-right-way/)
    

