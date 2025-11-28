## Notes

### Day 2
- Heh, the input list is sorted. All the unsafe reports comes first, and the all the safe reports come after.

### Day 4  
The test_input looks like this and contains 18 "XMAS"

```
|MMMSXXMASM| |....XXMAS.|
|MSAMXMSMSA| |.SAMXMS...|
|AMXSXMAAMM| |...S..A...|
|MSAMASMSMX| |..A.A.MS.X|
|XMASAMXAMM| |XMASAMX.MM|
|XXAMMXXAMA| |X.....XA.A|
|SMSMSASXSS| |S.S.S.S.SS|
|SAXAMASAAA| |.A.A.A.A.A|
|MAMMMXMMMM| |..M.M.M.MM|
|MXMXAXMASX| |.X.X.XMASX|

Horizontal right     3
Horizontal left      2
Vertical down        1
Vertical up          2
Diagonal right up    4
Diagonal right down  1
Diagonal left up     4
Diagonal left down   1
SUM                 18
```

Will probably just brute-force by running through the input 1 time for each direction.

I think masks could be used to allow for only running through the input once, but then I would also need to keep track of which matches that have already been found to avoid doubles, so just running through one direction at the time will probably be easier.

To avoid doubles on the diagonal, I think I should just search `right-down` + `right-up` (both for `XMAS` and `SAMX`. That should cover all diagonal while avoiding doubles, right?)