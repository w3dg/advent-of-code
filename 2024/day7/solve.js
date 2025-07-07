import { readFileSync } from "node:fs";

// const input = readFileSync("./sample.txt", "utf8");
const input = readFileSync("./input.txt", "utf8");

let lines = input.split("\n");

const equations = lines.map((l) => {
  let [target, nums] = l.split(":");
  target = Number(target);
  nums = nums.trim().split(" ").map(Number);
  const equation = { target, nums };
  return equation;
});

function check(currVal, nums, target) {
  if (nums.length === 0) {
    return currVal === target;
  } else {
    let r1 = currVal + nums[0];
    let res1 = check(r1, nums.slice(1), target);

    let r2 = currVal * nums[0];
    let res2 = check(r2, nums.slice(1), target);

    return res1 || res2;
  }
}

function puzzle1() {
  let sum = 0;
  equations.forEach((eqn) => {
    const { target, nums } = eqn;
    const canBeMadeTrue = check(nums[0], nums.slice(1), target);

    if (canBeMadeTrue) {
      sum += target;
    }
  });

  return sum;
}

console.log("Puzzle 1 is:", puzzle1());

// same code as before but the concat is added
function check2(currVal, nums, target) {
  if (nums.length === 0) {
    return currVal === target;
  } else {
    let r1 = currVal + nums[0];
    let res1 = check2(r1, nums.slice(1), target);

    let r2 = currVal * nums[0];
    let res2 = check2(r2, nums.slice(1), target);

    let r3 = Number("" + currVal + nums[0]);
    let res3 = check2(r3, nums.slice(1), target);

    return res1 || res2 || res3;
  }
}

function puzzle2() {
  let sum = 0;
  equations.forEach((eqn) => {
    const { target, nums } = eqn;
    const canBeMadeTrue = check2(nums[0], nums.slice(1), target);
    if (canBeMadeTrue) {
      sum += target;
    }
  });

  return sum;
}

console.log("Puzzle 2", puzzle2());
