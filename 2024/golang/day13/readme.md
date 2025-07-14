See the `day13` folder in the parent directory of `golang` folder for a solution and a better explanation using `numpy` in python

```
94a+22b=8400
34a+67b=5400
```

```
x1 a + x2 b = tX
y1 a + y2 b = tY
```

```
b = (tY - y1 a) / y2
```

```
y2*x1 a + x2(tY - y1 a) = y2 * tX

y2*x1 a - x2*y1 a = y2 * tX - x2*tY

a = (y2*tX - x2*tY) / (y2*x1 - x2*y1)
```
