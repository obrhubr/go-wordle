# Go-Wordle (or Gordle)

This was inspired by https://www.youtube.com/watch?v=v68zYyaEmEA. I used 3B1B's algorithm to create a similar bot.

### Usage

Run `go run ./cmd/game/main.go` and enter the word you chose. Then the program will prompt you for the pattern that you got back. A 2 is a green letter, a 1 is a yellow letter and a 0 is a gray letter. If you entered `thorn` for example, and the `t` (first letter) was green, the `h` (second letter) was yellow and the rest were gray you would enter: `2 1 0 0 0`.

```
     Wordle Bot
--------------------- 
 
 
-------------------------------
|  word   |      score        |
-------------------------------
| tares   | 8.464541687215732 | 
| teras   | 8.384621063737569 | 
| pelas   | 8.330901493791202 | 
| peats   | 8.246791141310844 | 
| tears   | 8.134442220668847 | 
| bares   | 8.085543735774847 | 
| dores   | 8.084691108022328 | 
| pears   | 8.082520569242503 | 
| pores   | 8.055957327018294 | 
| peart   | 8.027039950859331 | 
---------------------------
Enter the word you have chosen 
-> trace
Enter the pattern (ints seperated by spaces; 0 is gray, 1 is yellow and 2 is green; ex: 1 0 0 0 2) 
-> 2 1 0 0 0
There are 39 words left. 
-------------------------------
|  word   |      score        |
-------------------------------
| thorn   | 4.100451478180524 | 
| torus   | 4.087196023361733 | 
| tumor   | 3.8959397224821326 | 
| turps   | 3.848931434320989 | 
| thrid   | 3.7888536092792475 | 
| third   | 3.7618513684596655 | 
| tiros   | 3.7435893955388577 | 
| thorp   | 3.7211265853645457 | 
| torsi   | 3.7150004004184036 | 
| thirl   | 3.7090864102188643 | 
---------------------------
Enter the word you have chosen 
-> thorn
Enter the pattern (ints seperated by spaces; 0 is gray, 1 is yellow and 2 is green; ex: 1 0 0 0 2) 
-> 2 2 2 2 2
The secret word is: thorn
```
