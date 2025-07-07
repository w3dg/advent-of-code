/**
 * @param arr2D the 2d array to search in
 * @param startI the index of the ith row of the element from where to search
 * @param startJ the index of the jth col of the element from where to search
 * @return boolean true | false
 */
export function checkVisibility(arr2D, startI, startJ) {
  let topVisible = true;
  let leftVisible = true;
  let bottomVisible = true;
  let rightVisible = true;

  const startingPos = arr2D[startI][startJ];

  // if at edges, already visible
  if (startI === 0 || startJ == 0 || startI === arr2D.length - 1 || startJ === arr2D[0].length - 1) {
    return true;
  }

  // if inside, only visible on top bottom left right if all trees till edge
  //  are lower than the given trees height

  // bottom
  for (let i = startI + 1; i < arr2D.length; i++) {
    if (arr2D[i][startJ] >= startingPos) {
      bottomVisible = false;
    }
  }

  for (let i = startI - 1; i >= 0; i--) {
    if (arr2D[i][startJ] >= startingPos) {
      topVisible = false;
    }
  }

  for (let j = startJ + 1; j < arr2D[0].length; j++) {
    if (arr2D[startI][j] >= startingPos) {
      rightVisible = false;
    }
  }

  for (let j = startJ - 1; j >= 0; j--) {
    if (arr2D[startI][j] >= startingPos) {
      leftVisible = false;
    }
  }

  return topVisible || leftVisible || bottomVisible || rightVisible;
}

/**
 * @param arr2D the 2d array to search in
 * @param startI the index of the ith row of the element from where to search
 * @param startJ the index of the jth col of the element from where to search
 * @return scenic score of the given element
 */

export function calculateScenicScore(arr2D, startI, startJ) {
  let scenicScoresInDirections = [];
  const startingPos = arr2D[startI][startJ];

  if (startI === 0 || startJ == 0 || startI === arr2D.length - 1 || startJ === arr2D[0].length - 1) {
    return -1;
  }

  // bottom
  for (let i = startI + 1; i < arr2D.length; i++) {
    if (arr2D[i][startJ] >= startingPos || i == arr2D.length - 1) {
      scenicScoresInDirections.push(i - startI);
      break;
    }
  }

  for (let i = startI - 1; i >= 0; i--) {
    if (arr2D[i][startJ] >= startingPos || i == 0) {
      scenicScoresInDirections.push(startI - i);
      break;
    }
  }

  for (let j = startJ + 1; j < arr2D[0].length; j++) {
    if (arr2D[startI][j] >= startingPos || j == arr2D[0].length - 1) {
      scenicScoresInDirections.push(j - startJ);
      break;
    }
  }

  for (let j = startJ - 1; j >= 0; j--) {
    if (arr2D[startI][j] >= startingPos || j == 0) {
      scenicScoresInDirections.push(startJ - j);
      break;
    }
  }

  const score = scenicScoresInDirections.reduce((p, c) => p * c, 1);

  return score;
}
