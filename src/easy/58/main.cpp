#include <iostream>
#include <string>
#include <string_view>

class Solution {
public:
  int lengthOfLastWord(std::string s) {
    // split string using string_view
    std::string_view sv(s);
    std::string_view delimiter(" ");
    std::string_view token;
    size_t pos = 0;
    auto last_word_length = 0;
    while ((pos = sv.find(delimiter)) != std::string_view::npos) {
      token = sv.substr(0, pos);
      // std::cout << token << token.length() << std::endl;
      if (token.length() > 0) {
        last_word_length = token.length();
      }
      sv.remove_prefix(pos + delimiter.length());
    }
    return sv.length() > 0 ? sv.length() : last_word_length;
  }
};

int main() {
  Solution s;
  std::cout << s.lengthOfLastWord("Hello World") << std::endl;
  std::cout << s.lengthOfLastWord("   fly me   to   the moon  ") << std::endl;
  return 0;
}
