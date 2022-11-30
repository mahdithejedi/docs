from typing import List
from functools import partial
from abc import ABC


class _BasePrototype:
    def __init__(self, _type: str):
        self._type = _type

    def __str__(self):
        return self._type


class _BasePrototypeBuilder:
    def __new__(cls, *args, **kwargs):
        return type(
            args[0], (_BasePrototype,), {}
        )


Wheel = _BasePrototypeBuilder('Wheel')
Windows = _BasePrototypeBuilder('Windows')


class BaseCarBuilder(ABC):
    def __init__(self):
        self._wheels: List[Wheel] = []
        self._body: str = None
        self._windows: List[Windows] = []

    def get_wheel(self) -> List[Windows]:
        return self._wheels

    def get_windows(self):
        return self._windows

    def get_body(self):
        return self._body


class Manager:
    def __init__(self, builder: BaseCarBuilder):
        self._builder: BaseCarBuilder = builder
        self._car = ''

    def create_car(self):
        self.make_without_wheels()
        self._car += ' '.join(
            [' wheel %s is %s ' % (index, str(wh)) for index, wh in enumerate(self._builder.get_wheel())]) + '\n'
        return self._car

    def make_without_wheels(self):
        self._car += (' body is ' + str(self._builder.get_body()) + '\n')
        self._car += ' '.join(
            ['windows %s is %s ' % (index, str(win)) for index, win in enumerate(
                self._builder.get_windows()
            )]) + '\n'
        return self._car

    def reset(self):
        self._builder = None
        self._car = ''

    def __set_builder(self, builder: BaseCarBuilder):
        self.reset()
        self._builder = builder

    builder = property(fset=__set_builder, doc="'builder' manager")


class SimpleCar(BaseCarBuilder):
    def set_wheels(self, wheel_type: str):
        self._wheels = [Wheel(wheel_type) for _ in range(4)]

    def set_windows(self, window_type: str):
        self._windows = [Windows(window_type) for _ in range(4)]

    def set_body(self, body: str):
        self._body = body


if __name__ == '__main__':
    car = SimpleCar()
    car.set_wheels("Truck")
    car.set_windows("pride")
    car.set_body("simple")
    manager = Manager(car)
    print(manager.create_car())
