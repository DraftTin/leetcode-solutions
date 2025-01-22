class Calculator {
  /**
   * @param {number} value
   */
  constructor(value) {
    this.result = value;
    this.error = null;
  }

  /**
   * @param {number} value
   * @return {Calculator}
   */
  add(value) {
    this.result += value;
    return this;
  }

  /**
   * @param {number} value
   * @return {Calculator}
   */
  subtract(value) {
    this.result -= value;
    return this;
  }

  /**
   * @param {number} value
   * @return {Calculator}
   */
  multiply(value) {
    this.result *= value;
    return this;
  }

  /**
   * @param {number} value
   * @return {Calculator}
   */
  divide(value) {
    if (value == 0) {
      this.error = "Division by zero is not allowed";
    } else {
      this.result /= value;
    }
    return this;
  }

  /**
   * @param {number} value
   * @return {Calculator}
   */
  power(value) {
    this.result = Math.pow(this.result, value);
    return this;
  }

  /**
   * @return {number}
   */
  getResult() {
    if (this.error) {
      return this.error;
    }
    return this.result;
  }
}
