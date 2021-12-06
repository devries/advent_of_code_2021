#!/usr/bin/env python

def main():
    with open('../inputs/day06.txt', 'r') as f:
        content = f.read()

    numbers = [int(v) for v in content.split(',')]

    population = [0]*9

    for n in numbers:
        population[n]+=1

    # Part A
    pop_a = population[:]
    evolve(pop_a, 80)
    print("Part A: ", sum(pop_a))

    # Part B
    pop_b = population[:]
    evolve(pop_b, 256)
    print("Part B: ", sum(pop_b))

def evolve(population, ngen):
    for g in range(ngen):
        pop0 = population[0]
        for i, v in enumerate(population):
            if i==0:
                continue
            population[i-1]=v

        population[8]=pop0
        population[6]+=pop0
    
if __name__=="__main__":
    main()
