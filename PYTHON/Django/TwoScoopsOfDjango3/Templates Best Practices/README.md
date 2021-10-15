# 14 Templates Best Practices

## 1 Keep Templates mostly in _templates_
instead of having _templates_ in each app and facing extra nesting , keep all _templates_ in a __templates/__ dir
<br />
there is an exception which you have an third party app which has separate __templates/__ dir for it self

## 2 Template architecture patterns

### 2.1 2-tier
you have views and template 

### 2.2 3-tier
view and two layer of templates.
<br />
you have a basic template by the name of __base.html__ and each app inherit it and have it's base template by the format of __base\_<APP\_NAME>.html__

### 2.3 Flat is better than nested
> Giving you template block as shallow an inheritance as possible will make your template easier to work with and more maintainable
<br />
you reusable template instead of having nested template

## 3 Limit Processing in template
before loop over a query set in template ask this question:
*   1- How large the query set is?
*   2- How large are the objects being retrieved? Are all the fields needed in this template?
*   3- During each iteration of the loop, how much processing occurs?

### 3.2 Filtering With Conditions in Templates
Try your best to prevent any if/else condition in templates.
<br />

### 3.3 Complex Implied Queries in Templates
consider this query:

```html
{# users = Users.objects.all() #}
{% for user in users %}
    <h2> {user.username} </h2>
    <h3> {user.city.code} </h2> <!-- a foriegn key to city and catch code 
        here we will have extra query to database for catching the city tables and get it's related code
    -->
```
in this case it's better to use [`select_related()`](https://docs.djangoproject.com/en/3.2/ref/models/querysets/#select-related) and prevent extra query
```html
{# users = Users.objects.all().select_related('city') #}
```

### 3.4 Hidden CPU Load in Template
In template maybe one line of code isn't a big deal but in fact it's take a long time and complex process for that line of code especially when you work with template tag.
<br />
this can get more serious if you work with template tag image processing . Try to keep logic and any processing out of template to prevent CPU loads


### 3.5 Hidden REST API Calls in Templates
Call REST APIs not by template but by JS
<br />
Use JS to call your REST APIs

## 6 `block.super` Gives the Power of Control
As you may know by using `extends` each block you use will override the parent block for example
```html
<!-- base.html -->
{% block js %}
<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js" integrity="sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM" crossorigin="anonymous"></script>
<script src="https://cdn.jsdelivr.net/npm/@popperjs/core@2.9.2/dist/umd/popper.min.js" integrity="sha384-IQsoLXl5PILFhosVNubq5LC7Qb9DXgDA9i+tQ8Zj3iwWAwPtgFTxbJ8NT4GN1R8p" crossorigin="anonymous"></script>
<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.min.js" integrity="sha384-cVKIPhGWiC2Al4u+LWgxfKTRIcfu0JTxR+EQDz/bgldoEyl4H0zUF0QKbrJ0EcQF" crossorigin="anonymous"></script>
{% endblock %}
```
if you want add a new js when you have to do this
```html
<!-- base_users.html -->
{% extends "base.html" %}
{% block js %}
<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js" integrity="sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM" crossorigin="anonymous"></script>
<script src="https://cdn.jsdelivr.net/npm/@popperjs/core@2.9.2/dist/umd/popper.min.js" integrity="sha384-IQsoLXl5PILFhosVNubq5LC7Qb9DXgDA9i+tQ8Zj3iwWAwPtgFTxbJ8NT4GN1R8p" crossorigin="anonymous"></script>
<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.min.js" integrity="sha384-cVKIPhGWiC2Al4u+LWgxfKTRIcfu0JTxR+EQDz/bgldoEyl4H0zUF0QKbrJ0EcQF" crossorigin="anonymous"></script>
<!-- parent tag content because by jinja override   parent content -->
<script src="https://code.jquery.com/jquery-3.6.0.min.js" integrity="sha256-/xUj+3OJU5yExlq6GSYGSHk7tPXikynS7ogEvDej/m4=" crossorigin="anonymous"></script>
{% endblock %}
```
you can use `{{block.super}}`
```html
{% extends "base.html" %}
{% block js %}
{{block.super}}
<script src="https://code.jquery.com/jquery-3.6.0.min.js" integrity="sha256-/xUj+3OJU5yExlq6GSYGSHk7tPXikynS7ogEvDej/m4=" crossorigin="anonymous"></script>
{% endblock %}
```

## 7 Useful things to consider

### 7.1 Avoid hard-code style into Python
Avoid using _css_, _js_ or other view code inside python files

### 7.2 Common Conventions
* Use dashes over commas in template names
* Include name of tag when closing a tag

### 7.3 Use Implicit and Named Explicit Context Object Properly

### 7.4 Use URL names over Hardcoded path

## 8 Error Page Templates
It's better to serve your error codes via a static file server (eg: Nginx or Apache) or if you use pass services check if you can upload any specific error page. It's because if your django site goes down entirely you can still show error page 
<br />
Have minimalism but good error page

## 9 Follow a Minimalism Approach
it's not good to have a 404 page which there is a layout missing or any other problem