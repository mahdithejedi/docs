# **Queries and the Database Layer**

## 1 Use `get_object_or_404`
if you want to retrieve a single object data use [`get_object_or_404`](https://docs.djangoproject.com/en/3.2/topics/http/shortcuts/#get-object-or-404) inside **view**

***DANGER: you have to use `get_object_or_404` JUST INSIDE VIEW***

## 2 Be Careful With Queries That Might Throw Exceptions
<br />

### 2.1 ObjectDoesNotExists vs DoesNotExists
[`ObjectDoesNotExists`](https://docs.djangoproject.com/en/3.2/ref/exceptions/#objectdoesnotexist) => Any Model
<br />
[`DoesNotExists`](https://docs.djangoproject.com/en/3.2/ref/models/class/#django.db.models.Model.DoesNotExist) => Specific Model
<br />
<small> 
[example](https://github.com/feldroy/two-scoops-of-django-3.x/blob/master/code/chapter_07_example_01.py)
 </small>

### 2.2  When You Just Want One Object but Get Three Back
In this case you should use `MultipleObjectsReturned`
<br />
<small>
[example](https://github.com/feldroy/two-scoops-of-django-3.x/blob/master/code/chapter_07_example_02.py)
</small>

## 3 Use LazyEvaluation to Make Queries Legible
Break queries into multiple lines to make code cleaner and better
Instead of [this code](https://github.com/feldroy/two-scoops-of-django-3.x/blob/master/code/chapter_07_example_03.py):
```python
# Do this!
from django.db.models import Q

from promos.models import Promo

def fun_function(name=None):
    """Find working ice cream promo"""
    results = Promo.objects.active()
    results = results.filter(
                Q(name__startswith=name) |
                Q(description__icontains=name)
            )
    results = results.exclude(status='melted')
    results = results.select_related('flavors')
    return results
``` 

Use [this one](https://github.com/feldroy/two-scoops-of-django-3.x/blob/master/code/chapter_07_example_04.py):
```python
# Do this!
from django.db.models import Q

from promos.models import Promo

def fun_function(name=None):
    """Find working ice cream promo"""
    results = Promo.objects.active()
    results = results.filter(
                Q(name__startswith=name) |
                Q(description__icontains=name)
            )
    results = results.exclude(status='melted')
    results = results.select_related('flavors')
    return results
```

### 3.1 Chaining Queries for Legibility
Use [This format](https://github.com/feldroy/two-scoops-of-django-3.x/blob/master/code/chapter_07_example_05.py) for chaining queries
```python
# Do this!
from django.db.models import Q

from promos.models import Promo

def fun_function(name=None):
  """Find working ice cream promo"""
  qs = (Promo
          .objects
          .active()
          .filter(
              Q(name__startswith=name) |
              Q(description__icontains=name)
          )
          .exclude(status='melted')
          .select_related('flavors')
      )
    return qs
```

## 4 Lean on Advanced Query Tools

### 4.1 Query Expressions
consider [this example](https://github.com/feldroy/two-scoops-of-django-3.x/blob/master/code/chapter_07_example_07.py):
```python
# Don't do this!
from models.customers import Customer

customers = []
for customer in Customer.objects.iterator():
    if customer.scoops_ordered > customer.store_visits:
        customers.append(customer)
```
Use Django [`F`](https://docs.djangoproject.com/en/3.2/ref/models/expressions/) [for example](https://github.com/feldroy/two-scoops-of-django-3.x/blob/master/code/chapter_07_example_08.py):
```python
from django.db.models import F

from models.customers import Customer

customers = Customer.objects.filter(scoops_ordered__gt=F('store_visits'))
```

### 4.2 Database Functions
use can use DB functions like `UPPER`, `LOWER` etc inside your code

## 5 Don't Drop Down to Row SQL Until It's necessary
Use it when you are SURE that this way can lead to better performance for example when you have complex queries over millions of rows of data. But this can break ORM cycle and you can use `row()` or **`extra()`** functions.

## 6 Add Indexed as Needed:
[read mode on docs](https://docs.djangoproject.com/en/3.2/ref/models/indexes/)

## 7 Transactions
Use can use transactions when you have multiple insertions in a view this can lead to data integrity while cause performance issues

### 7.1 Wrapping Each HTTP Request in a  Transaction
you can set[ `ATOMIC_REQUESTS`](https://docs.djangoproject.com/en/3.2/ref/settings/#std:setting-DATABASE-ATOMIC_REQUESTS) in django settings but this can fuck up performance because each HTTP request no matter it's update on database or get data from DB (which it doesn't need transaction because no data would change while reading) will ba inside a transaction, you can either use [`transaction.non_atomic_request`](https://docs.djangoproject.com/en/3.2/topics/db/transactions/#django.db.transaction.non_atomic_requests).
<br />
consider this situation when you want to insert a Email Confirmation into DB and error happened, if you put all this inside a transaction you will lose your data so use can use again `non_atomic_request` and explicitly use `transaction.atomic()` inside view. consider [this example](https://github.com/feldroy/two-scoops-of-django-3.x/blob/master/code/chapter_07_example_11.py):
```python
# flavors/views.py

from django.db import transaction
from django.http import HttpResponse
from django.shortcuts import get_object_or_404
from django.utils import timezone

from .models import Flavor

@transaction.non_atomic_requests
def posting_flavor_status(request, pk, status):
    flavor = get_object_or_404(Flavor, pk=pk)

    # This will execute in autocommit mode (Django's default).
    flavor.latest_status_change_attempt = timezone.now()
    flavor.save()

    with transaction.atomic():
        # This code executes inside a transaction.
        flavor.status = status
        flavor.latest_status_change_success = timezone.now()
        flavor.save()
        return HttpResponse('Hooray')

    # If the transaction fails, return the appropriate status
    return HttpResponse('Sadness', status_code=400)
```
second way is to instead of use `ATOMIC_REQUESTS=True` you can set it to False and:

### 7.2 Explicit Transaction Declaration
You can explicitly declare atomic transaction whenever you want for example you can use transaction on `create`, `bulk_create`, `get_or_create`, `update`, `delete` and ignore transaction on `get`, `filter`, `count`, `iterate`, `exists`, `exclude`, `in_bulk` etc
