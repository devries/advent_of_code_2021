#!/usr/bin/env python
import itertools

def main():
    with open('../inputs/day05.txt', 'r') as f:
        lines = f.readlines()

    grid = {}
    for line in lines:
        p1, p2 = line.split(' -> ')
        x1, y1 = [int(v.strip()) for v in p1.split(',')]
        x2, y2 = [int(v.strip()) for v in p2.split(',')]

        if x1==x2:
            deltay = int((y2-y1)/abs(y2-y1))

            for y in range(y1, y2, deltay):
                grid[(x1,y)]=grid.setdefault((x1,y), 0)+1
            else:
                grid[(x2,y2)]=grid.setdefault((x2,y2), 0)+1


        elif y1==y2:
            deltax = int((x2-x1)/abs(x2-x1))

            for x in range(x1, x2, deltax):
                grid[(x,y1)]=grid.setdefault((x,y1), 0)+1
            else:
                grid[(x2,y2)]=grid.setdefault((x2,y2), 0)+1

    sum = 0
    for k,v in grid.items():
        if v >=2:
            sum+=1

    print("Part A:",sum)

    grid = {}
    for line in lines:
        p1, p2 = line.split(' -> ')
        x1, y1 = [int(v.strip()) for v in p1.split(',')]
        x2, y2 = [int(v.strip()) for v in p2.split(',')]

        if x1==x2:
            xpoints = [x1]*int(abs(y2-y1))
        else:
            deltax = (x2-x1)//abs(x2-x1)
            xpoints = range(x1, x2, deltax)

        if y1==y2:
            ypoints = [y2]*int(abs(x2-x1))
        else:
            deltay = (y2-y1)//abs(y2-y1)
            ypoints = range(y1, y2, deltay)


        points = list(zip(xpoints, ypoints))
        points+=[(x2,y2)]

        for x, y in points:
            grid[(x,y)]=grid.setdefault((x,y), 0)+1

    sum = 0
    for k,v in grid.items():
        if v >=2:
            sum+=1

    print("Part B:",sum)
    
if __name__=="__main__":
    main()
