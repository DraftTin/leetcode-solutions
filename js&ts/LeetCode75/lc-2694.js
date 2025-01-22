class EventEmitter {
  constructor() {
    this.emitter = new Map();
  }
  /**
   * @param {string} eventName
   * @param {Function} callback
   * @return {Object}
   */
  subscribe(eventName, callback) {
    if (!this.emitter.has(eventName)) {
      this.emitter.set(eventName, []);
    }
    this.emitter.get(eventName).push(callback);
    return {
      unsubscribe: () => {
        this.emitter.get(eventName).forEach((func, index) => {
          if (func == callback) {
            this.emitter.get(eventName).splice(index, 1);
          }
        });
      },
    };
  }

  /**
   * @param {string} eventName
   * @param {Array} args
   * @return {Array}
   */
  emit(eventName, args = []) {
    let res = [];
    if (!this.emitter.has(eventName)) {
      return res;
    }
    this.emitter.get(eventName).forEach((f) => {
      res.push(f(...args));
    });
    return res;
  }
}

/**
 * const emitter = new EventEmitter();
 *
 * // Subscribe to the onClick event with onClickCallback
 * function onClickCallback() { return 99 }
 * const sub = emitter.subscribe('onClick', onClickCallback);
 *
 * emitter.emit('onClick'); // [99]
 * sub.unsubscribe(); // undefined
 * emitter.emit('onClick'); // []
 */
