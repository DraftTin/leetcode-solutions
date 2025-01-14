function map(arr: number[], fn: (n: number, i: number) => number): number[] {
  let result: number[] = [];
  arr.forEach((element, index) => {
    result.push(fn(element, index));
  });
  return result;
}
