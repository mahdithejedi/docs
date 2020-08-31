#!/usr/bin/env python
# coding: utf-8

# In[1]:


from abc import ABCMeta, abstractclassmethod


# In[2]:


class AbstractInterface(metaclass=ABCMeta):
    @abstractclassmethod
    def ShowSockets(self):
        NotImplemented()


# In[3]:


class Bridge(AbstractInterface):
    def __init__(self):
        self._implementation = None


# In[4]:


class FireWall(Bridge):
    def __init__(self, implementation: 'ImplementationInterface'):
        self._implementation = implementation
        
    def ShowSockets(self):
        print('sockets for firewall app are')
        self._implementation.get_sockets()


# In[5]:


class NetworkManager(Bridge):
    def __init__(self, implementation: 'ImplementationInterface'):
        self._implementation = implementation
        
    def ShowSockets(self):
        print('sockets for network manager app are')
        self._implementation.get_sockets()


# In[6]:


### Implenetors


# In[7]:


class ImplementationInterface(metaclass=ABCMeta):
    @abstractclassmethod
    def get_sockets(self):
        raise NotImplemented()


# In[8]:


class Unix(ImplementationInterface):
    def _find_sockets(self):
        for num in range(1, 19):
            print("socket numeber{} found".format(num))
    def get_sockets(self):
        self._find_sockets()
        print ("Some Unix sockets")
        return "Unix"


# In[9]:


class Linux(Unix):
    def get_sockets(self):
        self._find_sockets()
        print("Some linux sockets")
        return "Linux"


# In[10]:


# You can have debian linux which inherit Linux and Ubuntu which inheric debian
# Inheritance for Ubuntu
# Ubuntu 18.1 -> Ubuntu 18 -> Ubuntu -> Debian -> Linux -> Unix -> ImplementationInterface


# In[11]:


class Windows(ImplementationInterface):
    def get_sockets(self):
        print("Some windows socket")
        return "windows"


# In[12]:


def os():
    state = 'run'
    unix = Unix()
    windows = Windows()
    
    unix_run_fire_wall = FireWall(unix).ShowSockets()
    unix_run_network_manager = NetworkManager(unix).ShowSockets()
    # ----
    win_run_fire_wall = FireWall(windows).ShowSockets()
    unix_run_network_manager = NetworkManager(windows).ShowSockets()


# In[13]:


if __name__ == '__main__':
    os()


# In[ ]:




