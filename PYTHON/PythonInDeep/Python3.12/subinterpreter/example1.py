import _xxsubinterpreters as interpreters
import threading
import time

_shared_object = "HI"


def _run(objecter):
    import time;
    time.sleep(10);
    print(objecter)


# sleeper(2)
i = interpreters.create()
_t = threading.Thread(target=interpreters.run_string, name="interpreter%s" % i,
                      args=(i, _run, {"objecter": _shared_object})
                      )
_t.start()
# interpreters.run_string(i, """import time; time.sleep(10); print(objecter)""", {"objecter": _shared_object})
_shared_object = "BUY"
print("BYBY")
while interpreters.is_running(i):
    time.sleep(0.5)

interpreters.destroy(i)
_t.join()
