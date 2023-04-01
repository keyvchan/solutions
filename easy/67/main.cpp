#include <algorithm>
#include <string>

class Solution {
public:
  std::string addBinary(std::string a, std::string b) {
    // create a string view
    auto a_view = std::string_view(a);
    auto b_view = std::string_view(b);

    // from back to front
    auto a_it = a_view.rbegin();
    auto b_it = b_view.rbegin();

    // carry
    int carry = 0;

    // result
    std::string result;

    // loop until both a and b are empty
    while (a_it != a_view.rend() || b_it != b_view.rend()) {
      // get the current digit
      int a_digit = a_it != a_view.rend() ? *a_it - '0' : 0;
      int b_digit = b_it != b_view.rend() ? *b_it - '0' : 0;

      // calculate the sum
      int sum = a_digit + b_digit + carry;

      // update the carry
      carry = sum / 2;

      // update the result
      result.push_back(sum % 2 + '0');

      // move to the next digit
      if (a_it != a_view.rend()) {
        ++a_it;
      }
      if (b_it != b_view.rend()) {
        ++b_it;
      }
    }

    // if there is still a carry, add it to the result
    if (carry) {
      result.push_back(carry + '0');
    }

    // reverse the result
    std::reverse(result.begin(), result.end());

    return result;
  }
};
