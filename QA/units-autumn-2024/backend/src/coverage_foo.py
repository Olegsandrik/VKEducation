import math


class MyClass:

    def __init__(self):
        pass

    def my_foo(self, a, b):
        if a == 3:
            a += 1
        if b == 4:
            b = b / a

        return a + b

    # my_foo('a', 1.3)
    # my_foo(None, 1)
