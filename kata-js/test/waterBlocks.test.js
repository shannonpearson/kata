const assert = require('assert');
const {
  calcBlocksBetween,
  findLocalMax,
  findRight,
  waterBlocks,
} = require('../src/waterBlocks');

describe('findBlocksBetween', () => {
  // make this a loop with different inputs and expects
  it('should return expected result', () => {
    // Arrange
    const input = [5, 4, 0, 3, 2, 6];

    // Act
    const result = calcBlocksBetween(input);

    // Assert
    assert.equal(result, 11);
  });
  it('should return expected result', () => {
    // Arrange
    const input = [6, 4, 0, 3, 2, 5];

    // Act
    const result = calcBlocksBetween(input);

    // Assert
    assert.equal(result, 11);
  });
  it('should return 0 if input is invalid', () => {
    // Arrange
    const input = [3, 2, 4, 0, 2];

    // Act
    const result = calcBlocksBetween(input);

    // Assert
    assert.equal(result, 0);
  });
  it('should return 0 if array is ordered descending', () => {
    // Arrange
    const input = [5, 4, 3, 1];

    // Act
    const result = calcBlocksBetween(input);

    // Assert
    assert.equal(result, 0);
  });
  it('should return 0 if array is ordered ascending', () => {
    // Arrange
    const input = [1, 3, 5, 6];

    // Act
    const result = calcBlocksBetween(input);

    // Assert
    assert.equal(result, 0);
  });
});

describe('findLocalMax', () => {
  it('should find local max at beginning', () => {
    // Arrange
    const input = [5, 4, 2, 1];
    // Act
    const result = findLocalMax(input);
    // Assert
    assert.equal(result, 0);
  });
  it('should find local max at end', () => {
    // Arrange
    const input = [1, 3, 4, 6];

    // Act
    const result = findLocalMax(input);

    // Assert
    assert.equal(result, 3);
  });
  it('should find first local max', () => {
    // Arrange
    const input = [1, 3, 5, 2, 6, 7, 0];

    // Act
    const result = findLocalMax(input);

    // Assert
    assert.equal(result, 2);
  });
});

describe('findRight', () => {
  it('should work for normal input', () => {
    // Arrange
    const input = [4, 0, 3, 1, 7, 5, 2, 8];

    // Act
    const result = findRight(input);

    // Assert
    assert.equal(result, 4);
  });
  it('should work for descending array', () => {
    // Arrange
    const input = [4, 3, 2, 1];

    // Act
    const result = findRight(input);

    // Assert
    assert.equal(result, 1);
  });
  it('should work for ascending array', () => {
    // Arrange
    const input = [4, 5, 6];

    // Act
    const result = findRight(input);

    // Assert
    assert.equal(result, 1);
  });
  it('should work for equal array', () => {
    // Arrange
    const input = [4, 4, 4, 4, 4];

    // Act
    const result = findRight(input);

    // Assert
    assert.equal(result, 1);
  });
  it('should return 0 for array length 1', () => {
    // Arrange
    const input = [4];

    // Act
    const result = findRight(input);

    // Assert
    assert.equal(result, 0);
  });
});

describe('waterBlocks', () => {
  it('should work', () => {
    // Arrange
    const input = [1, 2, 1, 4, 1, 3, 4, 3, 0, 2, 1, 3, 0, 2];

    // Act
    const result = waterBlocks(input);

    // Assert
    assert.equal(result, 13);
  });
  xit('should work', () => {
    // Arrange
    const input = [];

    // Act
    const result = waterBlocks(input);

    // Assert
    assert.equal(result, 0);
  });
  xit('should work', () => {
    // Arrange
    const input = [];

    // Act
    const result = waterBlocks(input);

    // Assert
    assert.equal(result, 0);
  });
  xit('should work', () => {
    // Arrange
    const input = [];

    // Act
    const result = waterBlocks(input);

    // Assert
    assert.equal(result, 0);
  });
});
