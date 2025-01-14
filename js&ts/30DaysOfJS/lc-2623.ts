type Fn = (...params: number[]) => number;

function memoize(fn: Fn): Fn {
  let dict = new Map();

  return function (...args) {
    let argsKey: string[] = [];
    args.forEach((element) => {
      argsKey.push(element.toString());
    });
    let key = argsKey.join(",");
    if (!dict.has(key)) {
      dict.set(key, fn(...args));
    }
    return dict.get(key);
  };
}

/**
 * let callCount = 0;
 * const memoizedFn = memoize(function (a, b) {
 *	 callCount += 1;
 *   return a + b;
 * })
 * memoizedFn(2, 3) // 5
 * memoizedFn(2, 3) // 5
 * console.log(callCount) // 1
 */
