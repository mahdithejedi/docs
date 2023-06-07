# Model Best Practices

## 1.Basics

### 1.1 Break Up Apps With To Many Models
it's better to have 200 apps with 10 models to have 200 models with 10 apps (_break up apps to smaller one_)

### 1.2 Be Careful With Model Inheritance
[Read more here](./InheritanceModel.py)


## 2 Migrations

### 2.1 Edit django migrations
You can edit your django migrations with [runpython](https://docs.djangoproject.com/en/3.2/ref/migration-operations/#django.db.migrations.operations.RunPython)  (<small>**recemended**</small>) or [runsql](https://docs.djangoproject.com/en/3.2/ref/migration-operations/#django.db.migrations.operations.RunSQL)
<br />
you can have a small example [here](./RunPythonExample.py)!


## 2.2 Deployment and Management of Migrations
1- ***FOR GOD SAKE Always backup your data before migrations***
<br />
2- before deployment check that you can rollback migrations!
<br />
3- In big project do extensive test against data of that size on staging server before run migrations on production server. Migrations sometimes can take hours
<br />
4- ***FOR GOD SAKE Always put Data Migrations code into source code***
 

# 4 Django Model Design

## 4.1 Start Normalize
you can use normalization to have better Django Model Design. you can read more about Normalization in [here](../../../../DB/normalization.md)
<br />
Another approach is [demoralize](https://en.wikipedia.org/wiki/Denormalization)

## 4.2 Cache Before Denoralizing
Before caching try to use cache in right place

## 4.3 Denormalize Only if Absolutely Needed

## 4.4 When to Use Null and Blank
for `CharField`, `TextField`, `SlugField`, `EmailField`, `CommaSeparatedIntegerField`, `IntegerField`, `UUIDField` Fields Type empty values store as **`NULL`** in DB if `null=True` and `unique=True` is set other wise it'll save it as **`empty string`**
<br />
for `FileField`, `ImageField` **DON'T USE `NULL=TRUE`**

## 4.5 Where to Use BinaryField
Use `BinaryField` if it's necessary because **read/write to a DB is always slower than a filesystem**

## 4.6 Try to Avoid Using Generic Relations
<!-- TODO -->

## 4.7 Make Choices and Sub-Choices Model Constants
```python
class IceCreamOrder(models.Model):
    FLAVOR_CHOCOLATE = 'ch'
    FLAVOR_VANILLA = 'vn'
    FLAVOR_STRAWBERRY = 'st'
    FLAVOR_CHUNKY_MUNKY = 'cm'

    FLAVOR_CHOICES = (
        (FLAVOR_CHOCOLATE, 'Chocolate'),
        (FLAVOR_VANILLA, 'Vanilla'),
        (FLAVOR_STRAWBERRY, 'Strawberry'),
        (FLAVOR_CHUNKY_MUNKY, 'Chunky Munky')
    )

    flavor = models.CharField(
        max_length=2,
        choices=FLAVOR_CHOICES
    )
```
you can query it like 
```python
IceCreamOrder.objects.filter(flavor=IceCreamOrder.FLAVOR_CHOCOLATE)
```

## 4.8 Using Enumeration Types for Choices
```python
class IceCreamOrder(models.Model):
    class Flavors(models.TextChoices):
        CHOCOLATE = 'ch', 'Chocolate'
        VANILLA = 'vn', 'Vanilla'
        STRAWBERRY = 'st', 'Strawberry'
        CHUNKY_MUNKY = 'cm', 'Chunky Munky'

    flavor = models.CharField(
        max_length=2,
        choices=Flavors.choices
    )
```
so you can query it like
```
IceCreamOrder.objects.filter(
    flavor=IceCreamOrder.Flavors.CHOCOLATE
)
```


# 5 META class
<!-- TODO -->

# 6 Model Managers
We can use ModelManagers to customize and add queries but there are some point we have to consider

1- when using model inheritance, children of abstract base classes receive their parent's model manager, the children of concrete base classes do not

2- the first manager applied to a model class is the one that Django treats as the default.

# 7 Understanding Fat Models
We can encapsulate logic in model with model methods, classmethods, properties and manager methods and this can lead to a great improvement and prevention of code reusing but we'll interfere a problem with hundreds or even thousands of complex code inside Model class.
**what can we do?**

## 7.1 Model Behaviors a.k.a Mixins
we can use mixins for models. [here](https://blog.kevinastone.com/django-model-behaviors) is a old article but good explanation 
<br />
<small>
before read above article it's better to read article's which list below to extend  your knowledge
<br />
[django manager vs query set](https://jairvercosa.medium.com/manger-vs-query-sets-in-django-e9af7ed744e0)
<br />
[django from_queryset](https://docs.djangoproject.com/en/3.2/topics/db/managers/#from-queryset)
<br />
[facade design pattern](https://en.wikipedia.org/wiki/Facade_pattern)
<br />
<small>
[lazy_load design pattern](https://www.geeksforgeeks.org/lazy-loading-design-pattern/)
</small>
</small>

## 7.2 Stateless Helper Functions
Moving logics out of models and into utility functions

## 7.3 Model Behaviors VS Helper Functions
It's better to use both of them together. How use them together is depends on experience and doesn't have absolute measurement