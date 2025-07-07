"""
This seems to be a linear equation solving problem
given 2 equations with 2 unknown variables, 
we need to find what are the variables to get to the the destination

for eg.
Button A: X+94, Y+34 : here: 94a, 34a
Button B: X+22, Y+67 : here: 22b,  67b
Prize: X=8400, Y=5400 : here: 8400  5400

94a + 22b = 8400
34a + 67b = 5400

(all are column vectors and hence the transpose)

[[94, 22], [34,67]]T [a, b]T = [8400, 5400]T

we need to solve for a and b if there exists a possible solution
also we cannot press a button for -ve times, so if the solution is negative, 
we reject it
if its not in whole numbers, we reject it

we can get the solution and check it backwards to see if it fits after rounding
as we do not want any fractional solutions

numpy has matrices and a linear algebra solver
"""

import numpy as np
import math

file = "./sample.txt"
# file = "./input.txt"
contents = ""
with open(file) as f:
    contents = f.read().strip()

machines = contents.split("\n\n")

p1 = 0
p2 = 0


def cost(a, b):
    return int(3 * a + b)


for m in machines:
    a, b, prize = m.split("\n")
    ax, ay = [
        int(a.split(",")[0].split()[-1].replace("X+", "")),
        int(a.split(",")[1].replace(" Y+", "")),
    ]
    bx, by = [
        int(b.split(",")[0].split()[-1].replace("X+", "")),
        int(b.split(",")[1].replace(" Y+", "")),
    ]
    px, py = [
        int(prize.split(",")[0].split()[-1].replace("X=", "")),
        int(prize.split(",")[1].replace(" Y=", "")),
    ]
    CORRECTION = 10000000000000
    M = np.array([[ax, bx], [ay, by]])
    R = np.array([px, py])
    R2 = np.array([px + CORRECTION, py + CORRECTION])
    soln = np.rint(
        np.linalg.solve(M, R)
    )  # rint for rounding the array elements to integers
    soln2 = np.rint(np.linalg.solve(M, R2))

    # checking if after rounding, we get to the original price
    # as we do not want floating values as solutions, we cross check the rounded-off solutions
    check_prize = M @ soln
    
    if np.all(M @ soln == R):
        p1 += cost(*soln)
    if np.all(M @ soln2 == R2):
        p2 += cost(*soln2)

print(p1)
print(p2)
