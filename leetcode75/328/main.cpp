#include <iostream>
#include <vector>
using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class Solution {
  public:
    ListNode *oddEvenList(ListNode *head) {
        // if nodes is 0
        if (head == nullptr) {
            return nullptr;
        }

        // if nodes is 1
        if (head->next == nullptr) {
            return head;
        }

        ListNode *oddHead = head;
        ListNode *oddTail = head;
        ListNode *evenHead = head->next;
        ListNode *evenTail = head->next;

        // let's have a counter to keep track of the nodes
        int counter = 1;

        ListNode *tail = head->next->next;

        while (tail != nullptr) {
            if (counter % 2 == 0) {
                // even, append to evenTail
                evenTail->next = tail;
                tail = tail->next;
                evenTail = evenTail->next;
                evenTail->next = nullptr;
            } else {
                // odd append to oddTail
                oddTail->next = tail;
                tail = tail->next;
                oddTail = oddTail->next;
                oddTail->next = nullptr;
            }

            counter++;
        }

        oddTail->next = evenHead;
        evenTail->next = nullptr;
        return oddHead;
    }
};

int main() {
    Solution s;
    ListNode *head = new ListNode(1);
    head->next = new ListNode(2);
    head->next->next = new ListNode(3);
    head->next->next->next = new ListNode(4);
    head->next->next->next->next = new ListNode(5);
    auto newHead = s.oddEvenList(head);

    while (newHead != nullptr) {
        cout << newHead->val << endl;
        newHead = newHead->next;
    }

    ListNode *head2 = new ListNode(2);
    head2->next = new ListNode(1);
    head2->next->next = new ListNode(3);
    head2->next->next->next = new ListNode(5);
    head2->next->next->next->next = new ListNode(6);
    head2->next->next->next->next->next = new ListNode(4);
    head2->next->next->next->next->next->next = new ListNode(7);
    auto newHead2 = s.oddEvenList(head2);

    while (newHead2 != nullptr) {
        cout << newHead2->val << endl;
        newHead2 = newHead2->next;
    }

    // 1 2 3
    ListNode *head3 = new ListNode(1);
    head3->next = new ListNode(2);
    head3->next->next = new ListNode(3);

    auto newHead3 = s.oddEvenList(head);

    while (newHead3 != nullptr) {
        cout << newHead3->val << endl;
        newHead3 = newHead3->next;
    }
    return 0;
}
