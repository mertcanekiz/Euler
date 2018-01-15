from euler import timer


@timer
def euler40():
    d = ''.join(str(i) for i in range(185187))
    print(int(d[1]) * int(d[10]) * int(d[100]) * int(d[1000])
          * int(d[10000]) * int(d[100000]) * int(d[1000000]))


euler40()
