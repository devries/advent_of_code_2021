#!/usr/bin/env python

def main():
    pairs = {'(': ')', '[': ']', '<': '>', '{': '}'}

    a_scores = {')': 3, ']': 57, '}': 1197, '>': 25137}
    b_scores = {')': 1, ']': 2, '}': 3, '>': 4}

    with open('../inputs/day10.txt', 'r') as f:
        lines = [ln.strip() for ln in f.readlines()]

    totalA = 0
    totalBs = []

    for ln in lines:
        stack = []

        for c in ln:
            if c in ['(', '[', '<', '{']:
                stack.append(c)
            else:
                cmatch = stack.pop()
                if pairs[cmatch]!=c:
                    totalA+=a_scores[c]
                    break

        else:
            subtotal = 0
            for r in reversed(stack):
                subtotal = subtotal*5+b_scores[pairs[r]]
            totalBs.append(subtotal)


    # Part A
    print("Part A: ", totalA)

    # Part B
    print("Part B: ", sorted(totalBs)[len(totalBs)//2])

    
if __name__=="__main__":
    main()
