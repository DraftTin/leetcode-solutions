type Fn = (...params: any[]) => Promise<any>;

function timeLimit(fn: Fn, t: number): Fn {
  return async function (...args) {
    return new Promise((resolve, reject) => {
      let clearTimeoutId = setTimeout(() => {
        clearTimeout(clearTimeoutId);
        reject("Time Limit Exceeded");
      });
      fn(...args)
        .then((result) => {
          clearTimeout(clearTimeoutId);
          resolve(result);
        })
        .catch((error) => {
          clearTimeout(clearTimeoutId);
          reject(error);
        });
    });
  };
}

/**
 * const limited = timeLimit((t) => new Promise(res => setTimeout(res, t)), 100);
 * limited(150).catch(console.log) // "Time Limit Exceeded" at t=100ms
 */
