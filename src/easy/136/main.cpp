#include <bits/stdc++.h>

class Solution {
public:
  int singleNumber(std::vector<int> &nums) {
    // xor all the numbers
    int result = 0;
    // range based for loop
    for (auto i : nums) {
      result ^= i;
    }
    return result;
  }
};
