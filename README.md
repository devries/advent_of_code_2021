# Advent of Code 2021

[![Tests](https://github.com/devries/advent_of_code_2021/actions/workflows/main.yml/badge.svg)](https://github.com/devries/advent_of_code_2021/actions/workflows/main.yml)
[![Stars: 44](https://img.shields.io/badge/⭐_Stars-44-yellow)](https://adventofcode.com/2021)

## Index

- [Day 1: Sonar Sweep](https://adventofcode.com/2021/day/1) - [part 1](day01_p1/main.go), [part 2](day01_p2/main.go)
- [Day 2: Dive!](https://adventofcode.com/2021/day/2) - [part 1](day02_p1/main.go), [part 2](day02_p2/main.go)
- [Day 3: Binary Diagnostic](https://adventofcode.com/2021/day/3) - [part 1](day03_p1/main.go), [part 2](day03_p2/main.go)
- [Day 4: Giant Squid](https://adventofcode.com/2021/day/4) - [part 1](day04_p1/main.go), [part 2](day04_p2/main.go)
- [Day 5: Hydrothermal Venture](https://adventofcode.com/2021/day/5) - [part 1](day05_p1/main.go), [part 2](day05_p2/main.go)
- [Day 6: Lanternfish](https://adventofcode.com/2021/day/6) - [part 1](day06_p1/main.go), [part 2](day06_p2/main.go)
- [Day 7: The Treachery of Whales](https://adventofcode.com/2021/day/7) - [part 1](day07_p1/main.go), [part 2](day07_p2/main.go)

  I enjoyed the second part of this exercise as I was able to use a simple
  minimization algorithm to bound and search the space. I divided the parameter
  space in half and then looked to see which half had a lower value.

- [Day 8: Seven Segment Search](https://adventofcode.com/2021/day/8) - [part 1](day08_p1/main.go), [part 2](day08_p2/main.go)

  This one was a great way to code up a set of deductions. I also ended up
  taking advantage of bit operations, representing each display as a byte.

- [Day 9: Smoke Basin](https://adventofcode.com/2021/day/9) - [part 1](day09_p1/main.go), [part 2](day09_p2/main.go)
- [Day 10: Syntax Scoring](https://adventofcode.com/2021/day/10) - [part 1](day10_p1/main.go), [part 2](day10_p2/main.go)
- [Day 11: Dumbo Octopus](https://adventofcode.com/2021/day/11) - [part 1](day11_p1/main.go), [part 2](day11_p2/main.go)
- [Day 12: Passage Pathing](https://adventofcode.com/2021/day/12) - [part 1](day12_p1/main.go), [part 2](day12_p2/main.go)

  I love the path searching problems. This one is a fairly small set of states,
  so I used a recursive search and kept track of if I had revisited a small
  cave by changing the maximum allowed small cave visits as I called the 
  recursive function.

- [Day 13: Transparent Origami](https://adventofcode.com/2021/day/13) - [part 1](day13_p1/main.go), [part 2](day13_p2/main.go)
- [Day 14: Extended Polymerization](https://adventofcode.com/2021/day/14) - [part 1](day14_p1/main.go), [part 2](day14_p2/main.go)

  I thought this one pretty straightforward, but I thought the trick was clever.
  The polymer itself is too long to track effectively, but each pair gives rise
  to two daugher pairs which share the same new central atom. By tracking the
  pairs it is possible to get sums for each atom pair in the polymer. I saw a
  lot of people counting this one strangely when they went to total up the
  atoms. Many people were summing all letters then dividing by two and accounting
  for each end. I thought it was much easier to count the first atom from each
  pair, knowing that the second atom is the first atom of some other pair. Then
  just add the last atom to the bunch.

- [Day 15: Chiton](https://adventofcode.com/2021/day/15) - [part 1](day15_p1/main.go), [part 2](day15_p2/main.go)

  This is another path search. Generally it is possible to do a breadth first
  search and not rely on algorithms like A* and Dijkstra, but I had never used
  Dijkstra before, so I decided to implement the algorithm and make use of Go's
  heap library. The toughest part of Go's heaps turns out to be figuring out
  when to use and not use pointers, but after you get the hang of them, they are
  pretty straightforward. They definitely speed up the implementation.

- [Day 16: Packet Decoder](https://adventofcode.com/2021/day/16) - [part 1](day16_p1/main.go), [part 2](day16_p2/main.go)

  I loved this problem. I used Go's big integer library to construct the bitfield
  and then just right shifted and masked the section of interest. I left the
  values as big integers just in case there were any gigantic numbers, but it
  was not necessary. I wrote a nice String method which printed the packets
  with levels of indentation to make them more easily readable as well.

- [Day 17: Trick Shot](https://adventofcode.com/2021/day/17) - [part 1](day17_p1/main.go), [part 2](day17_p2/main.go)

  This one sucked for me. I hard coded my solution initially and miscopied my
  lowest y value, so I was getting the wrong answer for a good hour before I
  thought to check my input. *Don't hard code things!* The second part was
  easy because I spent so much time on the first part.

- [Day 18: Snailfish](https://adventofcode.com/2021/day/18) - [part 1](day18_p1/main.go), [part 2](day18_p2/main.go)

  Another favorite. I first built a tree representation of each snailfish
  number, then I built a list with pointers to all number elements up to 4 levels
  and all pairs nested 4 levels. This allowed me to easily navigate left and right
  to do the explodes as well as search through to do the splits.
  I thought my solution was one of my most elegant of this advent.

- [Day 19: Beacon Scanner](https://adventofcode.com/2021/day/19) - [part 1](day19_p1/main.go), [part 2](day19_p2/main.go)

  This was some of my worst code. I found the rotations rather easily, but when
  it came time to find similar distances between two regions I ended up
  calculating metropolis distance, ordered element distances, and in the end I
  think I calculated too much. Better to just go through the rotations. In the
  end it does work though.

- [Day 20: Trench Map](https://adventofcode.com/2021/day/20) - [part 1](day20_p1/main.go), [part 2](day20_p2/main.go)
- [Day 21: Dirac Dice](https://adventofcode.com/2021/day/21) - [part 1](day21_p1/main.go), [part 2](day21_p2/main.go)

  I really liked this problem as well. I think I took a slightly different tack
  recording separately for each player the number of universes produced for
  each player state. The player state included the position on the board, the
  player's score, and the turn number. I calculated all states up to where the
  score first surpassed 21. Then I looked at all winning states (i.e. scores
  above 21) and multiplied them by the other player's losing (less than 21 score)
  states which had the appropriate number of turns. When looking at player 1's
  wins, I would look at losses from the previous player 2's turn. When looking
  at player 2's wins I would look at the losses from player 1's same turn. This
  complication in turns is because player 1 goes first.

- [Day 22: Reactor Reboot](https://adventofcode.com/2021/day/22) - [part 1](day22_p1/main.go), [part 2](day22_p2/main.go)

  I ended up doing a recursive solution of this. The idea is to move up the
  instructions from first to last. For each "on" instruction I added all the
  lights that were never turned on or off in subsequent instructions. For each
  intersecting subsequent instruction I would subtract off those positions
  which never intersect further instructions, and so on. It ended up being
  reasonably fast because there are only a limited number of overlapping
  instruction intersections.
