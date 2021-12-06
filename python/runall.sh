#!/bin/sh

for day in $(seq -f "%02g" 1 25); do
  filename="day${day}.py"
  if [ -f $filename ]; then
    echo "Day ${day}:"
    python $filename
    echo ""
  fi
done
