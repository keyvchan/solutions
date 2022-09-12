#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
  int bagOfTokensScore(vector<int> &tokens, int power) {
    if (tokens.size() == 0) {
      return 0;
    }
    sort(tokens.begin(), tokens.end());

    auto left = 0;
    auto right = tokens.size() - 1;

    auto score = 0;
    auto maxScore = 0;

    for (;;) {
      if (left > right) {
        break;
      }
      // check if we can gain score
      if (power >= tokens[left]) {
        // take it
        power -= tokens[left];
        score++;
        left++;

      } else {
        // check if we can gain power
        if (score >= 1) {
          // we can gain power
          power += tokens[right];
          score -= 1;
          right--;
        } else {
          // we can't gain power
          break;
        }
      }
      maxScore = max(maxScore, score);
    }
    return maxScore;
  }
};

int main() {
  Solution s;
  auto tokens = vector<int>{52, 65, 35, 88, 28, 1, 4, 68, 56, 95};
  auto result = s.bagOfTokensScore(tokens, 94);
  cout << result << endl;
}
