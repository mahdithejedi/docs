#!/usr/bin/env python
# coding: utf-8

# In[1]:


__source__ = "https://www.geeksforgeeks.org/strategy-method-python-design-patterns/"


# In[2]:


# We have a store which have diffrent discounts


# In[3]:


class Items:
    def __init__(self, price, discount_strategy = None):
        self.price = price
        self._discount_strategy = discount_strategy
        
    def calc_price(self, discount_kwargs = {}):
        self.price = self._discount_strategy(self.price, **discount_kwargs)
        return self.price
    
    def __repr__(self):
        return '{}'.format(self.price)


# In[4]:


def half_discount(price):
    return int(price/2)


# In[5]:


def holiday_discount(price, discount_per = 0.2):
    return int(price * discount_per)


# In[6]:


black_friday = Items(12, half_discount).calc_price()
print(black_friday)


# In[7]:


holiday = Items(12000, holiday_discount).calc_price(
    discount_kwargs = {
        'discount_per' : 0.3
    }
)
print(holiday)

