import { readFileSync } from "node:fs";

// const input = readFileSync("./sample.txt", "utf8");
const input = readFileSync("./input.txt", "utf8");

const grid = input.split("\n").map((x) => x.split(""));

function getOccurencesOfXMAS(str) {
  if (str.length < 4) return 0;
  let s = 0;
  for (let i = 0; i < str.length; i++) {
    let substr = str.slice(i, i + 4);
    if (substr == "XMAS" || substr == "SAMX") {
      s += 1;
    }
  }
  return s;
}

function puzzle1() {
  let sum = 0;
  // rows
  grid
    .map((row) => row.join(""))
    .forEach((str) => {
      sum += getOccurencesOfXMAS(str);
    });

  // cols
  for (let c = 0; c < grid[0].length; c++) {
    let colStr = "";
    for (let r = 0; r < grid.length; r++) {
      colStr += grid[r][c];
    }
    sum += getOccurencesOfXMAS(colStr);
  }

  // left diagonals
  // move across top row, diagonally
  for (let i = grid.length - 1; i >= 0; i--) {
    let ldiagStr = "";
    let r = 0,
      c = i;
    for (; c < grid[0].length; c++, r++) {
      ldiagStr += grid[r][c];
    }
    // console.log(ldiagStr);

    sum += getOccurencesOfXMAS(ldiagStr);
  }
  // move across the left margin diagonally
  for (let i = 1; i < grid.length; i++) {
    let ldiagStr = "";
    let r = i;
    let c = 0;
    for (; r < grid.length; c++, r++) {
      ldiagStr += grid[r][c];
    }
    // console.log(ldiagStr);

    sum += getOccurencesOfXMAS(ldiagStr);
  }

  // right diag
  // move across top row, diagonally
  for (let i = grid.length - 1; i >= 0; i--) {
    let ldiagStr = "";
    let r = 0,
      c = i;
    for (; c >= 0; c--, r++) {
      ldiagStr += grid[r][c];
    }
    // console.log(ldiagStr);

    sum += getOccurencesOfXMAS(ldiagStr);
  }

  // move across the right margin diagonally
  for (let i = 1; i < grid.length; i++) {
    let ldiagStr = "";
    let r = i;
    let c = grid[i].length - 1;
    for (; r < grid.length; c--, r++) {
      ldiagStr += grid[r][c];
    }
    // console.log(ldiagStr);

    sum += getOccurencesOfXMAS(ldiagStr);
  }

  return sum;
}

console.log("Puzzle 1 is:", puzzle1());

function isXMAS(subgrid) {
  let l = subgrid[0][0] + subgrid[1][1] + subgrid[2][2];
  let r = subgrid[0][2] + subgrid[1][1] + subgrid[2][0];

  return (l == "MAS" || l == "SAM") && (r == "MAS" || r == "SAM") ? 1 : 0;
}

function puzzle2() {
  // Okay X-MAS
  // 3x3 subgrid across whole grid
  let s = 0;

  for (let i = 0; i <= grid.length - 3; i++) {
    const threerows = grid.slice(i, i + 3);
    for (let j = 0; j <= grid[i].length - 3; j++) {
      const subgrid = threerows.map((row) => row.slice(j, j + 3));
      s += isXMAS(subgrid);
    }
  }

  return s;
}

console.log("Puzzle 2", puzzle2());
