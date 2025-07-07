import { readFileSync } from "fs";

import { motionsToCartesianInstr, updatePositionTail } from "./utils";

// const input = readFileSync(__dirname + "/sample.txt", "utf8").split("\n");
const input = readFileSync(__dirname + "/input.txt", "utf8").split("\n");

const motions = input.map((str) => {
  const [direction, steps] = str.split(" ");

  return {
    direction,
    steps: Number(steps),
  };
});

const cartesianMotions = motionsToCartesianInstr(motions);

function solve(LENGTH_OF_ROPE) {
  const rope = Array.from({ length: LENGTH_OF_ROPE }, () => {
    const obj = {
      x: 0,
      y: 0,
    };
    return obj;
  });

  // not needing an 2d array to keep track instead i work off of the coordinates, modelling a grid with x,y
  // need to track <unique> positions where the tail has been including the starting point

  const tailPositions = new Set(); //objects references will not be unique, a set with objects will  hence not work, hence I format it in a way, "x_y"
  tailPositions.add(rope[rope.length - 1].x + "_" + rope[rope.length - 1].y);

  cartesianMotions.forEach((motion) => {
    const { axis, value } = motion;
    const lb = Math.min(0, value);
    const ub = Math.max(0, value);

    if (axis === "X") {
      for (let i = lb; i < ub; i++) {
        for (let j = 0; j < LENGTH_OF_ROPE - 1; j++) {
          // everytime make a sliding window

          const slidingWindow = [rope[j], rope[j + 1]];
          if (j == 0) {
            // first time, update x here,
            rope[j].x += value > 0 ? 1 : -1;
          }
          const [_, updatedPt2] = updatePositionTail(slidingWindow);
          rope[j + 1] = updatedPt2;
        }

        // in the end, update the tail position of the rope
        tailPositions.add(rope[rope.length - 1].x + "_" + rope[rope.length - 1].y);
      }
    } else {
      for (let i = lb; i < ub; i++) {
        // everytime make a sliding window

        for (let j = 0; j < LENGTH_OF_ROPE - 1; j++) {
          const slidingWindow = [rope[j], rope[j + 1]];
          if (j == 0) {
            // first time, update y here,
            rope[j].y += value > 0 ? 1 : -1;
          }
          const [_, updatedPt2] = updatePositionTail(slidingWindow);
          rope[j + 1] = updatedPt2;
        }

        // in the end, update the tail position of the rope
        tailPositions.add(rope[rope.length - 1].x + "_" + rope[rope.length - 1].y);
      }
    }
  });

  // console.log(tailPositions);
  return tailPositions.size;
}

function puzzle1() {
  return solve(2);
}
function puzzle2() {
  return solve(10);
}

console.log("Puzzle 1: ", puzzle1());
console.log("Puzzle 2: ", puzzle2());
