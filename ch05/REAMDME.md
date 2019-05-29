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

#### NoOperands
一条指令 什么都不做
#### Index8Operands
#### Index16Operands
#### BranchOperands

### loads instructions 继承 NoOperands
- iload
- fload
- aload
- lload
- xaload

### 存储指令store instructions  继承 NoOperands
- lstore
### 栈指令 stack instructions  继承 NoOperands

- pop and pop2
- dup (复制栈顶数值并将复制值压入栈顶)
    - dup
    > 
     ```
     bottom -> top
    
      [...][c][b][a]
                   \_
                     |
                     V
      [...][c][b][a][a]
     ```
     
    - dup_x1
    ```   
       bottom -> top
       [...][c][b][a]
               __/
              |
              V
       [...][c][a][b][a]
    ```
    
    - dup_x2
    ```
     bottom -> top
     [...][c][b][a]
            _____/
           |
           V
     [...][a][c][b][a]
     ```
    
    - dup2
    ```
     bottom -> top
     [...][c][b][a]____
               \____   |
                    |  |
                    V  V
     [...][c][b][a][b][a]
     ```
    
    - dup2_x1
    ```
     bottom -> top
     [...][c][b][a]
            _/ __/
           |  |
           V  V
     [...][b][a][c][b][a]
     ```
    
    - dup2_x2
    ```
     bottom -> top
     [...][d][c][b][a]
            ____/ __/
           |   __/
           V  V
     [...][b][a][d][c][b][a]
     ```
    
- swap

### math instructions  继承 NoOperands

- 算术指令
    - 加法指令(add)
    - 减法指令(sub)
    - 乘法指令(mul)
    - 除法指令(div)
    - 求余指令(rem)
    - 取反指令(neg)
- 位移指令
    - ISHL // int左位移
    - ISHR  // int算术右位移
    - IUSHR  // int逻辑右位移
    - LSHL  // long左位移
    - LSHR  // long算术右位移
    - LUSHR  // long逻辑右位移 
- 布尔运算指令(bool instructions)
    - and
        - iand
        - land
    - xor
        - IXOR
        - LXOR
    - or
        - IOR
        - LOR
- iinc instruction
    iinc指令给局部变量表中的int变量增加常量值，局部变量表索
    引和常量值都由指令的操作数提供。
### (类型转换指令)type change instructions  继承 NoOperands

- d2x
- i2x
- f2x
- l2x


### 比较指令 compare instructions 

- lcmp
- fcmp
- dcmp
- ifcond
- if_icmp
- if_acmp

### 控制指令 control instructions

- goto

- tableswitch
- loopupswitch

Java语言中的switch-case语句有两种实现方式：如果case值可以
编码成一个索引表，则实现成tableswitch指令；否则实现成
lookupswitch指令。

下面这个Java方法中的switch-case可以编译成
tableswitch指令，代码如下：
```java
int chooseNear(int i) {
    switch (i) {
        case 0: return 0;
        case 1: return 1;
        case 2: return 2;
        default: return -1;
    }
}
```
下面这个Java方法中的switch-case则需要编译成lookupswitch指
令：
```java
int chooseFar(int i) {
    switch (i) {
        case -100: return -1;
        case 0: return 0;
        case 100: return 1;
        default: return -1;
    }
}
```

### 扩展指令

- `wide` 指令
加载类指令、存储类指令、ret指令和iinc指令需要按索引访问
局部变量表，索引以uint8的形式存在字节码中。对于大部分方法来
说，局部变量表大小都不会超过256，所以用一字节来表示索引就
够了。但是如果有方法的局部变量表超过这限制呢？Java虚拟机规
范定义了wide指令来扩展前述指令。

- `ifnull` 和 `ifnonull` 指令

- `goto_w` 指令
goto_w指令和goto指令的唯一区别就是索引从2字节变成了4
字节。

