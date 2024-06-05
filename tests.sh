#!/bin/bash

INPUT_FILE="txt/sample.txt"
OUTPUT_FILE="txt/result.txt"
GO_PROGRAM="go run ."

verify_output() {
  expected_output="$1"
  echo "$expected_output" > txt/expected_output.txt
  echo "\033[36m$expected_output\033[0m\n"
  if diff -q txt/expected_output.txt $OUTPUT_FILE >/dev/null; then
    echo "\033[32mTest succeeded!\033[0m\n"
  else
    echo "\033[31mTest failed!\033[0m\n"
  fi
  rm txt/expected_output.txt
}

print_io() {
    echo "\033[35m$(cat $INPUT_FILE)\033[0m\n"
    echo "\033[33m$(cat $OUTPUT_FILE)\033[0m"
}

echo "\n////////////////////////TESTS///////////////////////////"
echo "\n\033[35m◼\033[0m : Input \033[33m◼\033[0m : Output \033[36m◼\033[0m : Expected output"
echo "\nTest 1 :\n"
echo "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair." > $INPUT_FILE
$GO_PROGRAM $INPUT_FILE $OUTPUT_FILE
print_io
verify_output "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair."

echo "//////////////////////////////////////////////////////////\n"
echo "Test 2 :\n"
echo "Simply add 42 (hex) and 10 (bin) and you will see the result is 68." > $INPUT_FILE
$GO_PROGRAM $INPUT_FILE $OUTPUT_FILE
print_io
verify_output "Simply add 66 and 2 and you will see the result is 68."

echo "//////////////////////////////////////////////////////////\n"
echo "Test 3 :\n"
echo "There is no greater agony than bearing a untold story inside you." > $INPUT_FILE
$GO_PROGRAM $INPUT_FILE $OUTPUT_FILE
print_io
verify_output "There is no greater agony than bearing an untold story inside you."

echo "//////////////////////////////////////////////////////////\n"
echo "Test 4 :\n"
echo "Punctuation tests are ... kinda boring ,don't you think !?" > $INPUT_FILE
$GO_PROGRAM $INPUT_FILE $OUTPUT_FILE
print_io
verify_output "Punctuation tests are... kinda boring, don't you think!?"

echo "//////////////////////////////////////////////////////////\n"
echo "Test 5 :\n"
echo "This sentence specifically test (cap, 3) for arguments in flags (up, 3) . DO YOU THINK YOU (low, 4) GET IT RIGHT ?" > $INPUT_FILE
$GO_PROGRAM $INPUT_FILE $OUTPUT_FILE
print_io
verify_output "This Sentence Specifically Test for ARGUMENTS IN FLAGS. do you think you GET IT RIGHT?"

echo "//////////////////////////////////////////////////////////\n"
echo "Test 6 :\n"
echo "Here is a ' difficult one ', because you'll need to 'manage perfectly ' the quotes, and yes ' some times' it's ' better erase all your code and start again'." > $INPUT_FILE
$GO_PROGRAM $INPUT_FILE $OUTPUT_FILE
print_io
verify_output "Here is a 'difficult one', because you'll need to 'manage perfectly' the quotes, and yes 'some times' it's 'better erase all your code and start again'."

echo "//////////////////////////////////////////////////////////\n"
echo "Test 7 :\n"
echo "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?" > $INPUT_FILE
$GO_PROGRAM $INPUT_FILE $OUTPUT_FILE
print_io
verify_output "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"

echo "//////////////////////////////////////////////////////////\n"
echo "Test 8 :\n"
echo "I have' to pack' 101 (bin) outfits. Packed 1a (hex) just to be sure" > $INPUT_FILE
$GO_PROGRAM $INPUT_FILE $OUTPUT_FILE
print_io
verify_output "I have 'to pack' 5 outfits. Packed 26 just to be sure"

echo "//////////////////////////////////////////////////////////\n"
echo "Test 9 :\n"
echo "Don not be sad ,because     sad backwards is das . And das not good" > $INPUT_FILE
$GO_PROGRAM $INPUT_FILE $OUTPUT_FILE
print_io
verify_output "Don not be sad, because sad backwards is das. And das not good"

echo "//////////////////////////////////////////////////////////\n"
echo "Test 10 :\n"
echo "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '" > $INPUT_FILE
$GO_PROGRAM $INPUT_FILE $OUTPUT_FILE
print_io
verify_output "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"
