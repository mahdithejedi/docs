class Singleton(object):
    """ This is the best approch for Singleton"""
    _instance = None

    def __new__(cls):
        if cls._instance == None:
            cls._instance = super().__new__(cls)
        return cls._instance
