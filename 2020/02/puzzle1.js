/*
For example, suppose you have the following list:

1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
Each line gives the password policy and then the password. The password policy indicates the lowest and highest number of times a given letter must appear for the password to be valid. For example, 1-3 a means that the password must contain a at least 1 time and at most 3 times.

In the above example, 2 passwords are valid. The middle password, cdefg, is not; it contains no instances of b, but needs at least 1. The first and third passwords are valid: they contain one a or nine c, both within the limits of their respective policies.

How many passwords are valid according to their policies?

*/

const data = require("fs")
  .readFileSync("input.txt", { encoding: "utf-8" })
  .split("\r\n");

let validNoPasswds = 0;

data.forEach((entry) => {
  let count = 0;
  let [limits, letter, string] = entry.split(" ");
  const [min, max] = limits.split("-");
  letter = letter.replace(":", "");

  for (let i = 0; i < string.length; i++) {
    if (string.charAt(i) === letter) {
      count++;
    }
  }

  if (count >= min && count <= max) {
    validNoPasswds++;
  }
});

console.log(validNoPasswds);
