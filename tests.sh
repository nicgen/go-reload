#! /usr/bin/env bash

# todo create a list of tests
# input filename, output filename, input content, output content (result)
# create variables, pass, error
# find _test files, pause script and execute them in new tab
# find ../ -type f -iname "*_test.go"

# VARS
WHITE="\e[37m"
BG_GREEN="\e[42m"
BG_RED="\e[41m"

# TESTS
# input filename, output filename, raw text, solution
test_00=( "sample.txt" "output.txt" "This is the 1st test" "This is the 1st test" )
test_01=( "sample.txt" "output.txt" "This is the 2nd test" "This IS the 2nd test" )
test_02=( "sample.txt" "output.txt" "This is the 3rd test" "This IS the 3rd test" )
test_03=( "sample.txt" "output.txt" "This is the 4th test" "This IS the 4th test" )
test_04=( "sample.txt" "output.txt" "This is the 5th test" "This is the 5th test" )

arr1=(test_00 test_01 test_02 test_03 test_04)

arr=(arr1)

printTitle(){
cat << EOF
┌───────────────────────┐
│  GO-RELOAD            │
├───────────────────────┤
│  HellTests by nicgen  │
└───────────────────────┘
EOF
}

test_total=${#arr1[@]}
test_actual=0
test_passed=0

# tests batch
test(){
    echo "$1"
    cat "$1"
    echo "$2"
    cat "$2"
    test=$(comm -3 "$1" "$2")
    if [ -z  "$test" ]; then
        ((test_passed++))
        echo -ne "$test_actual/$test_total - ${WHITE}${BG_GREEN}PASS${NC}\033[0K\r"
        sleep 1
    else
        echo -ne "$test_actual/$test_total - ${WHITE}${BG_RED}ERROR${NC}\033[0K\r"
        sleep 1
        # echo "${test[@]}"
        echo -ne "Your result VS what was asked:\n$(cat $1)\n$(cat $2)"
        echo -e "$(comm -23 $1 $2)\n"
    fi
    # echo -e "──────────────────────────────────────────"
}

main() {
    declare -n elemnt elemnt_01
    for elemnt in "${arr[@]}"; do
        # echo "${elemnt[@]}"
        for elemnt_01 in "${elemnt[@]}"; do
            ((test_actual++)) # increment actual test
            # echo "${elemnt_01[@]}"
            fileIN="${elemnt_01[0]}"
            fileOUT="${elemnt_01[1]}"
            txtIN="${elemnt_01[2]}"
            txtOUT="${elemnt_01[3]}"
            echo "$txtIN" > "$fileIN"
            echo "$txtOUT" > "$fileOUT"
            # sleep 1
            test "$fileIN" "$fileOUT"
            # sleep 1
            rm "$fileIN" "$fileOUT"
            echo -e "──────────────────────────────────────────"
        done
    done
    echo -r "You succeded ($test_passed/$test_total)"
}
main
