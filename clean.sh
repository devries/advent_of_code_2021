#!/bin/sh

for day in $(seq -f "%02g" 1 25); do
  for part in 1 2; do
    dir="day${day}_p${part}"
    if [ -d $dir ]; then
      cwd=$(pwd)
      cd $dir
      rm $dir
      cd $cwd
    fi
  done
done
