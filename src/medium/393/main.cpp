#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
  int sequence_size(int data) {
    // get the sequence_size

    // check if data is 0xxxxxxx
    if ((data & 0x80) == 0x00) {
      // 1000 0000
      return 1;
    }
    if ((data & 0xE0) == 0xC0) {
      // 1110 0000
      return 2;
    }
    if ((data & 0xF0) == 0xE0) {
      // 1111 0000
      return 3;
    }
    if ((data & 0xF8) == 0xF0) {
      return 4;
    }
    return -1;
  }
  bool validUtf8(vector<int> &data) {
    // check the first byte
    auto pointer = 0;

    for (;;) {
      if (pointer >= data.size()) {
        break;
      }
      auto size = sequence_size(data[pointer]);
      if (size == -1) {
        return false;
      }
      // read size - 1 after current pointer
      int i = pointer + 1;
      for (; i < pointer + size; i++) {
        // check it
        if (i >= data.size()) {
          return false;
        }
        // check the most significant two bits
        if ((data[i] & 0xC0) != 0x80) {
          return false;
        }
      }
      pointer = i;
    }
    return true;
  }
};
