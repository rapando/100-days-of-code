#include <stdio.h>
#include <limits.h>
#include <stdlib.h>

// a structure to represent struct
struct Stack
{
    int top; // the top index. -1 if not available
    unsigned capacity;
    int *array; // the actual contents
};

// create a stack of a given capacity.
// it initializes the size as 0
struct Stack *createStack(unsigned capacity)
{
    struct Stack *stack = (struct Stack *)malloc(sizeof(struct Stack));
    stack->capacity = capacity;
    stack->top = -1;
    stack->array = (int *)malloc(stack->capacity * sizeof(int));
    return stack;
}

// Stack is full when top is equal to the last index
// e.g if top is 10 and capacity is 11, then the stack is full
int isFull(struct Stack *stack)
{
    return stack->top == stack->capacity - 1;
}

// Stack is empty when top is -1
int isEmpty(struct Stack *stack)
{
    return stack->top == -1;
}

// add an item to the stack and increase top by 1 at the same time
void push(struct Stack *stack, int item)
{
    if (isFull(stack))
    {
        // if the stack is full, don't push
        return;
    }
    stack->array[++stack->top] = item;
}

// remove the item from the stack and decrease top by 1 at the same time
int pop(struct Stack *stack)
{
    if (isEmpty(stack))
    {
        return INT_MIN;
    }
    return stack->array[stack->top--];
}

// get the top value from the stack without removing it
int peek(struct Stack *stack)
{
    if (isEmpty(stack))
    {
        return INT_MIN;
    }
    return stack->array[stack->top];
}
