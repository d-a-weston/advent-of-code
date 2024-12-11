import fs from "fs";

const part1 = (input: string) => {
  let stones = input.match(/\d+/g)?.map(Number) || [];

  for (let i = 0; i < 25; i++) {
    const newStones: number[] = [];

    for (let x = 0; x < stones.length; x++) {
      const stone = stones[x];

      if (stone === 0) {
        newStones.push(1);
        continue;
      }

      if (`${stone}`.length % 2 === 0) {
        const stringified = `${stone}`;

        const stoneA = parseInt(
          stringified.slice(0, Math.floor(stringified.length / 2)),
        );
        const stoneB = parseInt(
          stringified.slice(Math.floor(stringified.length / 2)),
        );

        newStones.push(stoneA, stoneB);

        continue;
      }

      newStones.push(stone * 2024);
    }

    stones = newStones;
  }

  return stones.length;
};

const cache: Record<string, number> = {};

const blink = (maxSteps: number, step: number, stone: number): number => {
  const key = `${step}-${stone}`;

  if (cache[key]) return cache[key];

  if (step === maxSteps) {
    return 1;
  }

  let result: number;

  if (stone === 0) {
    result = blink(maxSteps, step + 1, 1);
  } else if (`${stone}`.length % 2 === 0) {
    const stringified = `${stone}`;
    const mid = Math.floor(stringified.length / 2);

    const stoneA = parseInt(stringified.slice(0, mid));
    const stoneB = parseInt(stringified.slice(mid));

    result =
      blink(maxSteps, step + 1, stoneA) + blink(maxSteps, step + 1, stoneB);
  } else {
    result = blink(maxSteps, step + 1, stone * 2024);
  }

  cache[key] = result;
  return result;
};

const part2 = (input: string) => {
  const stones = input.match(/\d+/g)?.map(Number) || [];

  let total = 0;

  for (const stone of stones) {
    total += blink(75, 0, stone);
  }

  return total;
};

const input = fs.readFileSync("./input.txt", "utf8").trim();
console.log(`Part 1: ${part1(input)}`);
console.log(`Part 2: ${part2(input)}`);
