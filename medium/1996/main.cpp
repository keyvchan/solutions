#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
  int numberOfWeakCharacters(vector<vector<int>> &properties) {
    sort(properties.begin(), properties.end(), [](auto a, auto b) {
      if (a[0] == b[0]) {
        return a[1] > b[1];
      }
      return a[0] < b[0];
    });
    auto result = 0;
    // for every character
    auto current_max = INT_MIN;
    for (int i = properties.size() - 1; i >= 0; i--) {
      if (properties[i][1] < current_max) {
        result++;
      }
      current_max = max(current_max, properties[i][1]);
    }
    return result;
  }
};
