
read integers

total=0
for (( i=0; i<$integers; i++ )); do
  read num
  total=$((num + total))
done

result=$(echo "scale=4; $total / $integers" | bc)

printf "%.3f\n" "$result"

