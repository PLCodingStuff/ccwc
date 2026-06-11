# Create large_generic.txt
cat benchmarks/large_generic_part_1.txt benchmarks/large_generic_part_2.txt >> benchmarks/large_generic.txt

# Linux wc
echo "========================"
echo "wc"
echo "========================"
time wc -l -w -c benchmarks/large_generic.txt

time wc -l -w -c benchmarks/swarm/*.txt
time wc -l -w -c benchmarks/unicode.txt


# Your Go tool
echo "========================"
echo "ccwc"
echo "========================"
time ./ccwc -l -w -c benchmarks/large_generic.txt

# Your Go tool
time ./ccwc -l -w -c benchmarks/swarm/*.txt

# Your Go tool
time ./ccwc -l -w -c benchmarks/unicode.txt

rm benchmarks/large_generic.txt