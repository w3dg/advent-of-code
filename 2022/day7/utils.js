export function removeEmptyElement(arr) {
  const res = arr.filter((x) => x != "");
  return res;
}
