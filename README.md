This program can be used to create a unique ring "buffer" stack of stack size "length"

When you add element in stack, it checks if alreaedy present. In that case discard and return false else add and return true

Note the stack never grows beyond "length" elements. If it grows, the earlier elements are trimmed to maintain the length of stack as defined during initialisation
