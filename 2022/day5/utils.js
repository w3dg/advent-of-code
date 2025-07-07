export function makeMoveInstrObjects(moveInstr) {
  return moveInstr.split("\n").map((instruction) => {
    const [number, source, destination] = extractNumbersFromInstruction(instruction);

    return {
      numberOfObjects: Number(number),
      sourceStack: Number(source),
      destinationStack: Number(destination),
    };
  });
}

export function makeStackArray(stackinput) {
  const initial = stackinput.split("\n");
  const positionStr = initial[initial.length - 1];

  const resultArr = [];
  for (let i = 0; i < positionStr.length; i++) {
    // going through the label 1 2 3 4... through the position Str
    if (!Number.isNaN(Number.parseInt(positionStr.charAt(i)))) {
      let arr = [];
      const fixedIndex = i;
      // going through the stacks at the positions of each number labels
      // go through the same index (fixedIndex) in the above rows  to grab the stack from bottom to top
      let j = initial.length - 2;
      while (j >= 0) {
        arr.push(initial[j].charAt(fixedIndex));
        j--;
      }
      arr = arr.filter((element) => element !== " "); // filter out empty elements
      resultArr.push(arr);
    }
  }

  return resultArr;
}

export function extractNumbersFromInstruction(instruction) {
  const result = instruction.match(/move (\d+) from (\d+) to (\d+)/);
  return [result[1], result[2], result[3]];
}
