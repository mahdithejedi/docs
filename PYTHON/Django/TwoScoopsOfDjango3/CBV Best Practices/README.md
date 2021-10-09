# 10 Best Practices for CBVs
as you know django treats CBVs like functions. it calls their `as_view()` function

## 1 Guidelines
* Less view code is better
* Never Repeat code in views
* Views should handle presentation logic, keep business logic in models or forms if you can
* Keep your views simpler
* Keep your mixins simper

it's worth it to take a look at ccbv.co.uk

## 2 Use Mixins with CBVs
* The base view classes provided by Django always go to the **right**
* Mixins go to the left of the base view
* Mixins should not inherit from any other class. Keep you inheritance chain simple!

## 3 Which Django GCBV should be used to what task?

| Name        | Purpose     |
| ----------- | ----------- |
| [`View`](https://docs.djangoproject.com/en/3.2/ref/class-based-views/base/#django.views.generic.base.View)      | Base view       |
| [`RedirectView`](https://docs.djangoproject.com/en/3.2/ref/class-based-views/base/#redirectview)  | redirect user to another URL        |
| [`TemplateView`](https://docs.djangoproject.com/en/3.2/ref/class-based-views/base/#templateview) | Display a Django HTML template |
| [`ListView`](https://docs.djangoproject.com/en/3.2/ref/class-based-views/generic-display/#listview) | List Objects
| [`DetailView`](https://docs.djangoproject.com/en/3.2/ref/class-based-views/generic-display/#detailview) | Display on Object |
| [`FormView`](https://docs.djangoproject.com/en/3.2/ref/class-based-views/generic-editing/#django.views.generic.edit.FormView) | Submit a form |
| [`CreateView`](https://docs.djangoproject.com/en/3.2/ref/class-based-views/generic-editing/#django.views.generic.edit.CreateView) | Create on object |
| [`UpdateView`](https://docs.djangoproject.com/en/3.2/ref/class-based-views/generic-editing/#updateview) | Update an Object |
| [`DeleteView`](https://docs.djangoproject.com/en/3.2/ref/class-based-views/generic-editing/#django.views.generic.edit.DeleteView) | Delete an object |
| Generic date View | For display of objects that occur over a rang of time |


## 4. General Tips for Django CBVs
<br />

### 4.1 Constraining Django CBV/GCBV Access to Authentication User
Django have a decorator for authenticating use `django.contrib.auth.decorators.login_required` but this decorator won't keep logic out of view instead we can use django `django.contrib.auth.mixins.LoginRequiredMixin`.


**If you want to user `LoginRequiredMixin` and using `dispatch` method keep in mind to call ***`super().dispatch(request, *args, **kwargs)`*** method**

### 4.2 Performing Custom Actions in View with Valid forms
use `form_valid` method when you use a form in which case form is valid. you should return `django.http.HttpResonseRedirec`

### 4.3 Invalid Form
if form is invalid use `form_invalid` method

## 6. Using Just django.views.generic.View
There is two advantages over using `generic.View` and FBVs

*   Get rid of extra `if request.method == ...` and use functions
*   You have  the power of ObjectOriented!


