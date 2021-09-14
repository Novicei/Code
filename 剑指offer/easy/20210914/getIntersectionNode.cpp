class Solution {
public:
    ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
        ListNode *cur1=headA,*cur2=headB;
        while(cur1!=cur2){
            cur1=cur1==NULL?headB:cur1->next;
            cur2=cur2==NULL?headA:cur2->next;
        }
        return cur1;
    }
};
