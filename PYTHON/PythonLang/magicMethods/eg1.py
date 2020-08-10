class Tester:
    def __init__(self):
        print('This is init')
    def __call__(self):
        print('Called')

t = Tester()
# This is init
t()
# Called
