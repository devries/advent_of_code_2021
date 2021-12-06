#!/usr/bin/env python

def main():
    with open('../inputs/day02.txt', 'r') as f:
        lines = f.readlines()

    command_parts = [ln.strip().split() for ln in lines]

    commands = [(a,int(b)) for a,b in command_parts]

    # Part A
    horizontal = 0
    depth = 0
    for command, amount in commands:
        if command=="forward":
            horizontal+=amount
        elif command=="down":
            depth+=amount
        elif command=="up":
            depth-=amount

    print("Part A:", horizontal*depth)

    # Part B
    horizontal = 0
    depth = 0
    aim = 0
    for command, amount in commands:
        if command=="forward":
            horizontal+=amount
            depth+=aim*amount
        elif command=="down":
            aim+=amount
        elif command=="up":
            aim-=amount

    print("Part B:", horizontal*depth)

if __name__=="__main__":
    main()
