#include <algorithm>
#include <cassert>
#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
  vector<bool> kidsWithCandies(vector<int> &candies, int extraCandies) {
    auto max = *max_element(candies.begin(), candies.end());

    vector<bool> res;
    res.reserve(candies.size());

    for (auto &cur : candies) {
      res.emplace_back(cur + extraCandies >= max);
    }

    return res;
  }
};

int main() {
  vector<int> candies1 = {2, 3, 5, 1, 3};
  assert(Solution().kidsWithCandies(candies1, 3) ==
         vector<bool>({true, true, true, false, true}));
  vector<int> candies2 = {4, 2, 1, 1, 2};
  assert(Solution().kidsWithCandies(candies2, 1) ==
         vector<bool>({true, false, false, false, false}));
  vector<int> candies3 = {12, 1, 12};
  assert(Solution().kidsWithCandies(candies3, 10) ==
         vector<bool>({true, false, true}));
  return 0;
}
