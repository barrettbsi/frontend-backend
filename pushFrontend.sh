#!/bin/bash
cp Procfile.frontend Procfile
git add -A
rm Procfile
git commit -m "$1"
git push heroku master