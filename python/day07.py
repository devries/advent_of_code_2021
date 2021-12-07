#!/usr/bin/env python

def main():
    with open('../inputs/day07.txt', 'r') as f:
        content = f.read()

    numbers = [int(v) for v in content.split(',')]
    numbers = sorted(numbers)

    # Part A
    # Median
    x0 = numbers[len(numbers)//2]
    print("Part A: ", fuelA(numbers, x0))

    # Part B
    # Do bisection here
    interval = numbers[-1]-numbers[0]
    c = numbers[0]+interval//2
    known = {}
    interval//=4

    while True:
        h=c+interval
        l=c-interval

        cfuel = fuelB(numbers, c, known)
        hfuel = fuelB(numbers, h, known)
        lfuel = fuelB(numbers, l, known)

        if interval==1 and cfuel <= hfuel and cfuel<=lfuel:
            print("Part B: ", cfuel)
            break
        elif lfuel <= cfuel and lfuel <= hfuel:
            c = l
        elif hfuel <= cfuel and hfuel <= lfuel:
            c = h

        if interval>1:
            interval//=2

def fuelA(values, x0):
    fuel = 0
    for v in values:
        fuel+=abs(v-x0)

    return fuel

def fuelB(values, x0, known):
    if x0 in known:
        return known[x0]

    fuel = 0
    for v in values:
        fuel+=sum(range(1,abs(v-x0)+1))

    known[x0] = fuel
    return fuel

if __name__=="__main__":
    main()
