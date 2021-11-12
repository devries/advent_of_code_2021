#!/bin/sh

for day in $(seq -f "%02g" 1 25); do
  for part in 1 2; do
    dir="day${day}_p${part}"
    if [ -d $dir ]; then
      echo "Compiling $dir"
      cwd=$(pwd)
      cd $dir
      go build -v .
      cd $cwd
    fi
  done
done
