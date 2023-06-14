#include <bits/stdc++.h>

class Solution {
public:
  int numRescueBoats(std::vector<int> &people, int limit) {
    // we sort the people in ascending order
    std::sort(people.begin(), people.end());

    // we first take the weightiest person, then check if we can take more
    auto counts = 0;
    //
    auto left = 0;
    auto right = people.size() - 1;

    while (left <= right) {
      if (left == right) {
        // we can only take the last person
        counts++;
        break;
      }
      // check we have remaining space for the lightest person
      //
      // Noted that a boat can carry at most 2 people at the same time
      if (people[left] + people[right] <= limit) {
        // we can take lightest person and weightiest person, shift both
        left++;
        right--;

        counts++;
      } else {
        // we can just take the weightiest person
        right--;
        counts++;
      }
    }

    return counts;
  }
};
