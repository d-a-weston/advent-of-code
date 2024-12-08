import fs from "fs";

const isInBounds = (width: number, height: number, x: number, y: number) => {
  return x >= 0 && x <= width && y >= 0 && y <= height;
};

const part1 = (input: string) => {
  const grid: string[][] = input.split("\n").map((a) => a.split(""));

  const height = grid.length - 1;
  const width = grid[0].length - 1;

  const nodeLocations: Record<string, Record<string, number>[]> = {};

  for (let y = 0; y < grid.length; y++) {
    for (let x = 0; x < grid[y].length; x++) {
      const cell = grid[y][x];

      if (cell === ".") continue;

      if (nodeLocations[cell]) nodeLocations[cell].push({ x, y });
      else nodeLocations[cell] = [{ x, y }];
    }
  }

  const knownAntinodes: Record<string, boolean> = {};
  let total = 0;

  for (const locations of Object.values(nodeLocations)) {
    for (let i = 0; i < locations.length - 1; i++) {
      for (let j = i + 1; j < locations.length; j++) {
        const nodeA = locations[i];
        const nodeB = locations[j];

        const deltaX = nodeA.x - nodeB.x;
        const deltaY = nodeA.y - nodeB.y;

        const antiNodeA = {
          x: nodeA.x + deltaX,
          y: nodeA.y + deltaY,
        };

        const antiNodeB = {
          x: nodeB.x - deltaX,
          y: nodeB.y - deltaY,
        };

        if (
          isInBounds(width, height, antiNodeA.x, antiNodeA.y) &&
          !knownAntinodes[`${antiNodeA.x},${antiNodeA.y}`]
        ) {
          knownAntinodes[`${antiNodeA.x},${antiNodeA.y}`] = true;
          total += 1;
        }

        if (
          isInBounds(width, height, antiNodeB.x, antiNodeB.y) &&
          !knownAntinodes[`${antiNodeB.x},${antiNodeB.y}`]
        ) {
          knownAntinodes[`${antiNodeB.x},${antiNodeB.y}`] = true;
          total += 1;
        }
      }
    }
  }

  return total;
};

const part2 = (input: string) => {
  const grid: string[][] = input.split("\n").map((a) => a.split(""));

  const height = grid.length - 1;
  const width = grid[0].length - 1;

  const nodeLocations: Record<string, Record<string, number>[]> = {};
  const knownAntinodes: Record<string, boolean> = {};

  let total = 0;

  for (let y = 0; y < grid.length; y++) {
    for (let x = 0; x < grid[y].length; x++) {
      const cell = grid[y][x];

      if (cell === ".") continue;

      if (nodeLocations[cell]) nodeLocations[cell].push({ x, y });
      else nodeLocations[cell] = [{ x, y }];

      knownAntinodes[`${x},${y}`] = true;
      total += 1;
    }
  }

  for (const locations of Object.values(nodeLocations)) {
    for (let i = 0; i < locations.length - 1; i++) {
      for (let j = i + 1; j < locations.length; j++) {
        const nodeA = locations[i];
        const nodeB = locations[j];

        const deltaX = nodeA.x - nodeB.x;
        const deltaY = nodeA.y - nodeB.y;

        const antiNodeA = {
          x: nodeA.x + deltaX,
          y: nodeA.y + deltaY,
        };

        while (isInBounds(width, height, antiNodeA.x, antiNodeA.y)) {
          if (!knownAntinodes[`${antiNodeA.x},${antiNodeA.y}`]) {
            knownAntinodes[`${antiNodeA.x},${antiNodeA.y}`] = true;
            total += 1;
          }

          antiNodeA.x += deltaX;
          antiNodeA.y += deltaY;
        }

        const antiNodeB = {
          x: nodeB.x - deltaX,
          y: nodeB.y - deltaY,
        };

        while (isInBounds(width, height, antiNodeB.x, antiNodeB.y)) {
          if (!knownAntinodes[`${antiNodeB.x},${antiNodeB.y}`]) {
            knownAntinodes[`${antiNodeB.x},${antiNodeB.y}`] = true;
            total += 1;
          }

          antiNodeB.x -= deltaX;
          antiNodeB.y -= deltaY;
        }
      }
    }
  }

  return total;
};

const input = fs.readFileSync("./input.txt", "utf8").trim();
console.log(`Part 1: ${part1(input)}`);
console.log(`Part 2: ${part2(input)}`);
