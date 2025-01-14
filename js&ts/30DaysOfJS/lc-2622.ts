interface CachedValue {
  value: number;
  timeOutId: ReturnType<typeof setTimeout>;
}

class TimeLimitedCache {
  dict: Map<number, CachedValue>;
  constructor() {
    this.dict = new Map<number, CachedValue>();
  }

  set(key: number, value: number, duration: number): boolean {
    let res = false;
    if (this.dict.has(key)) {
      res = true;
      clearTimeout(this.dict.get(key)?.timeOutId);
    }
    let setTimeoutId = setTimeout(() => {
      this.dict.delete(key);
    }, duration);
    this.dict.set(key, { value: value, timeOutId: setTimeoutId });
    return res;
  }

  get(key: number): number {
    if (this.dict.has(key)) {
      return this.dict.get(key).value;
    }
    return -1;
  }

  count(): number {
    return this.dict.size;
  }
}

/**
 * const timeLimitedCache = new TimeLimitedCache()
 * timeLimitedCache.set(1, 42, 1000); // false
 * timeLimitedCache.get(1) // 42
 * timeLimitedCache.count() // 1
 */
