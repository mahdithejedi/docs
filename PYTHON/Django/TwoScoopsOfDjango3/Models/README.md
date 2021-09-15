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
3- In big project do extensive test againts data of that size on staging server before run migrations on production server. Migrations sometimes can take hours
<br />
4- ***FOR GOD SAKE Always put Data Migrations code into source code***
 


