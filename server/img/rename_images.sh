#!/bin/bash

count=1

for file in *.jpg; do
  new_name="img${count}.jpg"
  mv "$file" "$new_name"
  count=$((count+1))

  # Sal del bucle cuando hayas renombrado 10 archivos
  if [ $count -gt 10 ]; then
    break
  fi
done
