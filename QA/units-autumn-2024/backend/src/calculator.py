import math


class Calculator:

    def __init__(self):
        pass

    def addition(self, x1, x2):
        return x1 + x2

    def multiplication(self, x1, x2):
        return x1 * x2

    def subtraction(self, x1, x2):
        return x1 - x2

    def division(self, x1, x2):
        if x2 != 0:
            return x1 / x2

    def absolute(self, x):
        return abs(x)

    def degree(self, x, n):
        return x ** n

    def ln(self, x):
        return math.log(x)

    def log(self, x, n):
        return math.log(x, n)

    def sqrt(self, x):
        return x ** 0.5

    def nth_root(self, x, n):
        return x ** (1. / n)
