import { readFileSync } from "fs";

import { makeArray, fullyContains, doesOverlap } from "./utils";

// const sectionAssignments = readFileSync(__dirname + "/sample.txt", "utf8").split("\n");
const sectionAssignments = readFileSync(__dirname + "/input.txt", "utf8").split("\n");

const sectionAssignmentPairs = sectionAssignments.map((pair) => pair.split(","));

function puzzle1() {
  let fullyContainsCount = 0;
  sectionAssignmentPairs.forEach((pair) => {
    const [elf1, elf2] = pair;
    const elf1Elements = makeArray(elf1);
    const elf2Elements = makeArray(elf2);

    // if one of the totals fully contain the other
    if (fullyContains(elf1Elements, elf2Elements)) {
      fullyContainsCount += 1;
    }
  });

  return fullyContainsCount;
}

function puzzle2() {
  let overlapArrCount = 0;

  sectionAssignmentPairs.forEach((pair) => {
    const [elf1, elf2] = pair;
    const elf1Elements = makeArray(elf1);
    const elf2Elements = makeArray(elf2);

    // if overlap increment counter
    if (doesOverlap(elf1Elements, elf2Elements)) {
      overlapArrCount += 1;
    }
  });

  return overlapArrCount;
}

console.log("Puzzle 1: ", puzzle1());
console.log("Puzzle 2: ", puzzle2());
