cat input-partial.txt | sed -E 's/\|/ /g' | tsort
