// Due to the aforementioned Planck lengths, the rope must be quite short; in fact, the head (H) and tail (T)
//  must always be touching(diagonally adjacent and even overlapping both count as touching):
const checkIsTouching = (pt1, pt2) => {
  const validX = [pt1.x - 1, pt1.x, pt1.x + 1];
  const validY = [pt1.y - 1, pt1.y, pt1.y + 1];

  return validX.indexOf(pt2.x) !== -1 && validY.indexOf(pt2.y) !== -1;
};

// check pt2 is in st line of pt1
const checkIsInStraightLine = (pt1, pt2) => {
  // if same y level, they are on hosrizontally same line (X)
  if (pt1.y === pt2.y) return "X";
  // if same x level, they are on vertically same line (Y)
  if (pt1.x === pt2.x) return "Y";
  return "N";
};

export function motionsToCartesianInstr(motions) {
  return motions.map((motion) => {
    let { direction, steps } = motion;
    const ci = { axis: undefined, value: undefined };

    if (direction === "R" || direction === "L") {
      ci.axis = "X";
      ci.value = direction === "R" ? steps : -steps;
    } else {
      ci.axis = "Y";
      ci.value = direction === "U" ? steps : -steps;
    }

    return ci;
  });
}

// it is a rope with 2 knots, H, T or in a larger rope, H, 1
export function updatePositionTail(rope) {
  // console.log(rope);
  // if they are touching, is its fine, return

  if (checkIsTouching(rope[rope.length - 1], rope[0])) {
    return rope;
  }

  // Case 1: If the head is ever two steps directly up, down, left, or right from the pt2,
  // the tail must also move one step in that direction so it remains close enough:

  const isStraightAxis = checkIsInStraightLine(rope[0], rope[rope.length - 1]);

  if (isStraightAxis !== "N") {
    // one step in that direction
    if (isStraightAxis === "X") {
      rope[rope.length - 1].x = rope[rope.length - 1].x + (rope[0].x > rope[rope.length - 1].x ? 1 : -1);
    } else {
      rope[rope.length - 1].y = rope[rope.length - 1].y + (rope[0].y > rope[rope.length - 1].y ? 1 : -1);
    }
  }

  // Case 2: Otherwise, if the head and tail aren't touching and aren't in the same row or column,
  // the tail always moves one step diagonally to keep up:
  else {
    // standing at pt1, get the direction of pt2 diagonally in 4 directions by x y signs
    const xdiff = rope[0].x - rope[rope.length - 1].x;
    const ydiff = rope[0].y - rope[rope.length - 1].y;

    // update tail's coords in that direction by 1 unit
    rope[rope.length - 1].x = rope[rope.length - 1].x + (xdiff > 0 ? 1 : -1);
    rope[rope.length - 1].y = rope[rope.length - 1].y + (ydiff > 0 ? 1 : -1);
  }

  return rope;
}
