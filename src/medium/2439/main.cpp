#include <vector>

class Solution {
public:
  int minimizeArrayValue(std::vector<int> &nums) {
    long long answer = 0, prefixSum = 0;
    for (int i = 0; i < nums.size(); ++i) {
      prefixSum += nums[i];
      answer = std::max(answer, (prefixSum + i) / (i + 1));
    }

    return answer;
  }
};
