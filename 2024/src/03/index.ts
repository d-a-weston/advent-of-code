import fs from "fs";

const part1 = (input: string) => {
  const pattern = /mul\(\d+,\d+\)/g;
  const operations: string[] = input.match(pattern) || [];

  let total = 0;

  for (const op of operations) {
    const values: number[] = op.match(/\d+/g)?.map(Number) || [];

    const a = values[0];
    const b = values[1];

    total += a * b;
  }

  return total;
};

const part2 = (input: string) => {
  const doPattern = /do\(\)/;
  const dontPattern = /don't\(\)/;
  const mulPattern = /mul\(\d+,\d+\)/g;

  let total = 0;

  const doBlocks = input.split(doPattern);

  for (const doBlock of doBlocks) {
    const operations: string[] =
      doBlock.split(dontPattern)[0].match(mulPattern) || [];

    for (const op of operations) {
      const values: number[] = op.match(/\d+/g)?.map(Number) || [];

      const a = values[0];
      const b = values[1];

      total += a * b;
    }
  }

  return total;
};

const input = fs.readFileSync("./input.txt", "utf8").trim();
console.log(`Part 1: ${part1(input)}`);
console.log(`Part 2: ${part2(input)}`);
