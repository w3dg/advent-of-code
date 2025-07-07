import { readFileSync } from "node:fs";

// const input = readFileSync("./sample.txt", "utf8");
const input = readFileSync("./input.txt", "utf8");

const lines = input.split("\n");

let diskMap = lines[0].split("").map((x) => Number(x));

let disk = [];
let diskIndex = {};

let isFile = true;
let diskCounter = 0;
for (let i = 0; i < diskMap.length; i++, isFile = !isFile) {
  let times = diskMap[i];

  if (isFile) {
    let start = disk.length;
    let end = disk.length + times - 1;
    diskIndex[diskCounter] = { start, end };

    for (let j = 0; j < times; j++) {
      disk.push(diskCounter);
    }

    diskCounter++;
  } else {
    for (let j = 0; j < times; j++) {
      disk.push(null);
    }
  }
}

diskCounter--;

function puzzle1() {
  let disk_local = [...disk];

  let emptyPtr = 0;

  while (disk_local[emptyPtr] == null) {
    emptyPtr++;
  }

  let lastBlock = disk_local.length - 1;

  while (disk_local[lastBlock] == null) {
    lastBlock--;
  }

  // console.log(emptyPtr, lastBlock);

  // console.log("Before", disk);
  // while the two pointers do not cross each other,
  while (emptyPtr < lastBlock) {
    // find empty on the left,
    while (emptyPtr < lastBlock && disk_local[emptyPtr] != null) {
      emptyPtr++;
    }
    // find last block on the right
    while (emptyPtr < lastBlock && disk_local[lastBlock] == null) {
      lastBlock--;
    }

    // fill the empty on left with the block number and set the right one to null
    disk_local[emptyPtr] = disk_local[lastBlock];
    disk_local[lastBlock] = null;
    // update pointers to be pointing to next valid as well
    emptyPtr++;
    lastBlock--;
  }
  // sort of like a partition subroutine on the quicksort
  return calculateChecksum(disk_local);
  // console.log("After", disk);
}

console.log("Puzzle 1 is:", puzzle1());

// TODO: Optimise part 2 with a better solution, this takes ~7s
function puzzle2() {
  for (let i = diskCounter; i >= 0; i--) {
    const { start, end } = diskIndex[i];
    const sizeOfBlock = end - start + 1;
    // console.log(start, end, "size", sizeOfBlock);

    // find the empty space from the left to put the thing of the desired length
    let startPtr = 0;
    let endPtr = startPtr + sizeOfBlock - 1;

    let found = false;
    while (endPtr < disk.length && endPtr < start) {
      const currWindow = [];
      for (let i = startPtr; i <= endPtr; i++) {
        currWindow.push(disk[i]);
      }

      // find space from left until the thing start to put the thing, otherwise bail out
      if (currWindow.every((x) => x == null)) {
        found = true;
        break;
      } else {
        startPtr++;
        endPtr = startPtr + sizeOfBlock - 1;
      }
    }

    if (found) {
      // console.log(`Found space for ${i} at starting index`, startPtr);

      // update the disk value
      for (let p = startPtr; p <= endPtr; p++) {
        disk[p] = i;
      }
      for (let p = start; p <= end; p++) {
        disk[p] = null;
      }
    }
  }

  return calculateChecksum(disk);
}

console.log("Puzzle 2", puzzle2());

function calculateChecksum(disk) {
  let checksum = 0;
  for (let i = 0; i < disk.length; i++) {
    if (disk[i] != null) {
      const fileId = disk[i];
      checksum += fileId * i;
    }
  }

  return checksum;
}

function displayDisk() {
  console.log(
    disk
      .map((x) => {
        if (x === null) {
          return ".";
        } else return x;
      })
      .join("")
  );
}
