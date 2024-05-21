#! /usr/bin/env bash

# todo create a list of tests
# input filename, output filename, input content, output content (result)
# create variables, pass, error
# find _test files, pause script and execute them in new tab
# find ../ -type f -iname "*_test.go"

# TESTS
# USAGE:
# input filename, output filename, raw text, solution

# tests
t_0=( "sample.txt" "output.txt" "one" "one" )
t_1=( "sample.txt" "output.txt" "two" "two" )
t_2=( "sample.txt" "output.txt" "three" "three" )
t_3=( "sample.txt" "output.txt" "four" "three" )
# main file
t_00=( "sample.txt" "output.txt" "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair." "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair." )
t_01=( "sample.txt" "output.txt" "Simply add 42 (hex) and 10 (bin) and you will see the result is 68." "Simply add 66 and 2 and you will see the result is 68." )
t_02=( "sample.txt" "output.txt" "There is no greater agony than bearing a untold story inside you." "There is no greater agony than bearing an untold story inside you." )
t_03=( "sample.txt" "output.txt" "Punctuation tests are ... kinda boring ,don't you think !?" "Punctuation tests are... kinda boring, don't you think!?" )
# audit file
t_04=( "sample.txt" "output.txt" "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?" "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?" )
t_05=( "sample.txt" "output.txt" "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure" "I have to pack 5 outfits. Packed 26 just to be sure" )
t_06=( "sample.txt" "output.txt" "Don not be sad ,because sad backwards is das . And das not good" "Don not be sad, because sad backwards is das. And das not good" )
t_07=( "sample.txt" "output.txt" "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '" "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'" )
# optional
t_08=( "input.txt" "output.txt" "My NAME is TIMMY,(up, 3)(low, 3)(up) I am 65 (hex)(bin) years old and I like watching ' spongebob squarepants   '(cap)." "My name is TIMMY, I am 5 years old and I like watching 'Spongebob Squarepants'" )
# arrays
arr1=(t_00 t_01 t_02 t_03 t_04 t_05 t_06 t_07 t_08)
# test
# arr1=(t_0 t_1 t_2 t_3)
arr=(arr1)

# TESTS
test_total=${#arr1[@]} # number of tests
test_actual=0
test_passed=0

# COLORS
WHITE="\e[37m"
BG_GREEN="\e[42m"
BG_RED="\e[41m"
ERASE="$\033[0K\r"

# FUNCTIONS
printTitle(){
cat << EOF
┌───────────────────────┐
│  GO-RELOAD            │
├───────────────────────┤
│  HellTests by nicgen  │
└───────────────────────┘
EOF
}

test(){
    # echo "$1"
    # cat "$1"
    # echo "$2"
    # cat "$2"
    test=$(comm -3 "$1" "$2")
    if [ -z  "$test" ]; then
        ((test_passed++))
        echo -ne " $test_actual/$test_total - PASS\033[0K\r"
        sleep 1
    else
        # echo -ne "$test_actual/$test_total - ERROR\033[0K\r"
        echo -ne " $test_actual/$test_total - ERROR$ERASE"
        sleep 1
        # echo "${test[@]}"
        echo -ne "Error in $test_actual/$test_total:\nGOT: \t $(cat $1)\nWANTED:  $(cat $2)\n\n"
        # echo -e "$(comm -23 $1 $2)\n"
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
            # echo -e "──────────────────────────────────────────"
        done
    done
    echo -e "Score ($test_passed/$test_total)\n"
}

# EXECUTE
printTitle
main
