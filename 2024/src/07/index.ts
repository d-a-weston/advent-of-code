import fs from "fs";

const findPossibleResults = (result: number, values: number[]): number[] => {
  if (values.length === 0) return [result];

  const ifAdded = result + values[0];
  const ifMultiplied = result * values[0];

  return [
    ...findPossibleResults(ifAdded, values.slice(1)),
    ...findPossibleResults(ifMultiplied, values.slice(1)),
  ];
};

const findPossibleResultsWithConcat = (
  result: number,
  values: number[],
  concatted: boolean,
): number[] => {
  if (values.length === 0) return [result];

  const ifAdded = result + values[0];
  const ifMultiplied = result * values[0];
  const ifConcatted = parseInt(`${result}${values[0]}`);

  return [
    ...findPossibleResultsWithConcat(ifAdded, values.slice(1), concatted),
    ...findPossibleResultsWithConcat(ifMultiplied, values.slice(1), concatted),
    ...findPossibleResultsWithConcat(ifConcatted, values.slice(1), true),
  ];
};

const part1 = (input: string) => {
  const lines = input.split("\n");

  let total = 0;

  for (const line of lines) {
    const [result, ...values] = line.match(/\d+/g)?.map(Number) || [];

    const possibleResults: number[] = findPossibleResults(
      values[0],
      values.slice(1),
    );

    if (possibleResults.includes(result)) total += result;
  }

  return total;
};

const part2 = (input: string) => {
  const lines = input.split("\n");

  let total = 0;

  for (const line of lines) {
    const [result, ...values] = line.match(/\d+/g)?.map(Number) || [];

    const possibleResults: number[] = findPossibleResultsWithConcat(
      values[0],
      values.slice(1),
      false,
    );

    if (possibleResults.includes(result)) total += result;
  }

  return total;
};

const input = fs.readFileSync("./input.txt", "utf8").trim();
console.log(`Part 1: ${part1(input)}`);
console.log(`Part 2: ${part2(input)}`);
