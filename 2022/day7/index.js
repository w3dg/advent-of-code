import { readFileSync } from "fs";

import { terminalOutputParser } from "./terminalOutputParser";

const input = readFileSync(__dirname + "/sample.txt", "utf8");
// const input = readFileSync(__dirname + "/input.txt", "utf8");

// terminalOutputParser(input);

function puzzle1() {}

function puzzle2() {}

console.log("Puzzle 1: ", puzzle1());
console.log("Puzzle 2: ", puzzle2());
