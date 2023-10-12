awk '{
    printf ($0)
    printf NR
    if (NR%2 == 0) printf "\n"
    else printf ";"
}'
