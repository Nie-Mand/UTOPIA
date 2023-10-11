#!/bin/bash

for i in {1..999}; do 
    x=$(uuidgen | md5sum | cut -d' ' -f1)
    y=$(uuidgen | md5sum | cut -d' ' -f1)
    echo $x
    echo $x > $y
done