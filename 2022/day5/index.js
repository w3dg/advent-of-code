import { readFileSync } from "fs";

import { makeMoveInstrObjects, makeStackArray } from "./utils";

// const [stackinput, moveInstr] = readFileSync(__dirname + "/sample.txt", "utf8").split("\n\n");
const [stackinput, moveInstr] = readFileSync(__dirname + "/input.txt", "utf8").split("\n\n");

const stacks1 = [[], ...makeStackArray(stackinput)]; // add empty array at front to match up array indices with instructions
const stacks2 = [[], ...makeStackArray(stackinput)]; // add empty array at front to match up array indices with instructions

const moveInstrObjArray = makeMoveInstrObjects(moveInstr);

function puzzle1() {
  moveInstrObjArray.forEach((moveInstr) => {
    const { sourceStack, destinationStack } = moveInstr;

    for (let i = 0; i < moveInstr.numberOfObjects; i++) {
      const src = stacks1[sourceStack];
      const element = src[src.length - 1];
      src.splice(src.length - 1);

      stacks1[destinationStack].push(element);
    }
  });

  let str = "";
  stacks1.forEach((x) => {
    if (x.length !== 0) {
      str += x[x.length - 1];
    }
  });

  return str;
}

function puzzle2() {
  moveInstrObjArray.forEach((moveInstr) => {
    let { sourceStack, destinationStack, numberOfObjects } = moveInstr;

    // retain the order and move the items, thats the only difference from part 1

    const chunk = stacks2[sourceStack].slice(-numberOfObjects);
    stacks2[sourceStack] = stacks2[sourceStack].slice(0, stacks2[sourceStack].length - numberOfObjects);

    for (const elt of chunk) {
      stacks2[destinationStack].push(elt);
    }
  });

  let str = "";
  stacks2.forEach((x) => {
    if (x.length !== 0) {
      str += x[x.length - 1];
    }
  });

  return str;
}

console.log("Puzzle 1: ", puzzle1());
console.log("Puzzle 2: ", puzzle2());
