def fib(n):
    a, b = 0, 1
    while a<b:
        print a,
        a, b = b, a+b
fib(42)
