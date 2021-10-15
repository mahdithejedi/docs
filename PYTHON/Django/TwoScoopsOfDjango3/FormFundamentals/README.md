# 13 Form Fundamentals

## 1 Validate All Incoming Data With Django Forms
you can validate not only _HTTP requests_ with django forms but also any dictionary for example you want to validate  a CSV. you can use forms

```python
import csv

from django.utils.six import StringIS

from django import forms

from .model import User

class UserCreateForm(forms.ModelForm):
    class Meta:
        model = User
    
    def clean_username(self):
        username = self.cleaned_data['username']
        try:
            Users.object.get(username=username)
        except Users.DoesNotExists:
            return username
        finally:
            raise forms.ValidationError(
                'user %s already exists' % username
            )

def import_csv_users(csv_path, delimiter=','):
    rows = StringIO.StringIO(csv_path)
    record_added = 0
    errors = []

    for row in csv.DictReader(rows, delimiter):
        form = UserCreateForm(row)
        if form.is_valid:
            form.save()
            record_added += 1
        else:
            errors.append(
                form.errors
            )
    return record_added, errors
```

## 2 User the POST Method in HTML Forms
ANY forms submit that use alter data should done via POST method

## 3 Always User CSRF Protection With HTTP Forms THAT Modify Data

### 3.1 Posting Data Via AJAX
you shouldn't exempt CSRF in ajax

## 4 Understand How to Add Django Form Instance Attributes
```python
class FormTest(forms.ModelForm):
    def __init__(self, *args, **kwargs):
        self.user = kwargs.pop('user')
        super().__init__(*args, **kwargs)
```

## 5 How Form Validation Works
1 - `form.is_valid()` =>  2- `form.full_clean()` this will iterate over each field and validate them separately  -> Date coerced into python with `to_python` or `ValidateionError` will raise -> Date is validated for specific rule including custom validation -> if there is any custom `clean_<FIELD>()` avaibale they gonna call => 3- `form.clean` executed by `form.full_clean`

## 6 add error to form with `form.add_error`


## 7 Manage Custom Widget!
