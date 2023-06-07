#!/usr/bin/env python
# coding: utf-8

# In[1]:


from abc import ABC, abstractmethod


# In[2]:


## We should have two seprate class
## one for windows itself and one for it's scorling and other features
## But this two classes need a same Interface
class WindowInterface(ABC):
    @abstractmethod
    def build(self):
        raise NotImplementedError()


# In[3]:


class Window(WindowInterface):
    """
        The Window class
    """
    def build(self):
        print("this class build successfully")


# In[4]:


class AbstractWindowDecorator(WindowInterface, ABC):
    """
    Maintain a reference to a Window object and define an interface
    that conforms to Window's interface.
    """
    def __init__(self, window):
            self._window = window
    @abstractmethod
    def build(self):
        raise NotImplementedError()


# In[5]:


class BorderDecorator(AbstractWindowDecorator):
    def add_border(self):
        print('Border added')
    def build(self):
        self.add_border()
        self._window.build()


# In[6]:


class VerticalSBDecorator(AbstractWindowDecorator):
    def add_vertical_scroll_bar(self):
        print("Adding vertical scroll bar")

    def build(self):
        self.add_vertical_scroll_bar()
        self._window.build()


# In[7]:


class HorizontalSBDecorator(AbstractWindowDecorator):
    def add_horizontal_scroll_bar(self):
        print("Adding horizontal scroll bar")

    def build(self):
        self.add_horizontal_scroll_bar()
        self._window.build()

