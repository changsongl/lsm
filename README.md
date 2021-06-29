# lsm

### Introduction
A key/value db implemented by LSM. It is for studying LSM data structure.
I will design the db by the theories, and then compare with other famous
lsm databases. This README is writen before I write any code, because I 
can think through every detail before I start.

### What is LSM?
LSM is Log-Structured MergeTree. It is a data structure, but it's not a common
tree structure like B+ tree. Many popular databases are implemented by LSM.

The core concept of LSM is to use sequential IO to increase writing speed, but it has
**horizontal and vertical layers** in its design, so it decreases read speed.
It is popular for More-Write-And-Less-Read situation.

### LSM architecture
![alt LSM arch](./design/img/arch.jpeg)

### Reference
[LSM Tree-Based engine](https://www.jianshu.com/p/e89cd503c9ae?utm_campaign=hugo)
