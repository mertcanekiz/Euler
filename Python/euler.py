import time

class memoize:
    def __init__(self, f):
        self.f = f
        self.memo = {}
    def __call__(self, *args):
        if not args in self.memo:
            self.memo[args] = self.f(*args)
        return self.memo[args]

def run(func):
    start = time.time()
    func()
    end = time.time()
    print("Elapsed time: %.4f ms" %((end-start)*1000))

@memoize
def fib(n):
    if n == 0:
        return 0
    if n == 1:
        return 1
    return fib(n-1) + fib(n-2)
