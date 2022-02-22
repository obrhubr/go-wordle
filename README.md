# Go-Wordle (or Gordle)

This was inspired by https://www.youtube.com/watch?v=v68zYyaEmEA. I used 3B1B's algorithm to create a similar bot.

```
     Wordle Bot
--------------------- 
 
 
-------------------------------
|  word   |      score        |
-------------------------------
| trace   | 5.377566338603813 | 
| salet   | 5.303611974124567 | 
| parse   | 5.294365346728447 | 
| crate   | 5.287056928899425 | 
| carte   | 5.230307126327275 | 
| peart   | 5.223432537531815 | 
| slate   | 5.221421829329374 | 
| reast   | 5.2029841856220935 | 
| caret   | 5.186305073551834 | 
| carle   | 5.172858791190428 | 
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