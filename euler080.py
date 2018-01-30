from euler import timer
from decimal import Decimal, getcontext

getcontext().prec = 105

@timer
def euler80():
    print(sum(sum(int(d) for d in str(Decimal(n).sqrt())[:101] if d != '.') for n in range(101) if int(n**0.5) != n**0.5))


euler80()
