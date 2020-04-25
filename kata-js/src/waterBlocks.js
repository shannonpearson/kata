function calcBlocksBetween(arr) {
  let sum = 0;
  let minMax = Math.min(arr[0], arr[arr.length - 1]);
  for (let i = 1; i < arr.length - 1; i++) {
    if (arr[i] > minMax) return 0;
    sum += minMax - arr[i];
  }
  return sum;
}

function findLocalMax(arr) {
  for (let i = 0; i < arr.length - 1; i++) {
    if (arr[i] > arr[i + 1]) return i;
  }
  return arr.length - 1;
}

function findRight(arr) {
  if (arr.length === 1) {
    return 0;
  }
  for (let i = 1; i < arr.length; i++) {
    if (arr[i] >= arr[0]) return i;
  }
  return arr.indexOf(Math.max(...arr.slice(1)));
}

function waterBlocks(arr) {
  if (arr.length === 1) {
    return 0;
  }

  let blocks = 0;
  const localMax = findLocalMax(arr);
  let right = localMax + findRight(arr.slice(localMax));

  if (!right) blocks += waterBlocks(arr.slice(1));
  else {
    blocks += calcBlocksBetween(arr.slice(localMax, right + 1));
    if (right < arr.length - 1) {
      blocks += waterBlocks(arr.slice(right));
    }
  }

  return blocks;
}

module.exports = { calcBlocksBetween, findLocalMax, findRight, waterBlocks };
