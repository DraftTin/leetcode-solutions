/**
 * @param {string} s
 * @return {number}
 */
var partitionString = function (s) {
  let dict = {};
  let res = 1;
  for (let i = 0; i < s.length; i++) {
    if (dict[s[i]] == true) {
      dict = {};
      res++;
    }
    dict[s[i]] = true;
  }
  return res;
};
