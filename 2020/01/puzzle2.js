/*
--- Part Two ---
The Elves in accounting are thankful for your help; one of them even offers you a starfish coin they had left over from a past vacation. They offer you a second one if you can find three numbers in your expense report that meet the same criteria.

Using the above example again, the three entries that sum to 2020 are 979, 366, and 675. Multiplying them together produces the answer, 241861950.

In your expense report, what is the product of the three entries that sum to 2020?
*/

const data = require("fs")
  .readFileSync("input.txt", { encoding: "utf-8" })
  .split("\r\n")
  .map((num) => +num);

// sort ascending
ascdata = data.sort((a, b) => a - b);

let i = 0;
let j = 0;
let k = 0;
let answer;

for (i = 0; i < ascdata.length - 2; i++) {
  for (j = i + 1; j < ascdata.length - 1; j++) {
    for (k = j + 1; k < ascdata.length; k++) {
      if (ascdata[i] + ascdata[j] + ascdata[k] == 2020) {
        answer = ascdata[i] * ascdata[j] * ascdata[k];
        break;
      }
    }
  }
}
console.log(answer);
