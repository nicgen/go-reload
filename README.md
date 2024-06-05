# Go-reloaded

## Usage

Clone the repo:

```
git clone https://zone01normandie.org/git/ngenty/go-reloaded
cd go-reloaded
```

create a sample file (for ex: sample.txt)

```
harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '
```

execute the code with:

```
go run . "sample.txt" "output.txt"
```

## Objectives

> In this project you will use some of your old functions made in your old repository. You will use them with the objective of making a simple text completion/editing/auto-correction tool.

- [x] ~~Your project must be written in Go.~~
- [x] ~~The code should respect the [good practices](https://github.com/01-edu/public/blob/master/subjects/good-practices/README.md).~~
- [x] ~~(optional) It is recommended to have test files for [unit testing](https://go.dev/doc/tutorial/add-a-test).~~

- tool receive (as arguments):
    - the name of a file containing a text that needs modifications (input)
    - the name of the file the modified text should be placed in (the output)
- list of possible modifications that your program should execute:
    - conversion
        - `(hex)` -> replace the word before with its decimal version
            - note: the word will always be a **hexadecimal number**
            - ex:  `1E (hex)` -> 30; `1A` -> 26; `42` -> 66
        - `(bin)` -> replace the word before with its decimal version
            - note:  the word will always be a **binary number**
            - ex: `10 (bin)` -> 2; `101` -> 5
    - texte transformation
        - `(up, <num>?)` -> converts the word before with its Uppercase version
            - ex: "Ready, set, go (up) !" -> "Ready, set, GO !"
        - `(low, <num>?)` ->  converts the word before with its Lowercase version
            - ex: Ex: "I should stop SHOUTING (low)" -> "I should stop shouting"
        - `(cap, <num>?)` -> converts the word before with its capitalized version
            - ex:  "Welcome to the Brooklyn bridge (cap)" -> "Welcome to the Brooklyn Bridge"
        Note: For (low), (up), (cap) if a number appears next to it, like so: (low, <number>) it turns the previously specified number of words in lowercase, uppercase or capitalized accordingly. (Ex: "This is so exciting (up, 2)" -> "This is SO EXCITING")
    - punctuation
        - `.` `,` `!` `?` `:` `;` no space(s) before, one space after
        - note: group of multiples punctuation exists `...` or `!?` (or else)
        - `'` always surround a word or a group of words
    - vowels
        - `a` -> an if next word begins with a vowel `a`,`e`,`i`,`o`,`u`
            ex:  "There it was. A amazing rock!" -> "There it was. An amazing rock!

## tests

`it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.`  
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.

`Simply add 42 (hex) and 10 (bin) and you will see the result is 68.`  
Simply add 66 and 2 and you will see the result is 68.

`There is no greater agony than bearing a untold story inside you.`  
There is no greater agony than bearing an untold story inside you.

`Punctuation tests are ... kinda boring ,don't you think !?`  
Punctuation tests are... kinda boring, don't you think!?

audit (github)

`If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?`  
If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?

`I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure`  
I have to pack 5 outfits. Packed 26 just to be sure

`Don not be sad ,because sad backwards is das . And das not good`  
Don not be sad, because sad backwards is das. And das not good

`harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '`  
Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'

Note: :warning: a possible trick is to add multiple transformation to a word or group of words

[source](https://github.com/01-edu/public/tree/master/subjects/go-reloaded)
