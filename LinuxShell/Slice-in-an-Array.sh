while read line
do
    array=("${array[@]}" $line)
done

arr="${array[@]:3:5}"
echo ${arr[@]}
