from behave import *
from json import loads
from ast import literal_eval

def _decore_list(arr):
    return set(list(arr.split(',')))

@Given('a {Dictionary} of stations with {InputList} of state needed')
def give_data(context, Dictionary,InputList):
    context.stations = literal_eval(Dictionary)
    context.state_needed = _decore_list(InputList)

@When('we insert them')
def insersion(context):
    pass

@Then('should return minium Amount of stations as {OutputList}')
def confirm(context, OutputList):
    from implementation import station_set_example
    assert station_set_example(context.stations, context.state_needed) == _decore_list(OutputList)

