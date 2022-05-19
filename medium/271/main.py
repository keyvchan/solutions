import base64


class Solution:
    def encode(self, strs):
        """
        @param: strs: a list of strings
        @return: encodes a list of strings to a single string.
        """
        # write your code here
        ...
        final_str = ""
        for i, str in enumerate(strs):
            final_str += base64.b64encode(str.encode("utf-8")).decode("utf-8")
            if i != len(strs) - 1:
                final_str += "_"
        return final_str

    def decode(self, str):
        """
        @param: str: A string
        @return: dcodes a single string to a list of strings
        """
        # write your code here
        strs = str.split("_")
        for i, str in enumerate(strs):
            strs[i] = base64.b64decode(str).decode("utf-8")
        return strs


Solution().decode(Solution().encode(["hello", "world"]))
