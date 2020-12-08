/*
Specifically, they need you to find the two entries that sum to 2020 and then multiply those two numbers together.

For example, suppose your expense report contained the following:
1721
979
366
299
675
1456

In this list, the two entries that sum to 2020 are 1721 and 299. Multiplying them together produces 1721 * 299 = 514579, so the correct answer is 514579.

Of course, your expense report is much larger. Find the two entries that sum to 2020; what do you get if you multiply them together?
*/

const data = require("fs")
  .readFileSync("input.txt", { encoding: "utf-8" })
  .split("\r\n")
  .map((num) => +num);

// sort ascending
ascdata = data.sort((a, b) => a - b);

let i = 0;
let j = 0;
let answer = 0;

for (i = 0; i < ascdata.length - 1; i++) {
  for (j = i + 1; j < ascdata.length; j++) {
    if (ascdata[i] + ascdata[j] == 2020) {
      answer = ascdata[i] * ascdata[j];
      break;
    }
  }
}

console.log(answer);
