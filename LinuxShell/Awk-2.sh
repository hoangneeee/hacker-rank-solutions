awk '{
    if ($2 >= 50 && $3 >= 50 && $4 >= 50) {
      sts="Pass"
    } else {
      sts="Fail"
    }
    print $1, ":", sts
}'
