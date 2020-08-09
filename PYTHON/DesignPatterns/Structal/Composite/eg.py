from abc import ABC, abstractclassmethod

class Draw(ABC):
    @abstractclassmethod
    def render(self):
        raise NotImplementedError("This is an absctractnmethod! Do not Implement this")
        
class CompositeGraphic(Draw):
    def __init__(self):
        self.graphics = []
        
    def render(self):
        for graphic in self.graphics:
            graphic.render()
        
    def graphic_setter(self, graphic):
        self.graphics.append(graphic)
    def graphic_del(self, graphic):
        self.graphics.remove(graphic)
    
    graphic = property(None, graphic_setter, graphic_del, "This is for managing graphics")
    
class Ellipse:
    def __init__(self, name):
        self.name = name
    def render(self):
        print(f'name is {self.name}')
        
if __name__ == '__main__':
    ellipse1 = Ellipse("1")
    ellipse2 = Ellipse("2")
    ellipse3 = Ellipse("3")
    ellipse4 = Ellipse("4")
    
    graphic1 = CompositeGraphic()
    
    graphic1.graphic = ellipse1
    graphic1.graphic = ellipse2
    graphic1.graphic = ellipse3
    graphic1.graphic = ellipse4
    
    graphic1.render()
