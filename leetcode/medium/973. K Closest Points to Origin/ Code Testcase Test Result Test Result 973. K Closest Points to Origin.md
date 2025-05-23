# 973. [K ближайших точек к началу координат (K Closest Points to Origin)]()

`Medium`

Дан массив `points`, где `points[i] = [xi, yi]` представляет точку на плоскости **X-Y**, и целое число `k`. Верните `k` ближайших точек к началу координат `(0, 0)`.

Расстояние между двумя точками на плоскости **X-Y** — это евклидово расстояние (т.е. √(x₁ - x₂)² + (y₁ - y₂)²).

Вы можете вернуть ответ в **любом порядке**. Ответ **гарантированно** будет **уникальным** (за исключением порядка, в котором он находится).

**Пример 1:**\
![](https://assets.leetcode.com/uploads/2021/03/03/closestplane1.jpg)
```
Input: points = [[1,3],[-2,2]], k = 1
Output: [[-2,2]]
Explanation:
Расстояние между (1, 3) и началом координат равно sqrt(10).
Расстояние между (-2, 2) и началом координат равно sqrt(8).
Поскольку sqrt(8) < sqrt(10), (-2, 2) ближе к началу координат.
Нам нужны только k = 1 ближайшая точка к началу координат, поэтому ответ просто [[-2,2]].
```

**Пример 2:**
```
Input: points = [[3,3],[5,-1],[-2,4]], k = 2
Output: [[3,3],[-2,4]]
Explanation: Ответ `[[-2,4],[3,3]]` также будет принят.
```

**Ограничения:**

*   1 ≤ `k` ≤ `points.length` ≤ 10⁴
*   -10⁴ ≤ `xi`, `yi` ≤ 10⁴