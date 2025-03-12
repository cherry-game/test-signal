# test-signal

## 项目简介 | Project Introduction

这个项目主要用于研究和比较VSCode与GoLand在Go程序调试过程中的行为差异。
This project is dedicated to studying and comparing the debugging behavior differences between VSCode and GoLand for Go programs.

## 主要发现 | Key Findings

### 调试终止行为差异 | Debug Termination Behavior Difference

- **VSCode**:
    - 使用Shift+F5(停止按钮)终止调试时，dlv会直接杀死程序进程
    - When terminating debug session using Shift+F5 (stop button), dlv directly kills the program process

- **GoLand**:
    - 使用Stop按钮终止调试时，不会直接杀死程序进程，程序会接收signal，根据业务逻辑处理完成后并优雅地终止进程
    - When terminating debug session using stop button, it does not directly kill the program process; instead, the program receives a signal and gracefully terminates after completing business logic

## 技术细节 | Technical Details

此项目有助于理解不同IDE在Go程序调试实现上的差异，特别是在进程管理和信号处理方面的不同表现。
This project helps understand the differences in Go program debugging implementation between IDEs, especially in terms of process management and signal handling.

## 疑问 | Questions

怎样配置VSCode让它在调试时与GoLand效果一样?
How to configure VSCode to achieve the same debugging behavior as GoLand?