import unittest
from coverage_foo import MyClass
import math


class TestMyClass(unittest.TestCase):

    def setUp(self):
        self.my_class = MyClass()

    def test_my_class_1(self):
        self.assertEqual(self.my_class.my_foo(3, 4), 5)

    def test_my_class_2(self):
        self.assertRaises(ZeroDivisionError, self.my_class.my_foo, 0, 4)

    def test_my_class_3(self):
        self.assertRaises(TypeError, self.my_class.my_foo, None, None)

    def test_my_class_4(self):
        self.assertEqual(self.my_class.my_foo(math.inf, math.inf), math.inf)

    def test_my_class_5(self):
        self.assertEqual(self.my_class.my_foo(1.2, 2.3), 3.5)

    def test_my_class_6(self):
        self.assertRaises(TypeError, self.my_class.my_foo, 'abc', 4)

    def test_my_class_7(self):
        self.assertEqual(self.my_class.my_foo(-2, -4), -6)


if __name__ == "__main__":
    unittest.main()
