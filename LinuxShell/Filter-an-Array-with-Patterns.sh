array=($(cat))
for ele in "${array[@]}"; do
    if [[ "$ele" != *A* && "$ele" != *a* ]]; then
        echo $ele
    fi
done
