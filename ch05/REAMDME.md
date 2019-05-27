# 
字节码中存放编码后的Java虚拟机指令。每条指令都以一个单
字节的操作码（opcode）开头，这就是字节码名称的由来。由于只使
用一字节表示操作码，显而易见，Java虚拟机最多只能支持256（2
8 ）
条指令。到第八版为止，Java虚拟机规范已经定义了205条指令，操
作码分别是0（0x00）到202（0xCA）、254（0xFE）和255（0xFF）。这205
条指令构成了Java虚拟机的指令集（instruction set）。和汇编语言类
似，为了便于记忆，Java虚拟机规范给每个操作码都指定了一个助
记符（mnemonic）。比如操作码是0x00这条指令，因为它什么也不
做，所以它的助记符是nop（no operation）。
Java虚拟机使用的是变长指令，操作码后面可以跟零字节或多
字节的操作数（operand）。如果把指令想象成函数的话，操作数就是
它的参数。为了让编码后的字节码更加紧凑，很多操作码本身就隐
含了操作数，比如常数0推入操作数栈的指令是iconst_0。 



### base

#### bytecode_reader

#### instruction

##### NoOperands

##### Index8Operands
##### Index16Operands
##### BranchOperands

### loads instructions

### store instructions

### stack instructions

- pop and pop2
- dup
- swap

### math instructions

- 
- bool instructions
- iinc instruction
- type change instructions
- compare instructions