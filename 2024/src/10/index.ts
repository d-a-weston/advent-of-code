import fs from "fs";

const dirs = [
  [-1, 0],
  [0, 1],
  [1, 0],
  [0, -1],
];

const isInBounds = (width: number, height: number, x: number, y: number) => {
  return x >= 0 && x <= width && y >= 0 && y <= height;
};

const traverse = (
  nextLoc: number,
  pos: Record<string, number>,
  grid: number[][],
  width: number,
  height: number,
): Record<string, number>[] => {
  const { x, y } = pos;

  if (!isInBounds(width, height, x, y)) return [];

  const cell = grid[y][x];

  if (cell !== nextLoc) {
    return [];
  }

  if (cell === 9) return [{ x, y }];

  return [
    ...dirs
      .map((dir) =>
        traverse(
          cell + 1,
          { x: x + dir[0], y: y + dir[1] },
          grid,
          width,
          height,
        ),
      )
      .flat(),
  ];
};

const part1 = (input: string) => {
  const grid: number[][] = input
    .split("\n")
    .map((a) => a.split("").map(Number));

  const trailheads = [];

  for (let y = 0; y < grid.length; y++) {
    for (let x = 0; x < grid[y].length; x++) {
      if (grid[y][x] === 0) {
        trailheads.push({ x, y });
      }
    }
  }

  let total = 0;

  for (const trailhead of trailheads) {
    const seenPeaks: Record<string, boolean> = {};

    const trail = traverse(
      0,
      trailhead,
      grid,
      grid[0].length - 1,
      grid.length - 1,
    );

    trail.map(
      (peak: Record<string, number>) =>
        (seenPeaks[`${peak.x},${peak.y}`] = true),
    );

    total += Object.values(seenPeaks).length;
  }

  return total;
};

const part2 = (input: string) => {
  const grid: number[][] = input
    .split("\n")
    .map((a) => a.split("").map(Number));

  const trailheads = [];

  for (let y = 0; y < grid.length; y++) {
    for (let x = 0; x < grid[y].length; x++) {
      if (grid[y][x] === 0) {
        trailheads.push({ x, y });
      }
    }
  }

  let total = 0;

  for (const trailhead of trailheads) {
    const trail = traverse(
      0,
      trailhead,
      grid,
      grid[0].length - 1,
      grid.length - 1,
    );

    total += trail.length;
  }

  return total;
};

const input = fs.readFileSync("./input.txt", "utf8").trim();
console.log(`Part 1: ${part1(input)}`);
console.log(`Part 2: ${part2(input)}`);
