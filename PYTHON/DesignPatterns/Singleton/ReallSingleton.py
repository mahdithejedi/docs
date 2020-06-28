class SingleTon:
    _instance = None

    def __init__(self):
        if SingleTon._instance != None:
            pass
        else:
            SingleTon._instance = self

    @staticmethod
    def getIntance():
        """ Static Method for calling instance"""

        if SingleTon._instance == None:
            SingleTon()
        return SingleTon._instance
