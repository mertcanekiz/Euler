import time

def run(func):
    start = time.time()
    func()
    end = time.time()
    print("Elapsed time: %.4f ms" %((end-start)*1000))
