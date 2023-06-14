#include <string>
#include <unordered_set>

class Solution {
public:
  int partitionString(std::string s) {
    // from start to end
    auto counts = 0;

    // create a set to store the unique characters
    std::unordered_set<char> uniqueChars;

    // range based for loop
    for (auto ch : s) {
      if (uniqueChars.find(ch) == uniqueChars.end()) {
        uniqueChars.insert(ch);
      } else {
        uniqueChars.clear();
        uniqueChars.insert(ch);
        counts++;
      }
    }

    if (uniqueChars.size() > 0) {
      counts++;
    }

    return counts;
  }
};
