import fs from "fs";

const part1 = (input: string) => {
  const lines = input.split("\n");

  const left: number[] = [];
  const right: number[] = [];

  for (const line of lines) {
    const [leftValue, rightValue] = line.match(/\d+/g)?.map(Number) || [];

    left.push(leftValue);
    right.push(rightValue);
  }

  left.sort();
  right.sort();

  let diffs = 0;

  for (let i = 0; i < left.length; i++) {
    diffs += Math.abs(left[i] - right[i]);
  }

  return diffs;
};

const part2 = (input: string) => {
  const lines = input.split("\n");

  const left: number[] = [];
  const right: Record<number, number> = {};

  for (const line of lines) {
    const [leftValue, rightValue] = line.match(/\d+/g)?.map(Number) || [];

    left.push(leftValue);

    if (right[rightValue] >= 1) right[rightValue] += 1;
    else right[rightValue] = 1;
  }

  let similarity = 0;

  for (const id of left) {
    if (!right[id]) continue;

    const value = id * right[id];
    similarity += value;
  }

  return similarity;
};

const input = fs.readFileSync("./input.txt", "utf8").trim();
console.log(`Part 1: ${part1(input)}`);
console.log(`Part 2: ${part2(input)}`);
