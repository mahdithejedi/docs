# 12 Common Patterns Forms
forms are great with Models and CBVs here we cover 5 common form patterns

## 1 Pattern1: simple `ModelForm` with default validators
```python
class TestCreateView(LoginRequiredMixin, CreateView):
    model = Test
    fields = ['name', 'number']
```
In here
*   view auto-generate `ModelForms`
*   There `ModelForma` are relay on default validation rules of `Test` model

 
 ## 2 Custom Form Field Validators
 we can use functins to validate attributes for example
 ```python
 # core/validators.py
 from django.core.exceptions import ValidationError

 def gmail_validator(value):
     """
     test mail is gmail
     """
     if not value.endswith('@gmail.com') or value.endswith('@googlemail.com'):
         msg = '%s should be gmail' % value
         raise ValidationError(
             msg
         )
```
*Sometimes validators can be complex so write strict  test for them*


not you can *use it* either in models:

```python
# core/models
from django.db import models
from .validatos import from django import forms
from .models import Flavor
from core.validators import gmail_validator


class GmailAbstractModel(models.Model):
    gmail = models.Charfield(
        max_length=255, validators=[gmail_validator]
    )
    class Meta:
        abstract = True
```

or in the form:
```python
# mail/forms.py

from django import forms
from .models import Mail
from core.validators import gmail_validator

class MailForm(forms.Modelform):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.fields['gmail'].validators.append(gmail_validator)

    class Meta:
        model = Mail
```
and use this form in views
```python

class MailActionMixin:
    model = Mail
    fields = ['gmail', 'password']

class MailCreateView(LoginRequiredMixin, MailActionMixin, Createview):
    form_class = MailForm
```


## 3 Overriding clean stage of Validation
you can use `clean` and `clean_<FIELDـNAME>` to do multi-field or more validation. there methods will call after form validation for example you want user to have epecific mail addresses
```python
from .models import DomainMail
class MailForm(forms.Form):
    email = forms.Charfield()
    password = forms.Charfield()
    def clean_email(self):
        mail = self.clean_data['email']
        mail_name, host = mail.split('@')
        try:
            # check  dynamicly if given host name is in database
            DomainMail.object.get(domain=host)
        except DomianMail.DoesNotExists:
            raise forms.ValidationError(
                'mail is not valid'
            )
        return mail
```
and we want to make sure email and password are not the same
```python
def clean(self):
    cleaned_data = super().clean()
    email = clean_data['email']
    password = clean_data['password']
    if email == password:
        raise forms.ValidationError(
            'email and password can\'t be the same'
        )
    return cleaned_data
setattr(MailForm, 'clean', clean)
```


