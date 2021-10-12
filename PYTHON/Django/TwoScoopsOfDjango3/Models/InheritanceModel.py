

##############################
######  DJANGO OBJECTS  ######
#############################
# from django.db import models

# # Create your models here.

# class StudentManager(models.Manager):
# 	def get_moosavies(self):
# 		return super(StudentManager, self).get_queryset().filter(
# 				family__in = (
# 					'moosavi', 'Moosavi'
# 					)
# 			)


# class Student(models.Model):
# 	name = models.CharField(max_length=25)
# 	family = models.CharField(max_length=25)

# 	objects = StudentManager()

# 	class Meta:
# 		unique_together = ('name', 'family')

# 	def is_moosavi(self):
# 		return self.family == 'moosavi'

# class Course(models.Model):
# 	course_name = models.CharField(max_length=25)
# 	student = models.ManyToManyField(Student)



#####################################################
######  DJANGO INHERITANCE (multi-inhritance)  ######
#####################################################

#? each child of Place table will have a foriegn key to place
#? tables in DB:
#? SeprateLogic_place, SeprateLogic_resturant, SeprateLogic_school

#? if you add data to Place you **DON'T** have access to it from it's childer(like School) but
#? they have access to it's field

from django.db import models

class Place(models.Model):
	location = models.IntegerField(blank=True)
	name = models.CharField(max_length=80, blank=False, null=False)
	open_date = models.DateTimeField(blank=True)

class School(Place):
	student_count = models.SmallIntegerField(blank=True, null=True)

class Resturant(Place):
	cheff_count = models.SmallIntegerField(blank=True, null=True)



##########################################
######  DJANGO INHERITANCE (Proxy)  ######
##########################################

## this is used for having manged multiple Objects
#? READ MORE https://www.benlopatin.com/using-django-proxy-models/
#? Django OFFICIAL documents https://docs.djangoproject.com/en/3.2/topics/db/models/



#############################################
######  DJANGO INHERITANCE (Bastract)  ######
#############################################
#? READ MORE https://www.geeksforgeeks.org/how-to-create-abstract-model-class-in-django/
#? Django OFFICIAL documents https://docs.djangoproject.com/en/3.2/topics/db/models/
