import { readFileSync } from "fs";

import {} from "./utils";

// const datastream = readFileSync(__dirname + "/sample.txt", "utf8");
const datastream = readFileSync(__dirname + "/input.txt", "utf8");

function puzzle1() {
  for (let i = 3; i < datastream.length - 1; i++) {
    const last4 = [datastream.charAt(i), datastream.charAt(i - 1), datastream.charAt(i - 2), datastream.charAt(i - 3)];
    const startMarker = new Set(last4);
    // here iteration is one less than the character count (the answer) due to indices starting at 0 and character position being one extra.
    // hence we return i + 1
    if (startMarker.size === 4) {
      return i + 1;
    }
  }
}

function puzzle2() {
  for (let i = 13; i < datastream.length - 1; i++) {
    const last14 = datastream.slice(i - 13, i + 1); // last 14 characters from the current i position starting at an offset of 14 characters
    const startMarker = new Set(last14);
    if (startMarker.size === 14) {
      return i + 1;
    }
  }
}

console.log("Puzzle 1: ", puzzle1());
console.log("Puzzle 2: ", puzzle2());
