import time

from functools import wraps

from PIL import Image


def timeit(fn):
    wraps(fn)
    def measure_time(*args, **kwargs):
        s1 = time.time()
        res = fn(*args,**kwargs)
        print("%s executed in %s seconds" % (
            fn.__name__, time.time() - s1)
            )
        return res
    return measure_time
@profile
def julia():
    w, h, zoom = 1920,1080,1
    bitmap = Image.new("RGB", (w, h), "white")
    pix = bitmap.load()
    cX, cY = -0.7, 0.27015
    moveX, moveY = 0.0, 0.0
    maxIter = 255
    for x in range(w):
        for y in range(h):
            zx = 1.5*(x - w/2)/(0.5*zoom*w) + moveX
            zy = 1.0*(y - h/2)/(0.5*zoom*h) + moveY
            i = maxIter
            while zx*zx + zy*zy < 4 and i > 1:
                tmp = zx*zx - zy*zy + cX
                zy,zx = 2.0*zx*zy + cY, tmp
                i -= 1

    
            pix[x,y] = (i << 21) + (i << 10) + i*8


    bitmap.show()

julia()
