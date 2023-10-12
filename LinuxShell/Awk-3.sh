awk '{
    arg=($2+$3+$4)/3

    if (arg >= 80) {
      sts="A"
    } else if (arg >= 60) {
      sts="B"
    } else if (arg >= 50) {
      sts="C"
    } else {
      sts="FAIL"
    }
    print $0, ":", sts
}'
