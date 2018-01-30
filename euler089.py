from euler import timer
import re
from roman import *
import urllib.request

@timer
def euler089():
    numbers = ''.join(line for line in open('data/p089_roman.txt'))
    print(len(numbers) - len(re.sub("DCCCC|LXXXX|VIIII|CCCC|XXXX|IIII", '  ', numbers)))


euler089()
