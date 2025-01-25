class Solution:
    def countAndSay(self, n: int) -> str:
        res = ["1"]
        for i in range(n - 1):
            newRes = []
            j = 0
            while j < len(res):
                k = j + 1
                while k < len(res) and res[k] == res[j]:
                    k += 1
                newRes.append(str(k - j))
                newRes.append(res[j])
                j = k

            res = newRes
        return "".join(res)
