## 软件设计---文件加解密
#### 主要设计思路
##### 参考勒索病毒方式
+ RSA生成主密钥对-加密子密钥
+ AES生成子密钥对-加密文件
##### 执行步骤
1. 首先管理员生成的私钥A和公钥A，算法可以基于RSA
2. 然后生成随机生成密钥Z，算法基于AES
3. 将用户电脑上的文件通过密钥Z加密
4. 将用户电脑上的密钥Z通过公钥A加密
5. 删除密钥Z、公钥A、源数据

##### 开发进度
完成分段读写文件
两个方向：
1. AES分段加密解密功能
2. 多读多写合成最终文件


