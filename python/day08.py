#!/usr/bin/env python
from itertools import permutations

def main():
    with open('../inputs/day08.txt', 'r') as f:
        lines = f.readlines()

    # Part A
    total = 0
    for ln in lines:
        nums, disp = parse_line(ln)
        for d in disp:
            l = len(d)
            if l==2 or l==3 or l==4 or l==7:
                total+=1

    print("Part A:", total)

    # Part B
    # Find all possible permutations of abcdefg
    perms = permutations('abcdefg')

    # Find all possible letter to letter conversions
    conversions = [dict(zip(p,'abcdefg')) for p in perms]

    total = 0
    for ln in lines:
        nums, disp = parse_line(ln)

        for c in conversions:
            if check_conversion(nums, c):
                # This is the right value
                total+=convert_display(disp,c)

    print("Part B:", total)


def parse_line(ln):
    parts = ln.split(' | ')
    numbers = parts[0].split()
    display = parts[1].split()

    return (numbers, display)

def convert(values, c):
    # Convert a wire set using conversion dictionary
    return ''.join([c[v] for v in values])

def check_conversion(nums, c):
    # Check that all conversions of the numbers list work
    for n in nums:
        nc = convert(n, c)
        if get_value(nc)==None:
            return False

    return True

def get_value(active):
    # Sort wire list and return corresponding number
    s = ''.join(sorted(active))
    if s=='abcefg':
        return 0
    elif s=='cf':
        return 1
    elif s=='acdeg':
        return 2
    elif s=='acdfg':
        return 3
    elif s=='bcdf':
        return 4
    elif s=='abdfg':
        return 5
    elif s=='abdefg':
        return 6
    elif s=='acf':
        return 7
    elif s=='abcdefg':
        return 8
    elif s=='abcdfg':
        return 9
    else:
        return None

def convert_display(disp, c):
    # Convert the list of four active wire sets in display to a number
    multiplier = 1
    retval = 0

    for item in reversed(disp):
        retval+=multiplier*get_value(convert(item, c))
        multiplier*=10

    return retval

if __name__=="__main__":
    main()
