import { readFileSync } from "node:fs";

// const input = readFileSync("./sample1.txt", "utf8");
// const input = readFileSync("./sample2.txt", "utf8");
const input = readFileSync("./input.txt", "utf8");

const trailHeadsIndices = [];
const grid = input.split("\n").map((row, rowIndex) => {
  return row.split("").map((x, colIndex) => {
    if (x == "0") trailHeadsIndices.push([rowIndex, colIndex]);
    return Number(x);
  });
});

function applyDFS(curr, countRepeated = false) {
  let score = 0;
  const [row, col] = curr;

  let stack = [`${row}_${col}`];
  let visited = [];

  while (stack.length != 0) {
    let indexStr = stack.pop();
    visited.push(indexStr);
    const row = Number(indexStr.split("_")[0]);
    const col = Number(indexStr.split("_")[1]);
    let val = grid[row][col];
    if (val == 9) {
      score += 1;
    } else {
      // generate valid neighbours
      // generally we do not want to count already visited neighbors so we check,
      // if that is turned on so that we actually do check, we essentially turn off that check by AND with a True.
      // Logically, for any proposition p, (p AND True) === p
      if (
        col - 1 >= 0 &&
        grid[row][col - 1] == val + 1 &&
        stack.indexOf(`${row}_${col - 1}`) == -1 &&
        (countRepeated === false ? visited.indexOf(`${row}_${col - 1}`) == -1 : true)
      )
        stack.push(`${row}_${col - 1}`);
      if (
        col + 1 < grid[0].length &&
        grid[row][col + 1] == val + 1 &&
        stack.indexOf(`${row}_${col + 1}`) == -1 &&
        (countRepeated === false ? visited.indexOf(`${row}_${col + 1}`) == -1 : true)
      )
        stack.push(`${row}_${col + 1}`);
      if (
        row - 1 >= 0 &&
        grid[row - 1][col] == val + 1 &&
        stack.indexOf(`${row - 1}_${col}`) == -1 &&
        (countRepeated === false ? visited.indexOf(`${row - 1}_${col}`) == -1 : true)
      )
        stack.push(`${row - 1}_${col}`);
      if (
        row + 1 < grid.length &&
        grid[row + 1][col] == val + 1 &&
        stack.indexOf(`${row + 1}_${col}`) == -1 &&
        (countRepeated === false ? visited.indexOf(`${row + 1}_${col}`) == -1 : true)
      )
        stack.push(`${row + 1}_${col}`);
    }
  }

  return score;
}

function puzzle() {
  let sumOfScoresP1 = 0;
  let sumOfScoresP2 = 0;
  trailHeadsIndices.forEach((idx) => {
    // console.log(idx);
    sumOfScoresP1 += applyDFS(idx);
    sumOfScoresP2 += applyDFS(idx, true);
  });

  console.log("Puzzle 1 is:", sumOfScoresP1);
  console.log("Puzzle 2 is:", sumOfScoresP2);
}

puzzle();
