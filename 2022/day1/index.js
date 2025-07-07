// the number of Calories each Elf is carrying (your puzzle input)

import { readFileSync } from "fs";

const totalElfInventoryCollection = readFileSync(__dirname + "/sample.txt", "utf8").split("\n\n");
// const totalElfInventoryCollection = readFileSync(__dirname + "/input.txt", "utf8").split("\n\n");

const eachElfInventoryCollection = totalElfInventoryCollection.map((line) => {
  return line.split("\n").map((x) => Number(x));
});

console.log(totalElfInventoryCollection);
console.log(eachElfInventoryCollection);

function puzzle1() {
  const totalCaloriesPerElf = eachElfInventoryCollection.map((inventory) => {
    return inventory.reduce((prev, curr) => prev + curr, 0);
  });

  return Math.max(...totalCaloriesPerElf);
}

function puzzle2() {
  const totalCaloriesPerElf = eachElfInventoryCollection.map((inventory) => {
    return inventory.reduce((prev, curr) => prev + curr, 0);
  });

  const sortedCalories = totalCaloriesPerElf.sort((a, b) => b - a); // desc sort
  return sortedCalories[0] + sortedCalories[1] + sortedCalories[2];
}

console.log("Puzzle 1: ", puzzle1());
console.log("Puzzle 2: ", puzzle2());
