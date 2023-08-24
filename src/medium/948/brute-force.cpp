#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
  int dfs(vector<int> &tokens, int power, int score, unordered_set<int> &used) {
    if (used.size() == tokens.size()) {
      // we used all tokens, return
      return score;
    }
    auto result = 0;
    for (int i = 0; i < tokens.size(); i++) {
      if (used.find(i) != used.end()) {
        // we already use this token
        continue;
      }
      if (power >= tokens[i]) {
        // power is at least tokens[i], may play the current token or not

        // use this token
        used.insert(i);
        result = max(result, dfs(tokens, power - tokens[i], score + 1, used));
        used.erase(i);

        // not use this token
        used.insert(i);
        result = max(result, dfs(tokens, power, score, used));
        used.erase(i);
      }

      if (score >= 1) {
        // we may play another ways

        // use this token
        used.insert(i);
        result = max(result, dfs(tokens, power + tokens[i], score - 1, used));
        used.erase(i);

        used.insert(i);
        result = max(result, dfs(tokens, power, score, used));
        used.erase(i);
      }
    }
    return result;
  }

  int bagOfTokensScore(vector<int> &tokens, int power) {
    // sort the tokens
    sort(tokens.begin(), tokens.end());

    // used
    auto used = unordered_set<int>();

    return dfs(tokens, power, 0, used);
  }
};

int main() {
  Solution s;
  auto tokens = vector<int>{52, 65, 35, 88, 28, 1, 4, 68, 56, 95};
  auto result = s.bagOfTokensScore(tokens, 94);
  cout << result << endl;
}
