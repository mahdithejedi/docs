from behave import *

@given('we have behave installed')
def step_impl(context):
    print("context inside given is %s with type %s and dict %s" % (context, type(context), context.__dict__))

@when('we implement a test')
def step_impl(context):
    print("context inside when is %s with type %s and dict %s" % (context, type(context), context.__dict__))

@then('behave will test it for us!')
def step_impl(context):
    help(eq_)
    print("context inside then is %s with type %s and dict  %s" % (context, type(context), context.__dict__))

