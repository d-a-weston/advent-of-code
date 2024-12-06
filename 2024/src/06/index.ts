import fs from "fs";

interface Guard {
  x: number;
  y: number;
  facing: number;
}

const dirs: Record<number, number[]> = {
  0: [-1, 0],
  1: [0, 1],
  2: [1, 0],
  3: [0, -1],
};

const findStart = (grid: string[][]): Guard => {
  for (let y = 0; y < grid.length; y++) {
    for (let x = 0; x < grid[y].length; x++) {
      const cell = grid[y][x];

      switch (cell) {
        case "^":
          return { x, y, facing: 0 };
        case ">":
          return { x, y, facing: 1 };
        case "v":
          return { x, y, facing: 2 };
        case "<":
          return { x, y, facing: 3 };
      }
    }
  }

  return { x: 0, y: 0, facing: 0 };
};

const part1 = (input: string) => {
  const grid = input.split("\n").map((line) => line.split(""));

  const guard = findStart(grid);

  grid[guard.y][guard.x] = ".";

  const visited: Record<string, boolean> = {};

  let total = 0;
  let inBounds = true;

  while (inBounds) {
    const [modY, modX] = dirs[guard.facing];
    const lookAhead = grid[guard.y + modY]?.[guard.x + modX];

    if (!visited[`${guard.y},${guard.x}`]) {
      visited[`${guard.y},${guard.x}`] = true;
      total += 1;
    }

    switch (lookAhead) {
      case ".":
        guard.x += modX;
        guard.y += modY;
        break;
      case "#":
        guard.facing = (guard.facing + 1) % 4;
        break;
      case undefined:
        inBounds = false;
        break;
    }
  }

  return total;
};

const containsInfiniteLoop = (guard: Guard, grid: string[][]) => {
  const visited: Record<string, boolean> = {};

  let retracing = 0;
  let inBounds = true;

  while (inBounds) {
    const [modY, modX] = dirs[guard.facing];
    const lookAhead = grid[guard.y + modY]?.[guard.x + modX];

    if (!visited[`${guard.y},${guard.x}`]) {
      visited[`${guard.y},${guard.x}`] = true;
      retracing = 0;
    } else {
      retracing += 1;
    }

    if (retracing === 1000) return true;

    switch (lookAhead) {
      case ".":
        guard.x += modX;
        guard.y += modY;
        break;
      case "#":
        guard.facing = (guard.facing + 1) % 4;
        break;
      case undefined:
        inBounds = false;
        break;
    }
  }

  return false;
};

const part2 = (input: string) => {
  const grid = input.split("\n").map((line) => line.split(""));

  const guard = findStart(grid);

  grid[guard.y][guard.x] = ".";

  let total = 0;
  let test = 0;

  for (let y = 0; y < grid.length; y++) {
    for (let x = 0; x < grid[y].length; x++) {
      if (grid[y][x] === "#" || (guard.x === x && guard.y === y)) continue;

      const newGrid = JSON.parse(JSON.stringify(grid));
      newGrid[y][x] = "#";

      if (containsInfiniteLoop({ ...guard }, newGrid)) {
        total += 1;
      }

      test += 1;
    }
  }

  return total;
};

const input = fs.readFileSync("./input.txt", "utf8").trim();
console.log(`Part 1: ${part1(input)}`);
console.log(`Part 2: ${part2(input)}`);
