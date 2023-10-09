function main() {
    for (( i = 1; i <= 99; i++))
        do
            local check=$((i % 2))
            if [ $check != 0 ]
            then
                echo $i
            fi
        done
}

main
