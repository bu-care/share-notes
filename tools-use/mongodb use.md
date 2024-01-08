## 1. 创建数据库与用户

1. 在docker执行命令进入MongoDB shell

   ```shell
   root@mongo:/# mongosh
   Current Mongosh Log ID: 63fdcc0f807f2823203532ed
   Connecting to:          mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.6.0
   Using MongoDB:          6.0.2
   Using Mongosh:          1.6.0
   
   For mongosh info see: https://docs.mongodb.com/mongodb-shell/
   
   
   To help improve our products, anonymous usage data is collected and sent to MongoDB periodically (https://www.mongodb.com/legal/privacy-policy).
   You can opt-out by running the disableTelemetry() command.
   
   # 使用 use 进入 admin 数据库
   test> use admin
   switched to db admin
   ## 输入默认用户名密码登录
   admin> db.auth("root", "wireless_dev")
   { ok: 1 }
   ```

2. 创建 admin 的用户

   ```shell
   ## 创建命令
   db.createUser(
   	{
   	  user: "xbu",
   	  pwd: "wireless_dev",
   	  roles: [{ role: "root", db: "admin" }]
   	}
    );
   ## 注意这里role是root，之前使用 roles:[{ role: "userAdminAnyDatabase", db: "admin" }] 创建不成功，认证无法通过
   admin> db.createUser(
   ...     {
   ...       user: "xbu",
   ...       pwd: "wireless_dev",
   ...       roles: [{ role: "root", db: "admin" }]
   ...     }
   ...  );
   { ok: 1 }
   ```

3. 使用新创建的用户认证登录admin

   ```shell
   admin> db.auth("xbu", "wireless_dev")
   { ok: 1 }
   ## 查看现有的用户
   admin> show users
   [
     {
       _id: 'admin.root',
       userId: new UUID("ff030c30-73a0-4b05-8552-671bf18e2ba6"),
       user: 'root',
       db: 'admin',
       roles: [ { role: 'root', db: 'admin' } ],
       mechanisms: [ 'SCRAM-SHA-1', 'SCRAM-SHA-256' ]
     },
     {
       _id: 'admin.xbu',
       userId: new UUID("2207ba32-67a3-40fe-98c9-cddb5e4a9a96"),
       user: 'xbu',
       db: 'admin',
       roles: [ { role: 'root', db: 'admin' } ],
       mechanisms: [ 'SCRAM-SHA-1', 'SCRAM-SHA-256' ]
     }
   ]
   ```

4. 为数据库设置用户

   ```shell
   ## 目前还没有 dev_test 数据库，使用use dev_test创建 dev_test 数据库
   admin> use dev_test
   switched to db dev_test
   ## 此时数据库中显示没有集合
   dev_test> show collections
   
   ## 向数据库中插入一条数据
   dev_test> db.foo.insert({_id:1,name:"test"})
   { acknowledged: true, insertedIds: { '0': 1 } }
   ## 此时数据库中显示有了 foo 集合
   dev_test> show collections
   foo
   
   ## 将xbu 设置为 dev_test 数据库用户，设置权限为 roles: ["readWrite"]
   dev_test>  db.createUser(
   ...     {
   ...       user: "xbu",
   ...       pwd: "wireless_dev",
   ...       roles: ["readWrite"]
   ...     }
   ...  );
   { ok: 1 }
   
   ## 查看创建的用户
   dev_test> show users
   [
     {
       _id: 'dev_test.xbu',
       userId: new UUID("f7acb6b6-a938-4898-8769-f2bc00fefa95"),
       user: 'xbu',
       db: 'dev_test',
       roles: [ { role: 'readWrite', db: 'dev_test' } ],
       mechanisms: [ 'SCRAM-SHA-1', 'SCRAM-SHA-256' ]
     }
   ]
   ```





