#include <cassert>
#include <string>
using namespace std;

class Solution {
public:
  string mergeAlternately(string word1, string word2) {
    string res;

    int i = 0;
    for (i = 0; i < min(word1.size(), word2.size()); i++) {

      res += word1[i];
      res += word2[i];
    }
    if (i == word1.size()) {
      res += string_view(word2.data() + i, word2.size() - i);
    }
    if (i == word2.size()) {
      res += string_view(word1.data() + i, word1.size() - i);
    }

    return res;
  }
};

int main() {
  Solution s;
  // check output
  assert(s.mergeAlternately("abc", "pqr") == "apbqcr");
  assert(s.mergeAlternately("ab", "pqrs") == "apbqrs");
  assert(s.mergeAlternately("abcd", "pq") == "apbqcd");
}
