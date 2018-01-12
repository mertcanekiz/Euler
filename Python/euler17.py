from euler import timer
from num2words import num2words


@timer
def euler17():
    print(sum(len(num2words(i).replace(' ', '').replace('-', '')) for i in range(1, 1001)))


euler17()
