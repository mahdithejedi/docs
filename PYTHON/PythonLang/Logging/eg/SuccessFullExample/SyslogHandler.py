import logging
import logging.handlers

FMT='[%(asctime)s] %(filename)s -> %(funcName)s:%(module)s::%(levelname)s %(message)s'
DATEFMT='%Y:%m:%d %H:%M:%S'
APP_NAME = 'UssdApp'

#set handler for calling UNIX syslog to handle log throught UDP:514
#POINT: so our system should have a open port in 514
Handler=logging.handlers.SysLogHandler(address='/dev/log')
formatter  = logging.Formatter('%(message)s', DATEFMT)
Handler.setFormatter(formatter)
Handler.ident = APP_NAME

logger = logging.getLogger(APP_NAME)
logger.addHandler(Handler)
logger.setLevel(logging.DEBUG)

logger.info('\n\n\n\n\n\n\n\n\nNEW \n\n\n\n\n\n\n\n\n\n')

logger.critical('Oh SHIT ERROR! CRITICAL ONE IS HAPPENED')
