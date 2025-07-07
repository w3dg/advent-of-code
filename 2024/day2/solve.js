import { readFileSync } from "node:fs";

// const input = readFileSync("./sample.txt", "utf8");
const input = readFileSync("./input.txt", "utf8");

let reports = input.split("\n");
// each line is a report
// each report has levels separated by spaces
reports = reports.map((report) => report.split(" ").map(Number));

// either all increasing or all descreasing
// the diff between 2 successive levels is between 1 and 3 (both inclusive)
function isValidReport(report) {
  if (report[0] === report[1]) return false; // no incr or decr
  let initDiff = Math.abs(report[0] - report[1]);
  if (initDiff < 1 || initDiff > 3) return false;

  let isIncreasing = report[0] < report[1];
  for (let i = 1; i + 1 < report.length; i++) {
    if (isIncreasing && report[i] > report[i + 1]) return false;
    if (!isIncreasing && report[i] < report[i + 1]) return false;

    let difference = Math.abs(report[i] - report[i + 1]);
    if (difference < 1 || difference > 3) return false;
  }

  return true;
}

let safeReports = [];
let unsafeReports = [];

reports.forEach((report) => {
  if (isValidReport(report)) {
    safeReports.push(report);
  } else {
    unsafeReports.push(report);
  }
});

function puzzle1() {
  return safeReports.length;
}

console.log("Puzzle 1 is:", puzzle1());

function puzzle2() {
  const modifiedSafeReports = [];
  for (const report of unsafeReports) {
    for (let i = 0; i < report.length; i++) {
      const modified = report.toSpliced(i, 1);
      if (isValidReport(modified)) {
        modifiedSafeReports.push(report);
        break;
      }
    }
  }

  return safeReports.length + modifiedSafeReports.length;
}

console.log("Puzzle 2", puzzle2());
