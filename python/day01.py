#!/usr/bin/env python

def main():
    with open('../inputs/day01.txt', 'r') as f:
        values = [int(ln.strip()) for ln in f.readlines()]

    # Part A
    increases = sum([1 for (a,b) in zip(values,values[1:]) if a<b])
    print("Part A:",increases)

    # Part B
    windows = [sum(t) for t in zip(values, values[1:], values[2:])]
    increases = sum([1 for (a,b) in zip(windows, windows[1:]) if a<b])
    print("Part B:",increases)

if __name__=="__main__":
    main()
