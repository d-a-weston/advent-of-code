import fs from "fs";

const part1 = (input: string) => {
  const lines = input.split("\n");

  const haveSeen: Record<number, number[]> = {};
  const updates: number[][] = [];

  for (const line of lines) {
    if (line === "") continue;

    if (line.includes("|")) {
      const [x, y] = line.split("|").map(Number);

      if (!haveSeen[y]) {
        haveSeen[y] = [x];
      } else {
        haveSeen[y].push(x);
      }
    } else {
      updates.push(line.split(",").map(Number));
    }
  }

  let total = 0;

  for (const update of updates) {
    let passes = true;

    for (const [i, val] of update.entries()) {
      const partial = update.slice(i + 1);

      if (partial.some((e) => haveSeen[val]?.includes(e))) {
        passes = false;
        break;
      }
    }

    if (passes) total += update[Math.floor(update.length / 2)];
  }

  return total;
};

const part2 = (input: string) => {
  const lines = input.split("\n");

  const haveSeen: Record<number, number[]> = {};
  const updates: number[][] = [];

  for (const line of lines) {
    if (line === "") continue;

    if (line.includes("|")) {
      const [x, y] = line.split("|").map(Number);

      if (!haveSeen[y]) {
        haveSeen[y] = [x];
      } else {
        haveSeen[y].push(x);
      }
    } else {
      updates.push(line.split(",").map(Number));
    }
  }

  const faultyUpdates: number[][] = [];

  for (const update of updates) {
    let passes = true;

    for (const [i, val] of update.entries()) {
      const partial = update.slice(i + 1);

      if (partial.some((e) => haveSeen[val]?.includes(e))) {
        passes = false;
        break;
      }
    }

    if (!passes) faultyUpdates.push(update);
  }

  let total = 0;

  for (const faultyUpdate of faultyUpdates) {
    const fixedUpdate: number[] = [];

    for (const val of faultyUpdate) {
      for (let i = fixedUpdate.length; i >= 0; i--) {
        if (i === 0 || haveSeen[val]?.includes(fixedUpdate[i - 1])) {
          fixedUpdate.splice(i, 0, val);
          break;
        }
      }
    }

    total += fixedUpdate[Math.floor(fixedUpdate.length / 2)];
  }

  return total;
};

const input = fs.readFileSync("./input.txt", "utf8").trim();
console.log(`Part 1: ${part1(input)}`);
console.log(`Part 2: ${part2(input)}`);
