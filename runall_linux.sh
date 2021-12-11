#!/bin/sh

fullstart=$(date -u +%s.%N)
for day in $(seq -f "%02g" 1 25); do
  finecho=0
  start=$(date -u +%s.%N)
  for part in 1 2; do
    dir="day${day}_p${part}"
    if [ -d $dir ]; then
      if [ $part -eq 1 ]; then
        printf "Day $day "
        finecho=1
      fi
      cwd=$(pwd)
      cd $dir
      printf "part $part: "
      retval="$(./$dir)"
      printf "$retval "
      cd $cwd
    fi
  done
  if [ $finecho -eq 1 ]; then
    end=$(date -u +%s.%N)
    #duration=$( expr $end - $start)
    duration=$(echo "scale = 3; ($end - $start)/1.0" | bc)
    printf "time elapsed: ${duration}s\n"
    #if [ $duration -eq 0 ]; then
    #  printf "time elapsed: <1s\n"
    #else
    #  printf "time elapsed: ${duration}s\n"
    #fi
  fi
done
fullend=$(date -u +%s.%N)
duration=$(echo "scale = 3; ($fullend-$fullstart)/1.0" | bc)
echo ""
echo "Total time elapsed: ${duration}s"
