#!/bin/bash

for dir in dist/*/; do
	mkdir $dir/contrib
        cp {README.md,LICENSE} $dir
        cp contrib/*.tpl $dir/contrib
done
