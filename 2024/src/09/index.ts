import fs from "fs";

const part1 = (input: string) => {
  let disk = [];
  let id = 0;

  for (let i = 0; i < input.length; i++) {
    if (i % 2 === 0) {
      const len = parseInt(input[i]);

      for (let j = 0; j < len; j++) {
        disk.push(id);
      }

      id++;
    } else {
      const len = parseInt(input[i]);

      for (let j = 0; j < len; j++) {
        disk.push(".");
      }
    }
  }

  let total = 0;
  let j = disk.length - 1;

  for (let i = 0; i < disk.length; i++) {
    if (disk[i] === "." && j >= i) {
      while (disk[j] === "." && j > i) j--;

      disk[i] = disk[j];
      disk[j] = ".";
      j--;
    }

    if (disk[i] === ".") continue;

    total += parseInt(`${disk[i]}`) * i;
  }

  return total;
};

const part2 = (input: string) => {
  let disk = [];
  let id = 0;

  for (let i = 0; i < input.length; i++) {
    if (i % 2 === 0) {
      const len = parseInt(input[i]);
      for (let j = 0; j < len; j++) {
        disk.push(id);
      }
      id++;
    } else {
      const len = parseInt(input[i]);
      for (let j = 0; j < len; j++) {
        disk.push(".");
      }
    }
  }

  let total = 0;

  for (let currentId = id - 1; currentId >= 0; currentId--) {
    let fileStart = disk.indexOf(currentId);
    if (fileStart === -1) continue;

    let fileLength = 0;
    while (
      fileStart + fileLength < disk.length &&
      disk[fileStart + fileLength] === currentId
    ) {
      fileLength++;
    }

    let consecutiveSpaces = 0;
    let moveToIndex = -1;

    for (let i = 0; i < fileStart; i++) {
      if (disk[i] === ".") {
        consecutiveSpaces++;
        if (consecutiveSpaces === fileLength) {
          moveToIndex = i - fileLength + 1;
          break;
        }
      } else {
        consecutiveSpaces = 0;
      }
    }

    if (moveToIndex !== -1) {
      for (let i = 0; i < fileLength; i++) {
        disk[moveToIndex + i] = currentId;
        disk[fileStart + i] = ".";
      }
    }
  }

  for (let i = 0; i < disk.length; i++) {
    if (disk[i] === ".") continue;
    total += parseInt(`${disk[i]}`) * i;
  }

  return total;
};

const input = fs.readFileSync("./input.txt", "utf8").trim();
console.log(`Part 1: ${part1(input)}`);
console.log(`Part 2: ${part2(input)}`);
