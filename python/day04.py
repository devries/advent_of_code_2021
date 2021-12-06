#!/usr/bin/env python

def main():
    with open('../inputs/day04.txt', 'r') as f:
        lines = f.readlines()

    ordered_calls = [int(v) for v in lines[0].strip().split(',')]

    cards = []
    for i in range(2, len(lines), 6):
        cards.append(bingo_card(lines[i:i+5]))

    allcalls = set()
    firstdone = False
    for call in ordered_calls:
        allcalls.add(call)

        for c in cards:
            if not c.eliminated and c.bingo(allcalls):
                score = c.score(allcalls)*call
                c.eliminated = True
                if not firstdone:
                    print("Part A:", score)
                    firstdone = True

    print("Part B:", score)

class bingo_card:
    def __init__(self, lines):
        self.numbers = []
        for ln in lines:
            row = [int(v.strip()) for v in ln.strip().split()]
            self.numbers.append(row)
        self.eliminated = False
    
    def __repr__(self):
        return '\n'.join([' '.join([str(i) for i in r]) for r in self.numbers])

    def bingo(self, calls):
        # check rows
        for r in self.numbers:
            for n in r:
                if n not in calls:
                    break
            else:
                return True

        # Check columns
        for i in range(5):
            for r in self.numbers:
                if r[i] not in calls:
                    break
            else:
                return True

        return False

    def score(self, calls):
        total = 0
        for r in self.numbers:
            for n in r:
                if n not in calls:
                    total+=n

        return total

if __name__=="__main__":
    main()
