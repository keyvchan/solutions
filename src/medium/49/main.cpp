#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
  vector<vector<string>> groupAnagrams(vector<string> &strs) {
    auto result = unordered_map<string, vector<string>>();

    // count word frequence of every number
    for (auto &str : strs) {
      auto temp = str;
      sort(str.begin(), str.end());
      result[str].push_back(temp);
    }

    auto result_vec = vector<vector<string>>();
    for (auto val : result) {
      result_vec.push_back(val.second);
    }

    return result_vec;
  }
};
