### 递归
1. 递归的条件
   1. 一个问题分为几个子问题
   2. 除了数据规模，求解思路一致
   3. 存在递归终止条件
2. 怎么写递归
   1. 写出递推公式
   2. 找到终止条件
   3. ```N 台阶问题：假如有n个台阶，每次走一个或者两个，问有多少种走法？```
3. 一个问题分解为多个子问题，不要试图想清楚整个递归过程，这是一个误区。应该假设B/C/D子问题已经解决了，不要用人脑去分解递归的步骤。
4. 递归的问题
   1. 警惕堆栈溢出
   2. 改为非递归：由机器的栈改为自己维护的栈
   3. 