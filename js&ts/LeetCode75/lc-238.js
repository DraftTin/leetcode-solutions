/**
 * @param {number[]} nums
 * @return {number[]}
 */
var productExceptSelf = function (nums) {
  let left = new Array(nums.length);
  let right = new Array(nums.length);
  let val = 1;
  for (let i = 0; i < nums.length; ++i) {
    val *= nums[i];
    left[i] = val;
  }
  val = 1;
  for (let i = nums.length - 1; i >= 0; --i) {
    val *= nums[i];
    right[i] = val;
  }
  let res = Array(nums.length);
  res[0] = right[1];
  res[nums.length - 1] = left[nums.length - 2];
  for (let i = 1; i < nums.length - 1; ++i) {
    res[i] = left[i - 1] * right[i + 1];
  }
  return res;
};
