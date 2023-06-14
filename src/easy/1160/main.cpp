#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
  int countCharacters(vector<string> &words, string chars) {
    // count word frequency
    int chars_freq[26] = {};

    for (auto c : chars) {
      chars_freq[c - 'a']++;
    }

    auto total = 0;

    for (auto word : words) {
      int current_word_freq[26] = {};

      auto can_flag = true;

      for (auto w : word) {
        current_word_freq[w - 'a']++;
      }

      for (auto i = 0; i < 26; i++) {
        if (current_word_freq[i] > chars_freq[i]) {
          can_flag = false;
          break;
        }
      }
      if (can_flag == true) {
        total += word.size();
      }
    }
    return total;
  }
};
