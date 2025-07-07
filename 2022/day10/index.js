import { readFileSync } from "fs";

import {} from "./utils";

// const input = readFileSync(__dirname + "/sample-big.txt", "utf8").split("\n");
const input = readFileSync(__dirname + "/input.txt", "utf8").split("\n");

const cpuInstructions = input.map((inst) => {
  if (inst === "noop") {
    return {
      noOfCyclesToComplete: 1,
      instruction: inst,
    };
  } else {
    // its addx
    const [instruction, value] = inst.split(" ");
    return {
      noOfCyclesToComplete: 2,
      instruction,
      value: Number(value),
    };
  }
});

function puzzle1() {
  let cycleCount = 1; // ( starting at 1, only updates after the thing is done, what tripped me up )
  let registerX = 1;
  let cycleToBeChecked = [20, 60, 100, 140, 180, 220];
  let signalStrengths = [];

  cpuInstructions.forEach((instruction) => {
    for (let i = 0; i < instruction.noOfCyclesToComplete; i++) {
      if (cycleToBeChecked.indexOf(cycleCount) != -1) {
        // the cycle number multiplied by the value of the X register
        // console.log("Cycle " + cycleCount, registerX);
        const strength = cycleCount * registerX;
        signalStrengths.push(strength);
      }
      cycleCount++;
    }

    if (instruction.noOfCyclesToComplete === 2) {
      // update if it is a addx
      registerX += instruction.value;
    }
    // console.log(instruction, "done", registerX);
  });

  // console.log(signalStrengths);
  const sumOfStrengths = signalStrengths.reduce((prev, curr) => prev + curr, 0);
  // console.log(sumOfStrengths);
  return sumOfStrengths;
}

function puzzle2() {
  // Keep track of the "CRT Screen" as a 2d array
}

console.log("Puzzle 1: ", puzzle1());
console.log("Puzzle 2: ", puzzle2());
