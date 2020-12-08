/*
Each policy actually describes two positions in the password, where 1 means the first character, 2 means the second character, and so on. (Be careful; Toboggan Corporate Policies have no concept of "index zero"!) Exactly one of these positions must contain the given letter. Other occurrences of the letter are irrelevant for the purposes of policy enforcement.

Given the same example list from above:

1-3 a: abcde is valid: position 1 contains a and position 3 does not.
1-3 b: cdefg is invalid: neither position 1 nor position 3 contains b.
2-9 c: ccccccccc is invalid: both position 2 and position 9 contain c.
How many passwords are valid according to the new interpretation of the policies?
*/

const data = require("fs")
  .readFileSync("input.txt", { encoding: "utf-8" })
  .split("\r\n");

let validNoPasswds = 0;

data.forEach((entry) => {
  let count = 0;
  let [positions, letter, string] = entry.split(" ");
  const [p1, p2] = positions.split("-");
  letter = letter.replace(":", "");

  if (string.charAt(p1 - 1) === letter) count++;
  if (string.charAt(p2 - 1) === letter) count++;

  if (count === 1) validNoPasswds++;
});

console.log(validNoPasswds);
