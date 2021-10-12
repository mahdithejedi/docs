# 8 Functions anc Class-Based views
FBV=> Function Base View
<br />
CBV=> Class Base View


## 1 FVB or CBV?
you can use CBV for views that need to be subclasses and can be done easily with mixins otherwise you can use FVB

## 2 Keep View Logic Out of URLConfs

## 3 Stick to Loose Coupling in URLConfs
remember that we should keep the code *Decouple* look at [this example](https://github.com/feldroy/two-scoops-of-django-3.x/blob/master/code/chapter_08_example_01.py):
```python
# Don't do this!
from django.urls import path
from django.views.generic import DetailView

from tastings.models import Tasting

urlpatterns = [
    path('<int:pk>',
        DetailView.as_view(
            model=Tasting,
            template_name='tastings/detail.html'),
        name='detail'),
    path('<int:pk>/results/',
        DetailView.as_view(
            model=Tasting,
            template_name='tastings/results.html'),
        name='results'),
]
```
hare we have models, views and urls all coupled together but it's not good, what is we want to add authentication or do any change? it's better to separate urls and views look at example of a [view.py](https://github.com/feldroy/two-scoops-of-django-3.x/blob/master/code/chapter_08_example_02.py) and it's related [urls.py](https://github.com/feldroy/two-scoops-of-django-3.x/blob/master/code/chapter_08_example_03.py):
```python
## Do this!
## views.py
from django.urls import reverse
from django.views.generic import ListView, DetailView, UpdateView

from .models import Tasting

class TasteListView(ListView):
    model = Tasting

class TasteDetailView(DetailView):
    model = Tasting

class TasteResultsView(TasteDetailView):
    template_name = 'tastings/results.html'

class TasteUpdateView(UpdateView):
    model = Tasting

    def get_success_url(self):
        return reverse('tastings:detail',
            kwargs={'pk': self.object.pk})
```
<br />

```python
## DO this!
## urls.py
from django.urls import path

from . import views

urlpatterns = [
    path(
        route='',
        view=views.TasteListView.as_view(),
        name='list'
    ),
    path(
        route='<int:pk>/',
        view=views.TasteDetailView.as_view(),
        name='detail'
    ),
    path(
        route='<int:pk>/results/',
        view=views.TasteResultsView.as_view(),
        name='results'
    ),
    path(
        route='<int:pk>/update/',
        view=views.TasteUpdateView.as_view(),
        name='update'
    )
]
```
there is some advantages over this separation:
it's DRY, loosed coupled, URLConfs do one thins and do it well, Views are class-base, we have infinite flexibility

### 3.1 What if you use FVB?
it's the same, URL should be separated from views

## 4 User URL Namespaces
for example you have an app called `users` and a view called `details` instead of calling is `users_details` use django  [namespaces](https://docs.djangoproject.com/en/3.2/topics/http/urls/#url-namespaces) feature for example:
```python
 ## urls.py at the root of project
 urlpatterns += [
     path('users/', include('users.urls'),
        namespace='users'
     )
 ]
 ```
 and related view:
 ```python
 ## view.py
 class UserUpdateView(UpdateView):
     model= User
     def get_success_url(self):
         return reverse('users:details',
                kwargs={'pk': self.object.pk}
         )
```
and inside HTML template
```html
{% extends 'base.html' %}

{% block title %}Users{% endblock title %}

{% block content %}
<ul>
  {% for user in users %}
    <li>
      <a href="{% url 'users:detail' user.pk %}">{{ user.title }}</a>
      <small>
        (<a href="{% url 'users:update' user.pk %}">update</a>)
      </small>
    </li>
  {% endfor %}
</ul>
{% endblock content %}
```
but why it's user full?

### 4.1 Make for Shorter, More Intuitive, and Don't Repeat Yourself URL Names
<!-- TODO -->
it's more intuitive to use this way!!!


### 4.2 Increase Interoperability With Third-Party Libraries
Consider that you add new *user* third-party app to manage your users beside your own functionality, that will happen if that app use `users_detail` to get users details? it will collapse! your app have `users_details` and third party app to!
how to use it?

```python
# urls.py
urlpatterns += [
    path('users/', include('users.urls'), namespace='users'),
    path('usersmanager/', include('usermanager'), namespace='usersmanager') # third party app
]
```
by this way we can prevent collision between our app the third party application

```html
{% extends "base.html" %}
{% block title %}Users{% endblock title %}
{% block content %}
<p>
  <a href="{% url 'usermanager:create' %}">Create user</a>
</p>
<p>
  <a href="{% url 'users:details' %}">Users details</a>
</p>
{% endblock content %}
```
### 4.3 Easier Searches, Upgrades, and Refactors
for example you have a URL name `users_details`, when you search for `<APPNAME>_details` the result is what? variable? function name? or a URL name?
how can you filter out URL names? so it's better to use _\<APPNAME\>:details_ like URL name to make it more intuitive.

### 4.4 Allows for More APP and Template Reverse
by this way you can have better debugging and adding better modules


## 5 Try to Keep Business Logic Out of Views
it's better to move logic out of view so if you want to generate PDF or add a SOAP API or etc it won't be pain in the ass! =)

## 6 Django Views Are Functions
in Django CBVs have a function which is `as_view()`, this return a callable instance of the view.In other words: CBVs are manage like FBVs

## 7 Don't Use `locals()` as Views content
```python
# DONT DO THIS
def user_display(request, user_id):
    user = get_object_or_404(User, id=user_id)
    date = timezone.now()
    return render(
        request, 'user_display.html', locals()
    )
```
It's bad because
<br />
1- Why should html know about view locals? it's fuck up decoupling concept!
<br />
2- it's hard to maintain! inside html which variable should we user? by which name?
<br />
3- if any variable change you're gonna mess things up
<br />
So please for god sake **DON'T USE `locals()` AS VIEWS CONTEXT**
