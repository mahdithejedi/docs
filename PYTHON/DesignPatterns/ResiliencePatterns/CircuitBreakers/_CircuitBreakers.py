# from enum import Enum
# from datetime import datetime
# from socket import timeout
# import requests
# import subprocess as sp
# from time import sleep


# class CircuitBreakersStates(Enum):
#     OPEN = "OPEN"
#     CLOSED = "CLOSED"
#     HALF_OPEN  = "HALF_OPEN"


# class CircuitBreakersStateHandler:
#     def __init__(self, max_failed_count=10, success_state_count=10, max_retry=3):
#         self._failed_count = 0
#         self._success_count = 0
#         self._max_retry = max_retry
#         self._retries = 0
#         self._max_failed_count = max_failed_count
#         self._success_state_count = success_state_count
#         self._state = CircuitBreakersStates.OPEN
    
#     def open(self):
#         return self._state == CircuitBreakersStates.OPEN
    
#     def close(self):
#         return self._state == CircuitBreakersStates.CLOSED
    
#     def success(self):
#         if self._state == CircuitBreakersStates.HALF_OPEN and  self._success_count < self._success_state_count:
#             self._success_count += 1
#         elif self._state == CircuitBreakersStates.HALF_OPEN and self._success_count >= self._success_state_count:
#             self._success_count = 0
#             print("\n\n\n OPEND \n\n\n")
#             self._state = CircuitBreakersStates.OPEN

#     @property
#     def state(self):
#         return self._state
    
#     def failed(self):
#         if not self.__handle_retry():
#             return
#         if self._state == CircuitBreakersStates.OPEN:
#             self._failed_count = 1
#             self._state = CircuitBreakersStates.HALF_OPEN
#             print("\n\n\n\n going to half state \n\n\n\n")
#         elif self._state == CircuitBreakersStates.HALF_OPEN and self._failed_count < self._max_failed_count:
#             self._failed_count += 1
#         elif self._state == CircuitBreakersStates.HALF_OPEN and self._failed_count >= self._max_failed_count:
#             self._failed_count = 0
#             print("\n\n\n\n CLOSED!! \n\n\n\n")
#             self._state = CircuitBreakersStates.CLOSED
    
#     def __handle_retry(self):
#         if self._retries > self._max_retry:
#             self._state = CircuitBreakersStates.OPEN
#             print("OPENDDD!!!!!***")
#             self._failed_count = 0
#             self._retries = 0
#             return False
#         self._retries += 1
#         return True

# class CircuitBreakerError(Exception):
#     pass


# class CircuitBreakers:
#     def __init__(self, retry_count = 10, success_state_count = 10, state_handler=CircuitBreakersStateHandler, state_exceptions=(), 
#                  re_raise=True):
#         self.retry_count = retry_count
#         self.success_state_count = success_state_count
#         self.state_handler = state_handler(retry_count, success_state_count, 3)
#         self.state_exceptions = state_exceptions
#         self._re_raise = re_raise
#         self._latest_exception = None
    

#     def __call__(self, func):
#         def _inner_call(*args, **kwargs):
#             return self.call(func, *args, **kwargs)
#         return _inner_call
    
#     def call(self, func, *args, **kwargs):
#         if self.state_handler.close():
#                 if self._re_raise:
#                     raise self._latest_exception
#                 return
#         try:
#             return func(*args, **kwargs)
#         except Exception as e:
#             self.__handle_error(e)
#         finally:
#             self.__handle_success()

#     def __handle_error(self, e):
#         self.state_handler.failed()
#         if self.state_handler.close():
#             if  self._re_raise:
#                 raise self._latest_exception
#             else:
#                 return
#         if not issubclass(e.__class__, self.state_exceptions):
#             self.state_handler.failed()
#             if self._re_raise:
#                 self._latest_exception = e
#                 raise e
#         else:
#             self.state_handler.success()
            
#     def __handle_success(self):
#         self.state_handler.success()
#         if self.state_handler.close():
#             self._latest_exception = None


# class Singleton(object):
#     _instance = None
#     def __new__(class_, *args, **kwargs):
#         if not isinstance(class_._instance, class_):
#             class_._instance = object.__new__(class_, *args, **kwargs)
#         return class_._instance

# class CircuitBreakersHandler(Singleton):
#     _HANDLERS = {}
#     def get(self, name,  circuit_breaker):
#         if name  not  in self._HANDLERS:
#             self._HANDLERS[name] = circuit_breaker
#         return self._HANDLERS[name]

#     def delete(self, name):
#         if name in self._HANDLERS:
#             del self._HANDLERS[name]
    
    
            

# def ping(ip):
#     status, _ = sp.getstatusoutput("ping -c 1 -w2 -i20{}".format(ip))
#     if status == 1:
#         raise Exception("Error!")

# circuit_breakers = CircuitBreakers()

# @circuit_breakers
# def __run(num):
#     requests.get("%s" % num, timeout=3)

# def run():
#     # circuit_breaker_main = CircuitBreakersHandler()
#     # circuit_breaker = circuit_breaker_main.get("main", CircuitBreakers())
#     num = 30
#     for _ in range(25):
#         t1 = datetime.now()
#         try:
#             if _ > 9:
#                 __run("http://www.google.com")
#             else:
#                 __run("http://10.10.34.%s" % num)
#             print("success")
#         except Exception as e:
#             print("Error %s" % e)
#         print("Time: {}".format(datetime.now() - t1))
#         num += 1
#         sleep(.1)
        

# run()