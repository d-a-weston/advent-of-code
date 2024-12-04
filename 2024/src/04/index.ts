import fs from "fs";

const dirs = {
  up: [-1, 0],
  upRight: [-1, 1],
  right: [0, 1],
  downRight: [1, 1],
  down: [1, 0],
  downLeft: [1, -1],
  left: [0, -1],
  upLeft: [-1, -1],
};

const part1 = (input: string) => {
  const grid = input.split("\n").map((a) => a.split(""));

  let total = 0;

  for (let y = 0; y < grid.length; y++) {
    for (let x = 0; x < grid[y].length; x++) {
      if (grid[y][x] === "X") {
        for (const dir of Object.values(dirs)) {
          const [modY, modX] = dir;

          let curY = y;
          let curX = x;

          for (const char of ["M", "A", "S"]) {
            curY += modY;
            curX += modX;

            if (grid[curY]?.[curX] !== char) break;
            if (char === "S") total += 1;
          }
        }
      }
    }
  }

  return total;
};

const part2 = (input: string) => {
  const grid = input.split("\n").map((a) => a.split(""));

  let total = 0;

  for (let y = 0; y < grid.length; y++) {
    for (let x = 0; x < grid[y].length; x++) {
      if (grid[y][x] === "A") {
        const diag1 = `${grid[y - 1]?.[x - 1]}${grid[y + 1]?.[x + 1]}`;
        const diag2 = `${grid[y + 1]?.[x - 1]}${grid[y - 1]?.[x + 1]}`;

        if (
          (diag1 === "SM" || diag1 === "MS") &&
          (diag2 === "SM" || diag2 === "MS")
        )
          total += 1;
      }
    }
  }

  return total;
};

const input = fs.readFileSync("./input.txt", "utf8").trim();
console.log(`Part 1: ${part1(input)}`);
console.log(`Part 2: ${part2(input)}`);
