import { readFileSync } from "node:fs";

// const input = readFileSync("./sample.txt", "utf8");
const input = readFileSync("./input.txt", "utf8");

const lines = input.split("\n");

const list1 = [];
const list2 = [];

lines.forEach((line) => {
  const numbers = line
    .split(" ")
    .filter((x) => x.trim().length != 0)
    .map((x) => Number(x));
  list1.push(numbers.at(0));
  list2.push(numbers.at(-1));
});

list1.sort((a, b) => a - b);
list2.sort((a, b) => a - b);

// console.log(list1, list2);

function puzzle1() {
  let totalDistance = 0;
  for (let i = 0; i < list1.length; i++) {
    const distance = Math.abs(list1[i] - list2[i]);
    totalDistance += distance;
  }
  return totalDistance;
}

console.log("Puzzle 1 is:", puzzle1());

// Puzzle 2
// Bruteforce O(n^2), and maybe optimised with a hashmap and 2 pointers

function puzzle2() {
  let similarity = 0;

  for (let i = 0; i < list1.length; i++) {
    const currI = list1[i];
    let multiplier = 0;
    for (let j = 0; j < list2.length; j++) {
      if (currI == list2[j]) {
        multiplier++;
      }
    }

    similarity += currI * multiplier;
  }
  return similarity;
}
// console.log("Puzzle 2", puzzle2());

function puzzle2Again() {
  let pl1 = 0;
  let pl2 = 0;

  let similarity = 0;
  let countMap = new Map();
  while (pl1 < list1.length && pl2 < list2.length) {
    let currEl = list1[pl1];
    if (countMap.get(currEl) != undefined) {
      similarity += currEl * countMap.get(currEl);
    } else {
      let c = 0;
      while (pl2 < list2.length && list2[pl2] < currEl) {
        pl2++;
      }
      while (pl2 < list2.length && list2[pl2] == currEl) {
        c++;
        pl2++;
      }
      similarity += currEl * c;
      countMap.set(currEl, c);
    }
    pl1++;
  }
  return similarity;
}

console.log("Puzzle 2 is: ", puzzle2Again());
