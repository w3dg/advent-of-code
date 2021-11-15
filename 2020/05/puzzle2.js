/*
--- Part Two ---
Ding! The "fasten seat belt" signs have turned on. Time to find your seat.

It's a completely full flight, so your seat should be the only missing boarding pass in your list. However, there's a catch: some of the seats at the very front and back of the plane don't exist on this aircraft, so they'll be missing from your list as well.

Your seat wasn't at the very front or back, though; the seats with IDs +1 and -1 from yours will be in your list.

What is the ID of your seat?

*/

const fs = require("fs");
const input = fs.readFileSync("2020/05/input.txt", "utf8").split("\r\n");
// const input = fs.readFileSync("2020/05/sample.txt", "utf8").split("\r\n");

function binaryBoarding(boardingPass) {
  let rowChecks = boardingPass.slice(0, 7);
  let colChecks = boardingPass.slice(7, 10);
  let lowRow = 0;
  let highRow = 127;

  let lowCol = 0;
  let highCol = 7;

  let row = 0;
  let col = 0;
  rowChecks.split("").forEach((check) => {
    let mid = Math.floor((lowRow + highRow) / 2);
    if (check === "F") {
      highRow = mid;
      row = highRow;
    } else {
      lowRow = mid + 1;
      row = lowRow;
    }
  });

  colChecks.split("").forEach((check) => {
    let mid = Math.floor((lowCol + highCol) / 2);
    if (check === "L") {
      highCol = mid;
      col = highCol;
    } else {
      lowCol = mid + 1;
      col = lowCol;
    }
  });

  return [row, col];
}

let seatIDs = [];

input.forEach((b) => {
  const [row, col] = binaryBoarding(b);
  seatID = row * 8 + col;
  seatIDs.push(seatID);
});

let highest = 922;
let lowest = 0;

for (let i = lowest; i <= highest; i++) {
  if (
    seatIDs.includes(i + 1) &&
    seatIDs.includes(i - 1) &&
    !seatIDs.includes(i) // empty seat , seats before and after are present. hence check i as well. i missed this check.
  ) {
    console.log("Your Seat ID is ", i);
  }
}
