#!/usr/bin/env bash


TODAY=$(date +%Y-%m-%d-%HH-%MM-%SS)

VER="LTK"
VER+="_v0.9"
VER+="_1st_commit."
VER+="_"$TODAY
echo $VER

git tag -f $VER -m $VER"Release v0.9" 

git push --tags

git add .
git commit -a -m $VER
git push --set-upstream origin main

# ###################