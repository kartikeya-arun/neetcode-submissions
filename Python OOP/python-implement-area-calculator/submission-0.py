import math

class AreaCalc:
    # TODO: Implement calculate method
    def calculate(self, length, width=None):
        if width:
            return length * width
        else:
            return f"{math.pi * length * length:.2f}"
# Don't modify the following code
calc = AreaCalc()
print(calc.calculate(5))    
print(calc.calculate(4, 6))
