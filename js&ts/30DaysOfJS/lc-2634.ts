type Fn = (n: number, i: number) => any;

function filter(arr: number[], fn: Fn): number[] {
  let filteredArr: number[] = [];
  arr.forEach((element, index) => {
    if (fn(element, index)) {
      filteredArr.push(element);
    }
  });
  return filteredArr;
}
