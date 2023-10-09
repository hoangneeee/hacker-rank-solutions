while read line;
do
    echo "${line}" | cut -c3 | cut -c7
done
