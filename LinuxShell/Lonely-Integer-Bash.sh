read len
read -a array

declare -A hash_map

for ele in "${array[@]}"; do
    if [[ -v hash_map["$ele"] ]]; then
        unset hash_map["$ele"]
    else
        hash_map["$ele"]="$ele"
    fi
done

echo ${hash_map[@]}
