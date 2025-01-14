type Counter = {
  increment: () => number;
  decrement: () => number;
  reset: () => number;
};

function createCounter(init: number): Counter {
  let initVal = init;
  let curVal = init;

  return {
    increment() {
      return ++curVal;
    },
    decrement() {
      return --curVal;
    },
    reset() {
      return (curVal = initVal);
    },
  };
}

/**
 * const counter = createCounter(5)
 * counter.increment(); // 6
 * counter.reset(); // 5
 * counter.decrement(); // 4
 */
