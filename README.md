# JiangxiYouthStudy-Auto

江西省2022新版青年大学习自动大学习

在这里首先感谢XYZliang提供的思路， https://github.com/XYZliang/JiangxiYouthStudyMaker

在此基础上用go写了一个。

原理如下：

POST： http://osscache.vol.jxmfkj.com/pub/vol/volClass/join?accessToken=

param/json：{"course":"1","subOrg":null,"nid":"N0013************","cardNo":"姓名+学号"}

course --> 青年大学习的学习期数 --> http://osscache.vol.jxmfkj.com/html/assets/js/course_data.js   (查询最新青年大学习期数)

nid --> 江西省高校学院专业班级对应的团委组织所特有的nid号，以N0013开头.
            -->http://osscache.vol.jxmfkj.com/pub/vol/config/organization?pid=N0013 访问链接即可一步步找到自己对应的组织nid号
            
cardNo --> 填个人信息，姓名+学号


其实就是 每Post一下 青年大学习后台就有你的学习记录.

简单思路如上，如要丰富化在此基础上升级即可.（抛砖引玉嘿嘿嘿！）

同时我提供一个简单example



免责声明
此新系统还在开发当中，此脚本有随时失效的风险。
本程序为免费开源项目，仅供交流学习，遵循GPL v3开源协议，无任何形式的盈利行为。
本程序服务于原系统，旨在让原系统功能更强大。
本程序皆调用官方接口实现，无任何“Hack”行为，无破坏官方接口行为。
本程序仅做数据处理，不拦截、存储、篡改任何用户数据。
严禁使用本程序进行盈利、散播任何违法信息等行为。
本程序不作任何稳定性的承诺，如因使用本程序导致的问题，均与本软件无关。
