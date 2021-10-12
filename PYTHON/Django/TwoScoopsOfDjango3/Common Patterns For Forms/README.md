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

