import { readFileSync } from "fs";

// const rucksackCollection = readFileSync(__dirname + "/sample.txt", "utf8").split("\n");
const rucksackCollection = readFileSync(__dirname + "/input.txt", "utf8").split("\n");

const strsmall = "abcdefghijklmnopqrstuvwxyz";
const strlarge = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";

const rucksackPartCollection = rucksackCollection.map((str) => {
  const part1 = str
    .split("")
    .slice(0, str.length / 2)
    .join("");

  const part2 = str
    .split("")
    .slice(str.length / 2)
    .join("");
  return [part1, part2];
});

const calPriority = (character) => {
  if (strsmall.indexOf(character) !== -1) {
    return strsmall.indexOf(character) + 1;
  }

  if (strlarge.indexOf(character) !== -1) {
    return strlarge.indexOf(character) + 27;
  }
};

function findCommonLetter(rucksack) {
  let [part1, part2] = rucksack;

  part1 = part1.split("");
  part2 = part2.split("");

  for (let i = 0; i < part1.length; i++) {
    if (part2.indexOf(part1[i]) !== -1) {
      return part1[i];
    }
  }
}

function findCommonLetterInThreeString(allStrings) {
  let [part1, part2, part3] = allStrings;
  part1 = part1.split("");
  part2 = part2.split("");
  part3 = part3.split("");

  let common = "";

  strsmall.split("").forEach((letter) => {
    if (part1.indexOf(letter) !== -1 && part2.indexOf(letter) !== -1 && part3.indexOf(letter) !== -1) {
      common = letter;
    }
  });

  strlarge.split("").forEach((letter) => {
    if (part1.indexOf(letter) !== -1 && part2.indexOf(letter) !== -1 && part3.indexOf(letter) !== -1) {
      common = letter;
    }
  });

  return common;
}

function puzzle1() {
  let totalScore = 0;
  rucksackPartCollection.forEach((rucksack) => {
    const commonLetter = findCommonLetter(rucksack);
    totalScore += calPriority(commonLetter);
  });
  return totalScore;
}

function puzzle2() {
  let totalScore = 0;
  for (let i = 0; i < rucksackCollection.length; i += 3) {
    const collectionOfThree = [rucksackCollection[i], rucksackCollection[i + 1], rucksackCollection[i + 2]];
    const commonLetter = findCommonLetterInThreeString(collectionOfThree);
    totalScore += calPriority(commonLetter);
  }
  return totalScore;
}

console.log("Puzzle 1: ", puzzle1());
console.log("Puzzle 2: ", puzzle2());
