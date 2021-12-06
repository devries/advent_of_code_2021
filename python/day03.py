#!/usr/bin/env python

def main():
    with open('../inputs/day03.txt', 'r') as f:
        lines = [ln.strip() for ln in f.readlines()]

    width = len(lines[0])
    values = [int(v, 2) for v in lines]

    masks = [1<<i for i in range(width)]
    masks.reverse()

    gamma = 0
    epsilon = 0

    oxygen = 0
    oxygenset = set(values)

    co2 = 0
    co2set = set(values)

    for m in masks:
        zeros = sum([1 for v in values if m&v==0])
        ones = sum([1 for v in values if m&v>0])
        
        if ones > zeros:
            gamma |= m
        else:
            epsilon |= m

        # Oxygen
        if len(oxygenset)==0:
            pass
        elif len(oxygenset)==1:
            oxygen = oxygenset.pop()
        else:
            zeros = sum([1 for v in oxygenset if m&v==0])
            ones = sum([1 for v in oxygenset if m&v>0])

            if ones >= zeros:
                oxygen |= m
                oxygenset = set([v for v in oxygenset if m&v>0])
            else:
                oxygenset = set([v for v in oxygenset if m&v==0])

        # CO2
        if len(co2set)==0:
            pass
        elif len(co2set)==1:
            co2 = co2set.pop()
        else:
            zeros = sum([1 for v in co2set if m&v==0])
            ones = sum([1 for v in co2set if m&v>0])

            if ones < zeros:
                co2 |= m
                co2set = set([v for v in co2set if m&v>0])
            else:
                co2set = set([v for v in co2set if m&v==0])


    print("Part A:", gamma*epsilon)
    print("Part B:", oxygen*co2)

if __name__=="__main__":
    main()
