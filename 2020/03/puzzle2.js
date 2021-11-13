/*

--- Part Two ---
Time to check the rest of the slopes - you need to minimize the probability of a sudden arboreal stop, after all.

Determine the number of trees you would encounter if, for each of the following slopes, you start at the top-left corner and traverse the map all the way to the bottom:

Right 1, down 1.
Right 3, down 1. (This is the slope you already checked.)
Right 5, down 1.
Right 7, down 1.
Right 1, down 2.
In the above example, these slopes would find 2, 7, 3, 4, and 2 tree(s) respectively; multiplied together, these produce the answer 336.

What do you get if you multiply together the number of trees encountered on each of the listed slopes?
*/

const fs = require("fs");

// const input = fs.readFileSync("./2020/03/sample.txt", "utf8").split("\r\n");
const input = fs.readFileSync("./2020/03/input.txt", "utf8").split("\r\n");

const inputLength = input.length;

// let y_coordinate = 0; its basically the loop iterable variable

function determineTreesPerSlope(rightOffset, downOffset) {
  let count = 0;
  let x_coordinate = 1;
  for (let i = 0; i < inputLength - downOffset; ) {
    x_coordinate =
      x_coordinate + rightOffset > input[i].length
        ? (x_coordinate + rightOffset) % input[i].length
        : x_coordinate + rightOffset;
    i = i + downOffset;
    if (input[i][x_coordinate - 1] === "#") {
      count++;
    }
  }
  return count;
}

let slope1 = determineTreesPerSlope(1, 1);
let slope2 = determineTreesPerSlope(3, 1);
let slope3 = determineTreesPerSlope(5, 1);
let slope4 = determineTreesPerSlope(7, 1);
let slope5 = determineTreesPerSlope(1, 2);

console.log(slope1, slope2, slope3, slope4, slope5);
console.log(slope1 * slope2 * slope3 * slope4 * slope5);
