// the number of Calories each Elf is carrying (your puzzle input)
import { readFileSync } from "fs";

// const rounds = readFileSync(__dirname + "/sample.txt", "utf8").split("\n");
const rounds = readFileSync(__dirname + "/input.txt", "utf8").split("\n");
const roundOutcomes = rounds.map((r) => r.split(" "));

// console.log(rounds);

const points = {
  X: 1,
  Y: 2,
  Z: 3,
};

const outcomePoints = {
  lost: 0,
  draw: 3,
  win: 6,
};

function isDraw(m, o) {
  return (m === "X" && o === "A") || (m === "Y" && o === "B") || (m === "Z" && o == "C");
}

function puzzle1() {
  let totalPoints = 0;
  /*
  Then, a winner for that round is selected: Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock. 
  If both players choose the same shape, the round instead ends in a draw.
  */
  roundOutcomes.forEach(([other, mine]) => {
    if (mine == "X" && other == "C") {
      totalPoints += outcomePoints.win + points.X;
    } else if (mine == "Z" && other == "B") {
      totalPoints += outcomePoints.win + points.Z;
    } else if (mine == "Y" && other == "A") {
      totalPoints += outcomePoints.win + points.Y;
    } else if (isDraw(mine, other)) {
      totalPoints += outcomePoints.draw + points[mine];
    } else {
      totalPoints += outcomePoints.lost + points[mine];
    }
  });

  return totalPoints;
}
//  A,X for Rock, B,Y for Paper, and C,Z for Scissors.
function puzzle2() {
  let totalPoints = 0;

  roundOutcomes.forEach(([other, ending]) => {
    if (ending === "Y") {
      // draw
      const choice = other === "A" ? "X" : other === "B" ? "Y" : "Z";
      totalPoints += points[choice] + outcomePoints.draw;
    } else if (ending === "Z") {
      // win
      // Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock.
      let choice;
      if (other === "A") {
        choice = "Y";
      }
      if (other === "B") {
        choice = "Z";
      }
      if (other === "C") {
        choice = "X";
      }

      totalPoints += points[choice] + outcomePoints.win;
    } else {
      // lose
      let choice;

      if (other === "A") {
        choice = "Z";
      }
      if (other === "B") {
        choice = "X";
      }
      if (other === "C") {
        choice = "Y";
      }

      totalPoints += points[choice] + outcomePoints.lost;
    }
  });

  return totalPoints;
}

console.log("Puzzle 1: ", puzzle1());
console.log("Puzzle 2: ", puzzle2());
