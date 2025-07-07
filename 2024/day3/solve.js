import { readFileSync } from "node:fs";

// const input = readFileSync("./sample.txt", "utf8");
// const input = readFileSync("./sample2.txt", "utf8");
const input = readFileSync("./input.txt", "utf8");

const lines = input.split("\n").filter((x) => x.trim().length !== 0);

function puzzle1(lines) {
  let sum = 0;
  lines.forEach((line) => {
    const regex = /mul\((\d+),(\d+)\)/g;
    const array = [...line.matchAll(regex)];
    for (const m of array) {
      const op1 = Number(m[1]);
      const op2 = Number(m[2]);
      const pdt = op1 * op2;
      sum += pdt;
    }
  });

  return sum;
}

console.log("Puzzle 1 is:", puzzle1(lines));

function puzzle2() {
  let sum = 0;
  let enabled = true;
  lines.forEach((line) => {
    const matches = [...line.matchAll(/(do\(\))|(don't\(\))|(mul\((\d+),(\d+)\))/g)];
    for (const m of matches) {
      if (m[0] === "do()") {
        enabled = true;
      } else if (m[0] == "don't()") {
        enabled = false;
      } else {
        if (enabled) {
          sum += +m[4] * +m[5];
        }
      }
    }
  });

  return sum;
}

console.log("Puzzle 2", puzzle2());
