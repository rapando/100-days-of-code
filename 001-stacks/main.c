#include <stdio.h>
#include "stack.h"

int main()
{
    printf("Make a stack with capacity of 4\n");
    struct Stack *stack = createStack(4);

    printf("push 1,2,3,4\n");
    push(stack, 1);
    push(stack, 2);
    push(stack, 3);
    push(stack, 4);
    printf("Peek : %d\n", peek(stack));

    printf("Pop : %d\n", pop(stack));
    printf("Pop : %d\n", pop(stack));
    printf("Pop : %d\n", pop(stack));
    printf("Pop : %d\n", pop(stack));
    printf("Pop : %d\n", pop(stack)); // here we have read everything, we print INT_MIN

    return 0;
}
