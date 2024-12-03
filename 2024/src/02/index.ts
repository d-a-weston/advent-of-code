import fs from "fs";

const isSafe = (values: number[]) => {
  const order = values[0] > values[1] ? 0 : 1;

  for (let i = 1; i < values.length; i++) {
    const curr = values[i];
    const prev = values[i - 1];
    const diff = Math.abs(curr - prev);

    if (diff < 1 || diff > 3) return false;

    switch (order) {
      case 0:
        if (curr > prev) return false;
        break;
      case 1:
        if (curr < prev) return false;
        break;
    }
  }

  return true;
};

const part1 = (input: string) => {
  const lines = input.split("\n");

  let safe = 0;

  for (const line of lines) {
    const values: number[] = line.match(/\d+/g)?.map(Number) || [];

    if (isSafe(values)) safe += 1;
  }

  return safe;
};

const part2 = (input: string) => {
  const lines = input.split("\n");

  let safe = 0;

  for (const line of lines) {
    const values: number[] = line.match(/\d+/g)?.map(Number) || [];

    if (isSafe(values)) {
      safe += 1;
      continue;
    }

    const isSafeWithRemoved = values.some((_, i) => {
      const newValues = [...values];
      newValues.splice(i, 1);
      return isSafe(newValues);
    });

    if (isSafeWithRemoved) safe += 1;
  }

  return safe;
};

const input = fs.readFileSync("./input.txt", "utf8").trim();
console.log(`Part 1: ${part1(input)}`);
console.log(`Part 2: ${part2(input)}`);
