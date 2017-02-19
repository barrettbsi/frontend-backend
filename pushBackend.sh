#!/bin/bash
cp Procfile.backend Procfile
git add -A
rm Procfile
git commit -m "$1"
git push backend master