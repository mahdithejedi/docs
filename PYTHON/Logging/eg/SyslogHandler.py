import logging
import logging.handlers

#set handler for calling UNIX syslog to handle log throught UDP:514
#POINT: so our system should have a open port in 514
Handler=logging.handlers.SysLogHandler(address='/dev/log')

logging.basicConfig(
        format='[%(asctime)s] %(filename)s -> %(funcName)s:%(module)s::%(levelname)s %(message)s',
        datefmt='%Y:%m:%d %H:%M:%S',
        #say to python to use Handler instead of streamHandler
        handlers=(Handler, ))
logger = logging.getLogger('AppLogger')

logger.setLevel(logging.DEBUG)


logger.info('This  is  a simple info')

logger.critical('Oh SHIT ERROR!')
