# Camping trip dynamic programming exercise

Suppose you’re going camping. You have a knapsack that will hold 6 lb, and you can take the following items.
Each has a value, and the higher the value, the more important the item is:

- Water, 3 lb, 10
- Book, 1 lb, 3
- Food, 2 lb, 9
- Jacket, 2 lb, 5
- Camera, 1 lb, 6

What’s the optimal set of items to take on your camping trip?

## Solution

Draw the grid first:

|        | 1lb | 2lb | 3lb | 4lb | 5lb | 6ln |
|--------|-----|-----|-----|-----|-----|-----|
| water  |     |     |     |     |     |     |
| book   |     |     |     |     |     |     |
| food   |     |     |     |     |     |     |
| jacket |     |     |     |     |     |     |
| camera |     |     |     |     |     |     |

Now, calculate according to this rule:

```
                      1. the previous max(value at cell[i-1][j])
cell[i][j] = max of {
                      2. value of current item + value of remaining space (cell[i-1][j-item's weight]) (( exp: 3 - 1(1lb of the item and 3 is the current weight we are looking at ))
```

|             | 1lb  | 2lb  | 3lb     | 4lb     | 5lb       | 6ln       |
|-------------|------|------|---------|---------|-----------|-----------|
| water(3,10) | 0    | 0    | 10(w)   | 10(w)   | 10(w)     | 10(w)     |
| book(1,3)   | 3(b) | 3(b) | 10(w)   | 13(w,b) | 13(w,b)   | 13(w,b)   |
| food(2,9)   | 3(b) | 9(f) | 10(w)   | 13(w,b) | 19(w,f)   | 21(w,b,f) |
| jacket(2,5) | 3(b) | 9(f) | 10(w)   | 14(f,j) | 19(w,f)   | 21(w,b,f) |
| camera(1,6) | 6(c) | 9(f) | 15(f,c) | 16(w,c) | 20(f,j,c) | 25(w,f,c) |

- Dynamic programming is useful when you’re trying to optimize something given a constraint.
- You can use dynamic programming when the problem can be broken into discrete sub-problems, and they don’t depend on each other.
- Every dynamic-programming solution involves a grid.
- The values in the cells are usually what you’re trying to optimize. For the knapsack problem, the values were the value of the goods.
- Each cell is a sub-problem, so think about how you can divide your problem into sub-problems. That will help you figure out what the axes are.
