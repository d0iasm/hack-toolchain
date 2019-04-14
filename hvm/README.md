# Hack VM Translator
VM translator from .vm to .asm written in Go. The virtual matchine is stack-based. You can use this VM translator just by `hvm <xxx.vm>` 
## Commands
- Arithmetic commands: perform arithmetic and logical operations on the stack.
- Memory access commands: transfer data between the stack and virtual memory segments.
- Program flow commands: facilitate conditional and unconditional branching operations.
- Function calling commands: call functions and return from them. 

### Arithmetic commands
- add: x+y
- sub: x-y
- neg: -y
- eq: true if x=y and false otherwise
- gt: true if x\>y and false otherwise
- lt: true if x\<y and false otherwise
- and: x and y, bit-wise
- or: x or y, bit-wise
- not: not y, bit-wise

### Memory access commands
- push \<segment index\>: push the value of segment[index] onth the stack.
- pop \<segment index\>: pop the topmost stack item and store its value in segment[index].

### Program flow commands
todo

### function calling commands
todo

## Memory Segments
- argument: Stores the function's arguments.
- local: Stores the functions' local variables.
- static: Stores static variables shared by all functions in the same .vm file.
- constant: Pseudo-segment that holds all the constants in the range 0...32767.
- this/that: General-purpose segments that can be made to correspond to different areas in the heap. Serve various programming needs.
- pointer: Fixed 2-entry segment that holds the base addresses of this and that.
- temp: Fixed 8-entry segment that holds temporary variables for general use.

## RAM addresses Usage
- 0 ~ 15: Sixteen virtual registers, whose usage is described below
- 16 ~ 255: Static variables (of all the VM functions in the VM program)
- 256 ~ 2047: Stack
- 2048 ~ 16483: Heap (used to store objects and arrays)
- 16384 ~ 24575: Memory mapped I/O

## Registers
- RAM[0] SP Stack pointer: points to the next topmost location in the stack.
- RAM[1] LCL Points to the base of the current VM function's local segment.
- RAM[2] ARG Points to the base of the current VM function 's argument segment.
- RAM[3] THIS Points to the base of the current this segment (within the heap).
- RAM[4] THAT Points to the base of the current that segment (within the heap).
- RAM[5-12] TEMP Hold the contents of the temp segment.
- RAM[13-15] (-) Can be used by the VM implementation as general-purpose registers. 

## Assembly
### Registers as an assembly layer
- D: 16-bit register used to store data values.
- A: 16-bit register used as both a data register and an address register depends on the instruction context.
  - A register can access to the memory directly.

### Example Code
```
@16 // Store the number 16 to A register.

A=D-1 // Store the result of D minus 1 to A register.
D=!A // Store the result of bit NOT opeation of A to D register.

// M always refers to the memory word whose address is the current value of the A register.
D=M+1 // Store the result of Memory[A's data] plus 1 to D register.

// D = Memory[516] - 1
@516
D=M-1
```

## Implementation Steps
1. Handling stack arithmetic commands, which exists 9 types.
2. Handling memory access commands.
  a. Constant segment is already implemented at step 1.
  b. Local, argument, this, and that segments.
  c. Pointer and temp segments, in particular allowing modification of the bases of the this and that segments.
  d. Static segment.
3. todo


## Requirements
- go version go1.11.4
