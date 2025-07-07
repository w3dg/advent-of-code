import { readFileSync } from "fs";

import { checkVisibility, calculateScenicScore } from "./utils";

// const input = readFileSync(__dirname + "/sample.txt", "utf8");
const input = readFileSync(__dirname + "/input.txt", "utf8");

const grid = input.split("\n").map((line) => line.split("").map((x) => Number(x)));

function puzzle1() {
  let visibilityCounter = 0;

  for (let i = 0; i < grid.length; i++) {
    for (let j = 0; j < grid[i].length; j++) {
      if (checkVisibility(grid, i, j)) {
        visibilityCounter++;
      }
    }
  }

  return visibilityCounter;
}

function puzzle2() {
  let scenicScores = [];
  for (let i = 0; i < grid.length; i++) {
    for (let j = 0; j < grid[i].length; j++) {
      scenicScores.push(calculateScenicScore(grid, i, j));
    }
  }

  scenicScores = scenicScores.filter((x) => x !== -1);
  return Math.max(...scenicScores);
}

console.log("Puzzle 1: ", puzzle1());
console.log("Puzzle 2: ", puzzle2());
