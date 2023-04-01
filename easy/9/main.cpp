#include <iostream>
#include <string>

class Solution {
public:
  bool isPalindrome(int x) {
    // so we have a int, check if its a negative number
    //
    if (x < 0) {
      return false;
    }

    // its not a negative number, so we need to check if its a palindrome

    // convert the int to a string
    auto str = std::to_string(x);
    // use string_view
    auto str_view = std::string_view(str);

    // check if the string is a palindrome
    for (int i = 0; i < str_view.size() / 2; i++) {
      if (str_view[i] != str_view[str_view.size() - i - 1]) {
        return false;
      }
    }

    return true;
  }
};

int main() {
  Solution s;
  std::cout << s.isPalindrome(121) << std::endl;
  return 0;
}
