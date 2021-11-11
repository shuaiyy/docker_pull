#!/usr/bin/env bash

#./push_amd64 current_image_list.txt

for image in `cat ./current_image_list.txt`
do
   echo $image
   tag=`./sha256_amd64 $image`
done