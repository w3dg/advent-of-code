export function makeArray(assignment) {
  // '3-6' -> [3,4,5,6]

  let [start, end] = assignment.split("-");
  start = Number(start);
  end = Number(end);

  let arr = [];

  for (let i = start; i <= end; i++) {
    arr.push(i);
  }

  return arr;
}

export function fullyContains(arr1, arr2) {
  // whether arr1 is contained in arr2 or vice versa

  // grab the shorter length array or equal length and check
  // if every element of the shorter length array is
  // contained in the larger one

  let largerArr = arr1.length > arr2.length ? arr1 : arr2;
  let smallerArr = arr1.length < arr2.length ? arr1 : arr2;

  if (arr1.length === arr2.length) {
    largerArr = arr1;
    smallerArr = arr2;
  }

  for (let i = 0; i < smallerArr.length; i++) {
    const element = smallerArr[i];

    if (largerArr.indexOf(element) === -1) {
      return false;
    }
  }

  return true;
}

export function doesOverlap(arr1, arr2) {
  // whether arr1 and arr2 overlap at any point or not
  let largerArr = arr1.length > arr2.length ? arr1 : arr2;
  let smallerArr = arr1.length < arr2.length ? arr1 : arr2;

  if (arr1.length === arr2.length) {
    largerArr = arr1;
    smallerArr = arr2;
  }

  for (const elt of smallerArr) {
    if (largerArr.indexOf(elt) !== -1) {
      return true;
    }
  }

  return false;
}
