func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    } else if list2 == nil {
        return list1
    }
    //check, out of two linked list which linked list have lowest value on first node
    var head, curr *ListNode
    if list1.Val < list2.Val {
        head = list1
        curr = list1
        list1 = list1.Next //moving to next node of list1,
    } else {
        head = list2
        curr = list2
        list2 = list2.Next//moving to next node of list2
    }
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            curr.Next = list1
            curr = curr.Next
            list1 = list1.Next
        } else {
            curr.Next = list2
            curr = curr.Next
            list2 = list2.Next
        }
    }
    if list1 != nil {
        curr.Next = list1
    } else if list2 != nil {
        curr.Next = list2
    }
    return head
}