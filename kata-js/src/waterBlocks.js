/* linear with higher coefficient? */

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

/* linear 3n? */

const trap = (height) => {
  const n = height.length;

  const left = [];
  const right = [];

  let leftMax = 0;
  let rightMax = 0;

  for (let i = 0, j = n - 1; i < n, j >= 0; i++, j--) {
    // Scan from left
    left[i] = leftMax;
    leftMax = Math.max(leftMax, height[i]);

    // Scan from right
    right[j] = rightMax;
    rightMax = Math.max(rightMax, height[j]);
  }

  let total = 0;
  for (let i = 0; i < n; i++) {
    let water = Math.min(left[i], right[i]) - height[i];
    total += water > 0 ? water : 0;
  }

  return total;
};

const benchmark = (f) => {
  const start = Date.now();
  for (let i = 0; i < 100000; i++) f();
  console.log(`${Date.now() - start}ms`);
};

const input = [1, 2, 1, 4, 1, 3, 4, 3, 0, 2, 1, 3, 0, 2];
benchmark(() => waterBlocks(input));
benchmark(() => trap(input));

module.exports = { calcBlocksBetween, findLocalMax, findRight, waterBlocks };
