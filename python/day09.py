#!/usr/bin/env python
from functools import reduce

def main():
    grid = {}
    with open('../inputs/day09.txt', 'r') as f:
        for j, ln in enumerate(f.readlines()):
            for i, n in enumerate(ln.strip()):
                grid[(i,j)]=int(n)

    # Part A
    lowpts = []
    for k,v in grid.items():
        for neigh in neighbors(k):
            if neigh in grid and v >= grid[neigh]:
                break
        else:
            lowpts.append(k)

    risksum = sum([grid[p]+1 for p in lowpts])
    print("Part A: ", risksum)

    # Part B
    sizes = []
    for lowpt in lowpts:
        found = set([lowpt])
        wq = [lowpt]
        while len(wq)>0:
            for n in neighbors(wq.pop(0)):
                if n in grid and grid[n]<9 and n not in found:
                    wq.append(n)
                    found.add(n)
        sizes.append(len(found))

    sizes.sort(reverse=True)
    print("Part B: ", reduce(lambda a,b: a*b, sizes[:3], 1))

def neighbors(pt):
    for d in [(1,0),(-1,0),(0,1),(0,-1)]:
        yield tuple(a+b for a,b in zip(pt, d))

if __name__=="__main__":
    main()
