from euler import timer

def is_pentagonal(x):
    test = ((24 * x + 1)**0.5 + 1) / 6.0
    return test == int(test)

@timer
def euler45():
    i = 143
    while True:
        i += 1
        h = i*(2*i-1)
        if is_pentagonal(h):
            print(f"H({i}) = {h}")
            return


euler45()
