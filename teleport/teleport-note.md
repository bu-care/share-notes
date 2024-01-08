### Ubuntu 修改本地hosts文件

```shell
sudo vim /etc/hosts
```

将IP 与域名映射写入hosts文件

```
10.103.12.238   proxy.luna.teleport
10.103.12.238   *.proxy.luna.teleport
```

重启网络服务

```shell
sudo /etc/init.d/network-manager restart
```

## 自签名证书

### 1. 生成自签名证书

```shell
> tree /etc/ssl
├── myCA
│   ├── cacert.pem
│   ├── index.txt
│   ├── index.txt.attr
│   ├── index.txt.old
│   ├── newcerts
│   │   └── 01.pem  # $CNRequest's crt file,the same content with myCert's "*.snwl.com.crt"  
│   ├── private
│   │   └── cakey.pem
│   ├── serial
│   └── serial.old
├── myCert
│   ├── *.snwl.com.crt
│   ├── *.snwl.com.csr
│   ├── *.snwl.com.key
│   └── v3.ext
└── openssl.cnf
```

### 2. 安装自签名证书

make system to trust root certificate

#### 1. linux server (ubuntu/debian)

Install the CA Cert as a trusted root CA
\- Move the CA certificate (`ca.pem`) into `/usr/local/share/ca-certificates/ca.crt`.
\- Update the Cert Store with shell command `update-ca-certificates`

```shell
apt install ca-certificates
# Copy your CA to dir /usr/local/share/ca-certificates/
sudo cp foo.crt /usr/local/share/ca-certificates/foo.crt
sudo update-ca-certificates
# To remove/rebuild
sudo update-ca-certificates --fresh
```

there is a example for our self create CA, we must put  **‘/etc/ssl/myCA/cacert.pem’** mentioned above on every node which want to join the cluster （such as k3s-master, k3s-node1...）,

In this example, we use k3s-node1 to trust our CA '[utm2ca.snwl.com](http://utm2ca.snwl.com/)' 

```shell
# copy the cert file(/etc/ssl/myCA/cacert.pem) to k3s-node1 folder /usr/local/share/ca-certificates/ca.crt !!!
 
# before  we trust CA  'utm2ca.snwl.com' 
> wget https://teleport.snwl.com/
--2022-10-25 05:16:02--  https://teleport.snwl.com/
Resolving teleport.snwl.com (teleport.snwl.com)... 10.103.229.104
Connecting to teleport.snwl.com (teleport.snwl.com)|10.103.229.104|:443... connected.
ERROR: cannot verify teleport.snwl.com's certificate, issued by ‘CN=hnyin,OU=utm2,O=sonicwall,L=yangpu,ST=shanghai,C=CN’:
  Unable to locally verify the issuer's authority.
To connect to teleport.snwl.com insecurely, use `--no-check-certificate'.
 
 
> update-ca-certificates
Updating certificates in /etc/ssl/certs...
1 added, 0 removed; done.
Running hooks in /etc/ca-certificates/update.d...
done.
 
# after we trust CA 'utm2ca.snwl.com'  we can not see error
> wget https://teleport.snwl.com/
--2022-10-25 05:21:14--  https://teleport.snwl.com/
Resolving teleport.snwl.com (teleport.snwl.com)... 10.103.229.104
Connecting to teleport.snwl.com (teleport.snwl.com)|10.103.229.104|:443... connected.
HTTP request sent, awaiting response... 302 Found
Location: /web [following]
--2022-10-25 05:21:14--  https://teleport.snwl.com/web
Reusing existing connection to teleport.snwl.com:443.
HTTP request sent, awaiting response... 200 OK
Length: 698 [text/html]
Saving to: ‘index.html’
 
index.html                  100%[=========================================>]     698  --.-KB/s    in 0s
 
2022-10-25 05:21:14 (47.8 MB/s) - ‘index.html’ saved [698/698]
```

#### 2. windows machine 

命令行方式：

```shell
certutil -addstore -f "ROOT" new-root-certificate.crt
certutil -delstore "ROOT" serial-number-hex
```

手动操作：

- 将 CA server(10.103.229.104) /etc/ssl/myCA/cacert.pem 复制到您的 Windows 服务器并将其扩展名重命名为 crt，以便 Windows 可以识别证书文件
- 导入证书点击“安装证书”
- search for cert management 
- verify certificate is imported

#### 3. Mac OS X

```shell
sudo security add-trusted-cert -d -r trustRoot -k /Library/Keychains/System.keychain ~/new-root-certificate.crt
sudo security delete-certificate -c "<name of existing certificate>"
```

### 3.teleport使用自签名证书

1. 将生成的 myCert 目录下的自签名证书复制到teleport存放证书的目录下

```shell
cp ./selfSingedCert/myCert/*.snwl.com.key /etc/teleport/certs/*.snwl.com.key
cp ./selfSingedCert/myCert/*.snwl.com.crt /etc/teleport/certs/*.snwl.com.crt
```

2. 在teleport.yaml配置文件中的 proxy_service 中配置使用自签名证书

   ```yaml
   proxy_service:
     enabled: "yes"
     web_listen_addr: 0.0.0.0:443
     public_addr: teletest.snwl.com:443
     https_keypairs:
       - key_file: /etc/teleport/certs/*.snwl.com.key
         cert_file: /etc/teleport/certs/*.snwl.com.crt
     acme: { }
   ```

   

## 1. 安装 teleport

### 1. 通过Docker-compose 安装

官方的docker-compose文件没办法修改它的默认域名

```yaml
version: '2'
services:
  # This container depends on the config written by the configure container above, so it
  # sleeps for a second on startup to allow the configure container to run first.
  teleport:
    image: public.ecr.aws/gravitational/teleport-lab:10
    container_name: teleport
    entrypoint: /bin/sh
    hostname: luna.teleport
    command: -c "/usr/bin/dumb-init teleport start -d -c /etc/teleport.d/teleport.yaml"
    ports:
      - "3023:3023"
      - "3024:3024"
      - "3025:3025"
      - "443:443"
    volumes:
      - config:/etc/teleport
      - data:/var/lib/teleport
      - certs:/mnt/shared/certs
    networks:
      teleport:
        aliases:
          - proxy.luna.teleport

  # The bootstrap container generates certificates and then immediately exits.
  bootstrap:
    image: public.ecr.aws/gravitational/teleport-lab:10
    container_name: teleport-bootstrap
    entrypoint: /bin/sh
    command: -c "/etc/teleport.d/scripts/generate-certs.sh"
    volumes:
      - config:/etc/teleport
      - data:/var/lib/teleport
      - certs:/mnt/shared/certs
    depends_on:
      - teleport
    networks:
      - teleport

  #
  # openssh is a demo of openssh node
  #
  openssh:
    image: public.ecr.aws/gravitational/teleport-lab:10
    container_name: openssh
    hostname: mars.openssh.teleport
    entrypoint: /bin/sh
    command: -c "/etc/teleport.d/scripts/start-sshd.sh"
    mem_limit: 300m
    volumes:
      - certs:/mnt/shared/certs
    depends_on:
      - bootstrap
    networks:
      teleport:
        aliases:
          - mars.openssh.teleport

  #
  # term is a container with a terminal to try things out
  #
  term:
    image: public.ecr.aws/gravitational/teleport-lab:10
    hostname: term
    container_name: term
    entrypoint: /bin/sh
    command: -c "sleep infinity"
    mem_limit: 300m
    volumes:
      - certs:/mnt/shared/certs
    depends_on:
      - bootstrap
    networks:
      - teleport

volumes:
  certs:
  data:
  config:

networks:
  teleport:

```



```shell
docker-compose -f teleport-lab.yml up -d

docker-compose -f teleport-lab.yml down
```



### 2. 在宿主机安装

#### 1. 博客上的步骤

按照 https://cloud.tencent.com/developer/article/1453751 上面的步骤进行搭建公网 Master，但它是v2.3.5版本，GitHub上面找不到，所以现在了新的10.2.6版本。

为了让服务后台运行创建一个 systemd service 配置文件，注意这里的路径是/etc/teleport/teleport.yaml，而官方默认teleport.yaml的路径是/etc/teleport.yaml

```shell
cat > /etc/systemd/system/teleport.service <<EOF
[Unit]
Description=Teleport SSH Service
After=network.target

[Service]
Type=simple
Restart=always
ExecStart=/usr/local/bin/teleport start -c /etc/teleport.yaml

[Install]
WantedBy=multi-user.target
EOF
```

官方文档执行后生成的默认的teleport.yaml

```yaml
#
# A Sample Teleport configuration file.
#
# Things to update:
#  1. license.pem: You only need a license from https://dashboard.goteleport.com
#     if you are an Enterprise customer.
#
version: v2
teleport:
  nodename: xbu-desktop-1250
  data_dir: /var/lib/teleport
  advertise_ip: 10.103.12.50
  log:
    output: stderr
    severity: INFO
    format:
      output: text
  ca_pin: ""
  diag_addr: ""
auth_service:
  enabled: "yes"
  listen_addr: 0.0.0.0:3025
  cluster_name: 10.103.12.50
  proxy_listener_mode: multiplex
ssh_service:
  enabled: "yes"
  listen_addr: 0.0.0.0:3022
  commands:
  - name: hostname
    command: [hostname]
    period: 1m0s
proxy_service:
  enabled: "yes"
  web_listen_addr: 0.0.0.0:443
  public_addr: 10.103.12.50:443
```

对博客上的teleport.yaml文档做了修改

```yaml
teleport:
    nodename: xbu-desktop-1250
    data_dir: /var/lib/teleport
    auth_token: jYektagNTmhjv9Dh
    advertise_ip: 10.103.12.50
    auth_servers:
        - 0.0.0.0:3025
        - 0.0.0.0:3025
    connection_limits:
        max_connections: 1000
        max_users: 250
    log:
        output: stdout
        severity: INFO
        format:
        	output: text
    ca_pin: ""
    diag_addr: ""
    storage:
        type: bolt

# This section configures the 'auth service':
auth_service:
    # Turns 'auth' role on. Default is 'yes'
    enabled: yes
    listen_addr: 0.0.0.0:3025
    tokens:
        - "proxy,node:jYektagNTmhjv9Dh"
        - "auth:jYektagNTmhjv9Dh"
    cluster_name: 10.103.12.50
    proxy_listener_mode: multiplex

# This section configures the 'node service':
ssh_service:
    # Turns 'ssh' role on. Default is 'yes'
    enabled: yes
    listen_addr: 0.0.0.0:3022
    labels:
        role: master
    commands:
    - name: hostname             # this command will add a label like 'arch=x86_64' to a node
      command: [hostname]
      period: 1m0s
    # enables reading ~/.tsh/environment before creating a session. by default
    # set to false, can be set true here or as a command line flag.
    permit_user_env: false

# This section configures the 'proxy servie'
proxy_service:
    # Turns 'proxy' role on. Default is 'yes'
    enabled: yes
    listen_addr: 0.0.0.0:3023
    tunnel_listen_addr: 0.0.0.0:3024
    web_listen_addr: 0.0.0.0:443
    public_addr: 10.103.12.50:443
```



在写完 `/etc/teleport.yaml` 的文件之后，直接执行

```shell
systemctl enable teleport
systemctl start teleport
```

无法起来teleport，起来后就失败，我在官网上（https://goteleport.com/docs/deploy-a-cluster/open-source/）的步骤中看到的步骤，

```shell
# 删除之前的服务
sudo rm -f /etc/systemd/system/teleport.service

# 重新加载 systemctl daemon-reload
# Warning: The unit file, source configuration file or drop-ins of teleport.service changed on disk. Run 'systemctl daemon-reload' to reload units.
xbu@xbu-desktop-1250:/etc$ sudo systemctl daemon-reload

# 再次启动 teleport
sudo systemctl start teleport
```

#### 2. 官方步骤

生成yaml配置文件

```shell
DOMAIN=tele.xbu
EMAIL=xbu@sonicwall.com
teleport configure --acme --acme-email=${EMAIL?} --cluster-name=${DOMAIN?} | \
 sudo tee /etc/teleport.yaml > /dev/null
```

在您的Linux机器上，运行以下命令来启动  `teleport` 守护程序（这取决于您之前安装teleport的方式）。

- Package manager RPM/DEB

  ```shell
  sudo systemctl start teleport
  ```

- Source or custom install

在teleport“v10.1.4”及更高版本上使用 `teleport install systemd` 命令创建一个systemd单元文件。作为命令的一部分，输入要用于配置的单位文件路径。对于大多数使用情况，我们建议使用路径 `/etc/systemd/system/teleport.service`：

先执行这一步，然后在执行下面的启动，可以成功。

```shell
sudo teleport install systemd -o /etc/systemd/system/teleport.service
```

Enable and start the new `teleport` daemon:

```shell
sudo systemctl enable teleport && sudo systemctl start teleport
```



### 3. 通过docker安装



```shell
# 为Teleport创建本地配置和数据目录，这些目录将装入容器。
mkdir -p ~/teleport/config ~/teleport/data

# 生成一个示例Teleport配置并将其写入本地配置目录。这个容器将写入配置并立即退出——这是预期的。
docker run --hostname tele-host --rm \
  --entrypoint=/bin/sh \
  -v ~/teleport/config:/etc/teleport.d \
  public.ecr.aws/gravitational/teleport:10.3.2 -c "teleport configure > /etc/teleport.d/teleport.yaml"

# 使用挂载的配置和数据目录以及所有端口启动Teleport
docker run -d --hostname tele-host --name teleport \
  -v ~/teleport/config:/etc/teleport \
  -v ~/teleport/data:/var/lib/teleport \
  -p 3023:3023 -p 3025:3025 -p 3080:3080 \
  public.ecr.aws/gravitational/teleport:10.3.2
```

### 4. teleport添加用户

```shell
tctl users add xbu --roles=editor,access --logins=root,ubuntu,ec2-user
```



## 2. 添加节点

### 1. 添加单点登录 sso

#### 1. 创建 GitHub OAuth 应用程序

创建并注册 GitHub OAuth 应用程序。执行此操作时，请确保您的 OAuth 应用程序的“身份验证回调 URL”如下：

```
https://PROXY_ADDRESS/v1/webapi/github/
```

#### 2. 创建 GitHub 身份验证连接器

通过创建一个名为 github.yaml 的文件来定义 GitHub 身份验证连接器

使用个人账户（buxuehu）创建的 OAuth APP

```yaml
kind: github
version: v3
metadata:
  # Connector name that will be used with `tsh --auth=github login`
  name: github
spec:
  # Client ID of your GitHub OAuth App
  client_id: 7db14713d38dff535c8e
  # Client secret of your GitHub OAuth App
  client_secret: 0c3696c9dc920c8f0bef6a78e7e0f7e2fe2a264c
  # Connector display name that will be shown on the Web UI login screen
  display: GitHub
  # Callback URL that will be called after successful authentication
  redirect_url: https://xbu1250.tele:3080/v1/webapi/github/callback
  # Mapping of org/team memberships onto allowed roles
  teams_to_roles:
    - organization: buxuehu # GitHub organization name
      team: go-dev # GitHub team name within that organization
      # Maps octocats/admins to the "access" Teleport role
      roles:
        - access

```

使用组织账户（buxuehu）创建的 OAuth APP

```yaml
kind: github
version: v3
metadata:
  # Connector name that will be used with `tsh --auth=github login`
  name: github
spec:
  # Client ID of your GitHub OAuth App
  client_id: feba7bb753c28de130fb
  # Client secret of your GitHub OAuth App
  client_secret: 868d5efe378f22def7620c1a741637474f4e1d2b
  # Connector display name that will be shown on the Web UI login screen
  display: GitHub
  # Callback URL that will be called after successful authentication
  redirect_url: https://xbu1250.tele:3080/v1/webapi/github/callback
  # Mapping of org/team memberships onto allowed roles
  teams_to_roles:
    - organization: bu-family # GitHub organization name
      team: go-dev # GitHub team name within that organization
      # Maps octocats/admins to the "access" Teleport role
      roles:
        - access
```

使用 tctl 创建连接器

```shell
tctl create github.yaml
```

#### 3. 配置身份验证

获取现有的 cluster_auth_preference 资源：

```shell
tctl get cap > cap.yaml
```

如果没有定义 cluster_auth_preference，则 cap.yaml 将为空白。确保 cap.yaml 包含以下内容：

```yaml
kind: cluster_auth_preference
metadata:
  name: cluster-auth-preference
spec:
  type: github
  webauthn:
    rp_id: 'xbu1250.tele:3080'
version: v2

```

对于 rp_id，使用 Teleport Proxy Service 的公共地址。

Create the resource:

```shell
tctl create -f cap.yaml
#cluster auth preference has been updated
```

您还可以编辑 Teleport 配置文件以包含以下内容：

```yaml
# Snippet from /etc/teleport.yaml
auth_service:
  authentication:
    type: github

```

### 2. 将节点添加到集群

#### 1. 在集群机器上生成token

```shell
sudo tctl tokens add --type=node | grep -oP '(?<=token:\s).*' > token.file

# 生成APP的token
tctl tokens add \
    --type=app \
    --app-name=grafana \
    --app-uri=http://10.103.230.16:3000 \
    --insecure
    
tctl tokens add \
    --type=app \
    --app-name=fw23022 \
    --app-uri=https://10.103.230.22:443 \
    --insecure
    
tctl tokens add \
    --type=app \
    --app-name=test-report \
    --app-uri=http://10.103.49.247:8081
```



在节点机器执行 `teleport start` 将节点或应用通过token加入集群（节点机器不用 `tsh login` ，直接 `teleport start` 就行了）

```shell
sudo teleport start \
   --roles=node \
   --token=098a643fd02b96e8e51a36afcee240c4 \
   --auth-server=teletest.snwl.com:443  --insecure

teleport start \
   --roles=node \
   --nodename=ip-10-103-230-19 \
   --token=098a643fd02b96e8e51a36afcee240c4 \
   --auth-server=teletest.snwl.com:443 
   
sudo teleport start \
   --roles=node \
   --token=/teleport/token/teleport12-238 \
   --auth-server=proxy.luna.teleport:443  --insecure
```



#### 2. 节点上将APP加入集群

```shell
teleport app start \
   --token=dd3c09e59ab5cfd52a8490834f53d6eb \
   --ca-pin=sha256:fff4c6b6ffc63d15a1d1144ca623a4caab0cb94ae3b1c79327d993e116ed4249 \
   --auth-server=proxy.luna.teleport:443 \
   --name=grafana                        `# Change "grafana" to the name of your application.` \
   --uri=http://10.103.230.16:3000
   
sudo teleport app start \
  --name=grafana \
  --token=/tmp/token \
  --uri=http://localhost:3000 \
  --auth-server=https://proxy.luna.teleport:443  --insecure
  
sudo teleport app start \
  --name=fw23022 \
  --token=/tmp/xbu1250/fw23022-token \
  --uri=https://10.103.230.22:443 \
  --auth-server=https://proxy.luna.teleport:443 \ 
  --insecure
  
teleport app start \
  --name=demo-app \
  --token=/tele-data/tokens/demo-app \
  --uri=http://192.168.100.2:9000 \
  --auth-server=https://xbu1250.tele:3080 \
  --insecure
```

通过start_app.yaml文件启动，注意应用程序的public_addr“10.103.230.22”不能是IP地址，teleport应用程序访问使用DNS名称进行路由

```shell
# 生成一个node和app类型的token
tctl tokens add --type=node,app
```



```yaml
version: v3
teleport:
  nodename: ip-10-103-230-22
  data_dir: /var/lib/teleport
  ca_pin: "sha256:87d32aa01fb07f33982f2f569a765dcfd949ac00f0e606747de7182d3db935d2"
  auth_token: "node and app type token"
  proxy_server: teletest.snwl.com:443
  log:
    output: stderr
    severity: INFO
    format:
      output: text
  diag_addr: ""

auth_service:
  enabled: "no"
ssh_service:
  enabled: "yes"
  commands:
    # this command will add a label 'arch=x86_64' to a node
    - name: hostname # arch
      command: ['/bin/uname', '-p']
      period: 1h0m0s
proxy_service:
  enabled: "no"
  https_keypairs: []
  acme: {}
  
app_service:
    enabled: yes
    # Teleport provides a small debug app that can be used to make sure application
    # access is working correctly. It'll output JWTs so it can be useful
    # when extending your application.
    debug_app: true
    apps:
    - name: "fw23022"
      # URI and port the application is available at.
      uri: "https://10.103.230.22"
      # Optional application public address to override.
      public_addr: "fw23022.teletest.snwl.com"
      insecure_skip_verify: true
      rewrite:
        # Rewrite the "Location" header on redirect responses replacing the
        # host with the public address of this application.
        redirect:
        - "10.103.230.22"
        - "fw23022.teletest.snwl.com"
      # Optional static labels to assign to the app. Used in RBAC.
      labels:
        env: "prod"
      # Optional dynamic labels to assign to the app. Used in RBAC.
      commands:
      - name: "os"
        command: ["/usr/bin/uname"]
        period: "5s"
        
```



```yaml
version: v3
teleport:
  nodename: ip-10-103-229-71
  data_dir: /var/lib/teleport
  ca_pin: "sha256:87d32aa01fb07f33982f2f569a765dcfd949ac00f0e606747de7182d3db935d2"
  auth_token: "node and app type token"
  proxy_server: teletest.snwl.com:443
  log:
    output: stderr
    severity: INFO
    format:
      output: text
  diag_addr: ""

auth_service:
  enabled: "no"
ssh_service:
  enabled: "yes"
proxy_service:
  enabled: "no"
  https_keypairs: []
  acme: {}
  
app_service:
    enabled: yes
    debug_app: true
    apps:
    - name: "fw22971"
      uri: "https://10.103.229.71"
      public_addr: "fw22971.teletest.snwl.com"
      insecure_skip_verify: true
      rewrite:
        redirect:
        - "10.103.229.71"
        - "fw22971.teletest.snwl.com"
      labels:
        env: "prod"
      commands:
      - name: "os"
        command: ["/usr/bin/uname"]
        period: "5s"
```



```shell
teleport start -c /etc/teleport/teleport.yaml
sudo teleport start --config=./start_app.yaml
```



### 3. 添加数据库

#### 1. 创建数据库

通过docker compose创建MongoDB数据库

```yaml
version: v3
teleport:
  nodename: ip-10-103-230-19
  data_dir: /var/lib/teleport
  ca_pin: "sha256:87d32aa01fb07f33982f2f569a765dcfd949ac00f0e606747de7182d3db935d2"
  auth_token: b9c4d4f54aaf673baf3e89f05170ba5f
  proxy_server: teletest.snwl.com:443
  log:
    output: stderr
    severity: INFO
    format:
      output: text
  diag_addr: ""

auth_service:
  enabled: "no"
ssh_service:
  enabled: "yes"
proxy_service:
  enabled: "no"
  https_keypairs: []
  acme: {}

windows_desktop_service:
  enabled: yes
  ldap:
    addr:     '10.103.230.53:636'
    domain:   'teleport.dev'
    username: 'TELEPORT\svc-teleport'
    server_name: 'WIN-DK5CL2SUTL9.teleport.dev'
    insecure_skip_verify: false
    ldap_ca_cert: |
        -----BEGIN CERTIFICATE-----
        MIIDhTCCAm2gAwIBAgIQMkzkplo89aZMoKhy7XJ85TANBgkqhkiG9w0BAQsFADBV
        MRMwEQYKCZImiZPyLGQBGRYDZGV2MRgwFgYKCZImiZPyLGQBGRYIdGVsZXBvcnQx
        JDAiBgNVBAMTG3RlbGVwb3J0LVdJTi1ESzVDTDJTVVRMOS1DQTAeFw0yMjExMjUw
        MjIzNThaFw0yNzExMjUwMjMzNTNaMFUxEzARBgoJkiaJk/IsZAEZFgNkZXYxGDAW
        BgoJkiaJk/IsZAEZFgh0ZWxlcG9ydDEkMCIGA1UEAxMbdGVsZXBvcnQtV0lOLURL
        NUNMMlNVVEw5LUNBMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0iSH
        hHPRIG9AV/CeGCiWuA/oTJCXrVwsUMKehVfLMYNHhdmriO2g/1pf5l/63jO2+Wno
        vEAviMKXnuTai5vRDZMXHXhmyOKUPM0wUdbX8P839YtaVwG9o55N+QEp9oNVJZ2n
        Os+iP5HH8AJ2Ygh6oprGEf1iqguc9b+Q/OUN9kLAexDwSMiKxBKk8RIstwc8ECdf
        0NcFIIg8vnGn8Xae0k/lOpss8GSHk+VbNDC4sAOnyrqzlMI0Fjz9EG1QgwzVTVNu
        4XyKLYU2kCerSdvzH1XcEZvmHKkRfoOgxJ/N3MWXRzWlWnCa6ttOvZNiPSOTuNHZ
        0QwRIAa+kYnRGbQRCQIDAQABo1EwTzALBgNVHQ8EBAMCAYYwDwYDVR0TAQH/BAUw
        AwEB/zAdBgNVHQ4EFgQUzSmJ85OpCDbjcV+UnXF4m0qwEdEwEAYJKwYBBAGCNxUB
        BAMCAQAwDQYJKoZIhvcNAQELBQADggEBABeDJ0pkQRmz1V1MEejeChsTG1DFBnHo
        O2EERGSsJQ9HJnxvq0wwkMXM+xytdpDTInuTqczI/5c7C/kitXGUunfCYgSo4Pmw
        nxzOS4Hd47ORnymgxX/Sie18weZorSLJYvhYRXPYoppGeW+f2TVyGROC5MjXeAiU
        A5puDZ6Y3bsWu6qTuOj7EZCbhh0n15oJfOa754mp0Wj6oBKlFU7Ti/GL/CfUNmZK
        Qo3MlMDVCVsLTWFEsjMe6MpFkCCggk1+8g+XeAGjA5tu/coiCRlnavQN6+h8PoX+
        xwvn4AIU+smWgUXCJ/DlQfXx0NP0AcMcvvT1fM7Lpk8o0R3+MOHkHBs=
        -----END CERTIFICATE-----

  discovery:
    base_dn: '*'
  labels:
    teleport.internal/resource-id: c654dd8c-6d75-426b-bd31-033f33ab21ed
    
db_service:
  enabled: "yes"
  resources:
  - labels:
      "*": "*"

  databases:
  - name: "mysql01"
    description: "mysql test"
    protocol: "mysql"

    # Database connection endpoint. Must be reachable from Database Service.
    uri: "10.103.230.16:3306"
    tls:
      mode: insecure # verify-full
      #server_name: db.example.com
      #ca_cert_file: /path/to/pem
    mysql:
      server_version: 8.0.28
    dynamic_labels:
    - name: "hostname"
      command: ["hostname"]
      period: 1m0s
  - name: "mongo01"
    description: "mongodb"
    protocol: "mongodb"
    uri: "mongodb://10.103.230.16:27017"
    tls:
      mode: insecure # verify-full
      #server_name: db.example.com
      #ca_cert_file: /path/to/pem
    static_labels:
      env: "dev"
```



```yaml
# Use root/example as user/password credentials
version: '3.1'
services:
  mongo:
    image: mongo
    restart: always
    hostname: mongo
    ports:
      - 27017:27017
    environment:
      MONGO_INITDB_ROOT_USERNAME: root
      MONGO_INITDB_ROOT_PASSWORD: wireless_dev

  mongo-express:
    image: mongo-express
    restart: always
    hostname: mongo-express
    ports:
      - 8081:8081
    environment:
      ME_CONFIG_MONGODB_ADMINUSERNAME: root
      ME_CONFIG_MONGODB_ADMINPASSWORD: wireless_dev
      ME_CONFIG_MONGODB_URL: mongodb://root:wireless_dev@mongo:27017/
```

创建MySQL数据库

```yaml
# Use root/example as user/password credentials
version: '3.1'
services:
  db:
    image: mysql
    command: --default-authentication-plugin=mysql_native_password
    restart: always
    ports:
      - 3306:3306
    environment:
      MYSQL_ROOT_PASSWORD: wireless_dev

  adminer:
    image: adminer
    restart: always
    ports:
      - 8080:8080
```



```shell
# 启动
docker-compose -f teleport-lab.yml up -d
docker-compose -f teleport-lab.yml down
```

#### 2. 创建用户

数据库需要创建授权所用的用户

```shell
db.getSiblingDB("$external").runCommand(
  {
    createUser: "CN=xbu",
    roles: [
      { role: "readWriteAnyDatabase", db: "admin" }
    ]
  }
)
```

设置双向TLS（mTLS）

该命令将创建两个文件：带有 Teleport 证书颁发机构的 mongo.cas 和带有生成的证书和密钥对的 mongo.crt。您将需要这些文件在 MongoDB 服务器上启用双向 TLS。

```shell
tctl auth sign --format=mongodb --host=10.103.230.16:27017 --out=mongo --ttl=2190h
```

使用生成的密钥在 mongod.conf 配置文件中启用双向 TLS 并重新启动数据库：

```yaml
# 3.6-4.2版本 
net:
  ssl:
    mode: requireSSL
    PEMKeyFile: /etc/certs/mongo.crt
    CAFile: /etc/certs/mongo.cas
    
# 4.2+版本    
net:
  tls:
    mode: requireTLS
    certificateKeyFile: /etc/certs/mongo.crt
    CAFile: /etc/certs/mongo.cas

export PATH=$PATH:/Users/wirelessdev/programs/mongosh-1.6.0-darwin-x64/bin
```

创建token

```shell
tctl tokens add --type=db



```



```shell
teleport db start \
   --token=c77e2c00131ddf1714171568d835df50 \
   --ca-pin=sha256:b49d5a469d1c1624b6b9ea44b4b010092da692ca753ec758b0050dd364645843 \
   --auth-server=xbu1250.tele:3080 \
   --name=mongo01 \
   --protocol=mongodb \
   --uri=192.168.200.7:27017
   
teleport db start \
  --token=/tmp/teleport-token/mongo-token \
  --name=mongo \
  --auth-server=xbu1250.tele:3080 \
  --protocol=mongodb \
  --uri=10.103.230.16:27017 \
  --insecure
```



```shell
teleport db configure create \
   --token=3862a6294bd4f26959a724e689e115f8 \
   --ca-pin=sha256:16521b38e4735e25c7299e385f9ca7572e8e58263fb5b08d82ea8ed385d28012 \
   --proxy=xbu1250.tele:3080 \
   --name=mongo \
   --protocol=mongodb \
   --uri=10.103.230.16:27017 \
   --output file:///etc/teleport.yaml
   
   
teleport start -c /etc/teleport.yaml
```



### 4. 添加api

1. 首先是创建user

   在tele-host的主机上创建用来连接到proxy上的user，

   ```shell
   tctl users add api-admin --roles=editor
   ```

2. 生成  client credentials

   在要执行脚本的主机上执行login user，就会在 ~/.tsh 目录下生成登录的user的相关认证信息，可以用来给脚本中读取认证

   ```shell
   tsh login --user=api-admin --proxy=xbu1xbu1250.:3080 --insecure
   ```

3. 执行go脚本

   ```go
   package main
   
   import (
   	"context"
   	"log"
   
   	"github.com/gravitational/teleport/api/client"
   )
   
   func main() {
   	ctx := context.Background()
   
   	clt, err := client.New(ctx, client.Config{
   		Addrs: []string{
   			// Teleport Cloud customers should use <tenantname>.teleport.sh
   			"xbu1250.tele:443",
   			"xbu1250.tele:3025",
   			"xbu1250.tele:3024",
   			"xbu1250.tele:3080",
   		},
   		// Multiple Credentials can be provided to attempt to authenticate
   		// the client. At least one Credentials object must be provided.
   		Credentials: []client.Credentials{
         //第一个参数 dir 是配置文件目录。它将默认为“~/.tsh”，
         //第二个参数是 name 是配置文件名称。它将默认为当前活动的 tsh 配置文件：xbu1250.tele.yaml
   			client.LoadProfile("", ""),
         //client.LoadProfile("/Users/wirelessdev/.tsh", "xbu1250.tele")
   		},
   	})
   
   	if err != nil {
   		log.Fatalf("failed to create client: %v", err)
   	}
   
   	defer clt.Close()
   	resp, err := clt.Ping(ctx)
   	if err != nil {
   		log.Fatalf("failed to ping server: %v", err)
   	}
   
   	log.Printf("Example success!")
   	log.Printf("Example server response: %s", resp)
   	log.Printf("Server version: %s", resp.ServerVersion)
   }
   ```





### 5. 添加 desktop

#### 1. 添加win server

1. vcenter中创建win2016server虚拟机(密码：wir....._dev01)，win2019的虚拟机有问题，无法安装DNS server

首先关闭ipv6，在CMD下分别执行如下三条命令关闭Tunnel adapter，再去配置静态IP，然后再去搭建AD域服务，不然可能会导致teleport识别DHCP分配的默认(192.*)IP，而不去连接自己配置的静态IP。

```shell
netsh interface teredo set state disable
netsh interface 6to4 set state disable
netsh interface isatap set state disable
```

2. 服务端域服务搭建：首先需要安装 AD server、DNS server，然后再执行teleport中添加desktop步骤中的脚本（install-ad-cs.ps1）安装ADCS

```shell
$ErrorActionPreference = "Stop"

Add-WindowsFeature Adcs-Cert-Authority -IncludeManagementTools
Install-AdcsCertificationAuthority -CAType EnterpriseRootCA -Force
Restart-Computer -Force
```

3. 打开teleport中的配置链接（如https://teletest.snwl.com/v1/webapi/scripts/desktop-access/configure/d8a498ae69498606adbcd8b0afe0b821/configure-ad.ps1），执行teleport中的脚本（configure-ad.ps1）配置AD，生成运行 Windows 桌面服务实例的yaml文件

```yaml
version: v3
teleport:
  auth_token: d8a498ae69498606adbcd8b0afe0b821
  proxy_server: teletest.snwl.com:443

auth_service:
  enabled: no
ssh_service:
  enabled: no
proxy_service:
  enabled: no

windows_desktop_service:
  enabled: yes
  ldap:
    addr:     '10.103.230.53:636'
    domain:   'teleport.dev'
    username: 'TELEPORT\svc-teleport'
    server_name: 'WIN-DK5CL2SUTL9.teleport.dev'
    insecure_skip_verify: false
    ldap_ca_cert: |
        -----BEGIN CERTIFICATE-----
        MIIDhTCCAm2gAwIBAgIQMkzkplo89aZMoKhy7XJ85TANBgkqhkiG9w0BAQsFADBV
        MRMwEQYKCZImiZPyLGQBGRYDZGV2MRgwFgYKCZImiZPyLGQBGRYIdGVsZXBvcnQx
        JDAiBgNVBAMTG3RlbGVwb3J0LVdJTi1ESzVDTDJTVVRMOS1DQTAeFw0yMjExMjUw
        MjIzNThaFw0yNzExMjUwMjMzNTNaMFUxEzARBgoJkiaJk/IsZAEZFgNkZXYxGDAW
        BgoJkiaJk/IsZAEZFgh0ZWxlcG9ydDEkMCIGA1UEAxMbdGVsZXBvcnQtV0lOLURL
        NUNMMlNVVEw5LUNBMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0iSH
        hHPRIG9AV/CeGCiWuA/oTJCXrVwsUMKehVfLMYNHhdmriO2g/1pf5l/63jO2+Wno
        vEAviMKXnuTai5vRDZMXHXhmyOKUPM0wUdbX8P839YtaVwG9o55N+QEp9oNVJZ2n
        Os+iP5HH8AJ2Ygh6oprGEf1iqguc9b+Q/OUN9kLAexDwSMiKxBKk8RIstwc8ECdf
        0NcFIIg8vnGn8Xae0k/lOpss8GSHk+VbNDC4sAOnyrqzlMI0Fjz9EG1QgwzVTVNu
        4XyKLYU2kCerSdvzH1XcEZvmHKkRfoOgxJ/N3MWXRzWlWnCa6ttOvZNiPSOTuNHZ
        0QwRIAa+kYnRGbQRCQIDAQABo1EwTzALBgNVHQ8EBAMCAYYwDwYDVR0TAQH/BAUw
        AwEB/zAdBgNVHQ4EFgQUzSmJ85OpCDbjcV+UnXF4m0qwEdEwEAYJKwYBBAGCNxUB
        BAMCAQAwDQYJKoZIhvcNAQELBQADggEBABeDJ0pkQRmz1V1MEejeChsTG1DFBnHo
        O2EERGSsJQ9HJnxvq0wwkMXM+xytdpDTInuTqczI/5c7C/kitXGUunfCYgSo4Pmw
        nxzOS4Hd47ORnymgxX/Sie18weZorSLJYvhYRXPYoppGeW+f2TVyGROC5MjXeAiU
        A5puDZ6Y3bsWu6qTuOj7EZCbhh0n15oJfOa754mp0Wj6oBKlFU7Ti/GL/CfUNmZK
        Qo3MlMDVCVsLTWFEsjMe6MpFkCCggk1+8g+XeAGjA5tu/coiCRlnavQN6+h8PoX+
        xwvn4AIU+smWgUXCJ/DlQfXx0NP0AcMcvvT1fM7Lpk8o0R3+MOHkHBs=
        -----END CERTIFICATE-----

  discovery:
    base_dn: '*'
  labels:
    teleport.internal/resource-id: c654dd8c-6d75-426b-bd31-033f33ab21ed
```

4. 设置密码永不过期

   在 Administrator Directory Users and Computers >Users >Administrator 的属性的账户中将Administrator设置为密码永不过期

#### 2. 添加win10 desktop

youtube视频地址：https://www.youtube.com/watch?v=YvMqgcq0MTQ

1. 首先安装Active Directory Certificate Services（AD证书服务）

   依次点击：设置》应用》可选功能》添加功能》RSAT:Active Directory证书服务工具，进行安装



```shell
certutil –dspublish –f C:\Users\Administrator\user-ca.cer RootCA

certutil –dspublish –f C:\Users\Administrator\user-ca.cer NTAuthCA
certutil -pulse
```



2. 在teleport server中生成可以用于node和desktop的token

   ```shell
   root@tele-host:~# tctl tokens add --type=node,windowsdesktop
   The invite token: f8ec0df2587acfab121bd4f086b953d9
   This token will expire in 60 minutes.
   
   Run this on the new node to join the cluster:
   
   > teleport start \
      --roles=node,windowsdesktop \
      --token=f8ec0df2587acfab121bd4f086b953d9 \
      --ca-pin=sha256:87d32aa01fb07f33982f2f569a765dcfd949ac00f0e606747de7182d3db935d2 \
      --auth-server=172.25.0.2:3025
   ```
   
   将加入令牌复制到您将运行 Windows 桌面服务的实例上的一个文件中，然后使用以下配置：
   
   
   
   
   
   
   
   更新 /etc/teleport.yaml 后启动 Teleport：
   
   ```shell
   teleport configure | tee /etc/teleport.yaml > /dev/null
   
   
   teleport start -c /etc/teleport/teleport.yaml
   nohup teleport start -c /etc/teleport/teleport.yaml > ./nohup.log 2>&1 &
   nohup ./start_node.sh > ./start_node.log 2>&1 &
   
   # 查看teleport进程
   ps -aux |grep teleport
   ```
   
   

 错误



```
$ certutil -viewstore "ldap:///CN=NTAuthCertificates,CN=Public Key Services,CN=Services,CN=Configuration,DC=teletest,DC=com?caCertificate"



sudo teleport start \
   --roles=node \
   --token=/home/sonicwall/teleport/tokens/token.file \
   --auth-server=xbu1250.tele:3080
```





## 3. 通过一个yaml文件添加

```yaml
version: v2
teleport:
  nodename: k3s-master
  data_dir: /var/lib/teleport
  join_params:
    #token_name: 3c7d91b298cf29300f41f06e5026b52d 
    token_name: 17c1591ce0112f6b86ef1e2e8c9a01df
    method: token
  auth_servers:
  - teleport.snwl.com:443
  log:
    output: stderr
    severity: INFO
    format:
      output: text
  ca_pin: "sha256:81510e1fe940b84a879ad104f869592510bc8cbb70ac1ee74841f682c3cb5464"
  diag_addr: ""
auth_service:
  enabled: "no"
ssh_service:
  enabled: "yes"
proxy_service:
  enabled: "no"
  https_keypairs: []
  acme: {}
app_service:
  enabled: "yes"
  debug_app: false
  apps:
  - name: nginx
    uri: http://localhost:8000/
    public_addr: ""
    insecure_skip_verify: false

  - name: anothernginx
    uri: http://10.103.229.106:9000/
    public_addr: ""
    insecure_skip_verify: false

  - name: grafana
    uri: http://10.103.229.107:3000/
    public_addr: ""
    insecure_skip_verify: false
    #teleport:
    #  auth_token: 8ea133326394239c3b9df8035f2a197a
    # auth_servers: [ teleport.sonicwall.com:443 ]

windows_desktop_service:
  enabled: yes
  ldap:
    addr:     '10.103.229.108:636'
    domain:   'sonictest.com'
    username: 'SONICTEST\svc-teleport'
    server_name: 'WIN-N3MMU2EG4G8.sonictest.com'
    insecure_skip_verify: false
    ldap_ca_cert: |
        -----BEGIN CERTIFICATE-----
        MIIDiTCCAnGgAwIBAgIQG6/0DSLmKq5LHmVcBmnY3zANBgkqhkiG9w0BAQsFADBX
        MRMwEQYKCZImiZPyLGQBGRYDY29tMRkwFwYKCZImiZPyLGQBGRYJc29uaWN0ZXN0
        MSUwIwYDVQQDExxzb25pY3Rlc3QtV0lOLU4zTU1VMkVHNEc4LUNBMB4XDTIyMTAy
        NTA4NDYyOVoXDTI3MTAyNTA4NTYyOFowVzETMBEGCgmSJomT8ixkARkWA2NvbTEZ
        MBcGCgmSJomT8ixkARkWCXNvbmljdGVzdDElMCMGA1UEAxMcc29uaWN0ZXN0LVdJ
        Ti1OM01NVTJFRzRHOC1DQTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB
        ANs23VVDH8zs4rSWuA9Z/0iQTdUMrDLz9jJz4ZwRXwet7YnN/C+kO8F1EztQowdF
        nm0waHNtG3tZOarOVkZ9GftQBIlgNia6uKN717NF7zqOE/xGxt06q+qVP9BDd5nS
        A85Sb5MBFyE1k6W5izsYoyIB61LB1tEiK5wPLsMICTJXGbfEI709fsZIB7AuBwFp
        4XvjDtB4IhuCLeiqxQiYWz/RVxlUkjRQ81tQVJw7GIHis8O0urkV5fndIHFZavz2
        VKIyQH/HWssWAfHBH5NX7eDvj7uA3t5VNjbKOtj+Wu2fG9Nm6sHT4ZpEvWjqHOcB
        1IbKf8A/jKXY2t+SXuPVYn0CAwEAAaNRME8wCwYDVR0PBAQDAgGGMA8GA1UdEwEB
        /wQFMAMBAf8wHQYDVR0OBBYEFI2C+1fk0F9MzlXy2Xm8YxtDArUQMBAGCSsGAQQB
        gjcVAQQDAgEAMA0GCSqGSIb3DQEBCwUAA4IBAQBn5ecT9x6AveYqtJ2ns+di3lJd
        12Z55n2peSdO9zPAsYWNO5s21VwkuMPO7/zMuDDz/IZhLSOpbaDWXbCorfsMKhS3
        ZKQlJjJL7MsgzukeRdlR1RC2f+beFhqDFjp0AMOwoP7EGl5JIQy+c9VehuTbKl9B
        hJXKErIwSFPQ5zZaKh+PANn04MAl07G7bL2WjHwNaCQjI4m1qNw8wmptRbOT3rfu
        lwLtrRUNcv8V0s84N95IX5MPzvRYj8roA7ExtmSxpI8DSuPK66Gk8bIVv81IR/oE
        RVnOEmoRr2e7zP+5sEya8Z+hBHJVKxN3AghYiviXWdKrL3R5qviF4tE+SBvQ
        -----END CERTIFICATE-----

  discovery:
    base_dn: '*'
  labels:
    teleport.internal/resource-id: 6d22610a-c022-4de7-b136-cd91016b1a88
```





```yaml
version: v3
teleport:
  nodename: ip-10-103-230-19
  data_dir: /var/lib/teleport
  ca_pin: "sha256:87d32aa01fb07f33982f2f569a765dcfd949ac00f0e606747de7182d3db935d2"
  auth_token: f8ec0df2587acfab121bd4f086b953d9
  proxy_server: teletest.snwl.com:443
  log:
    output: stderr
    severity: INFO
    format:
      output: text
  diag_addr: ""

auth_service:
  enabled: "no"
ssh_service:
  enabled: "yes"
proxy_service:
  enabled: "no"
  https_keypairs: []
  acme: {}

windows_desktop_service:
  enabled: yes
  ldap:
    addr:     '10.103.230.53:636'
    domain:   'teleport.dev'
    username: 'TELEPORT\svc-teleport'
    server_name: 'WIN-DK5CL2SUTL9.teleport.dev'
    insecure_skip_verify: false
    ldap_ca_cert: |
        -----BEGIN CERTIFICATE-----
        MIIDhTCCAm2gAwIBAgIQMkzkplo89aZMoKhy7XJ85TANBgkqhkiG9w0BAQsFADBV
        MRMwEQYKCZImiZPyLGQBGRYDZGV2MRgwFgYKCZImiZPyLGQBGRYIdGVsZXBvcnQx
        JDAiBgNVBAMTG3RlbGVwb3J0LVdJTi1ESzVDTDJTVVRMOS1DQTAeFw0yMjExMjUw
        MjIzNThaFw0yNzExMjUwMjMzNTNaMFUxEzARBgoJkiaJk/IsZAEZFgNkZXYxGDAW
        BgoJkiaJk/IsZAEZFgh0ZWxlcG9ydDEkMCIGA1UEAxMbdGVsZXBvcnQtV0lOLURL
        NUNMMlNVVEw5LUNBMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0iSH
        hHPRIG9AV/CeGCiWuA/oTJCXrVwsUMKehVfLMYNHhdmriO2g/1pf5l/63jO2+Wno
        vEAviMKXnuTai5vRDZMXHXhmyOKUPM0wUdbX8P839YtaVwG9o55N+QEp9oNVJZ2n
        Os+iP5HH8AJ2Ygh6oprGEf1iqguc9b+Q/OUN9kLAexDwSMiKxBKk8RIstwc8ECdf
        0NcFIIg8vnGn8Xae0k/lOpss8GSHk+VbNDC4sAOnyrqzlMI0Fjz9EG1QgwzVTVNu
        4XyKLYU2kCerSdvzH1XcEZvmHKkRfoOgxJ/N3MWXRzWlWnCa6ttOvZNiPSOTuNHZ
        0QwRIAa+kYnRGbQRCQIDAQABo1EwTzALBgNVHQ8EBAMCAYYwDwYDVR0TAQH/BAUw
        AwEB/zAdBgNVHQ4EFgQUzSmJ85OpCDbjcV+UnXF4m0qwEdEwEAYJKwYBBAGCNxUB
        BAMCAQAwDQYJKoZIhvcNAQELBQADggEBABeDJ0pkQRmz1V1MEejeChsTG1DFBnHo
        O2EERGSsJQ9HJnxvq0wwkMXM+xytdpDTInuTqczI/5c7C/kitXGUunfCYgSo4Pmw
        nxzOS4Hd47ORnymgxX/Sie18weZorSLJYvhYRXPYoppGeW+f2TVyGROC5MjXeAiU
        A5puDZ6Y3bsWu6qTuOj7EZCbhh0n15oJfOa754mp0Wj6oBKlFU7Ti/GL/CfUNmZK
        Qo3MlMDVCVsLTWFEsjMe6MpFkCCggk1+8g+XeAGjA5tu/coiCRlnavQN6+h8PoX+
        xwvn4AIU+smWgUXCJ/DlQfXx0NP0AcMcvvT1fM7Lpk8o0R3+MOHkHBs=
        -----END CERTIFICATE-----

  discovery:
    base_dn: '*'
  labels:
    teleport.internal/resource-id: c654dd8c-6d75-426b-bd31-033f33ab21ed
```





```yaml
version: v3
teleport:
  nodename: ip-10-103-12-249
  data_dir: /var/lib/teleport
  ca_pin: "sha256:87d32aa01fb07f33982f2f569a765dcfd949ac00f0e606747de7182d3db935d2"
  auth_token: 1bccc6ffed303f38d3e53dc780cf9acd
  proxy_server: teletest.snwl.com:443
  log:
    output: stderr
    severity: INFO
    format:
      output: text
  diag_addr: ""

auth_service:
  enabled: "no"
ssh_service:
  enabled: "yes"
  commands:
    - name: hostname
      command: [ hostname ]
      period: 1h0m0s
proxy_service:
  enabled: "no"
  https_keypairs: []
  acme: {}
  
app_service:
    enabled: yes
    debug_app: true
    apps:
    - name: "fw12249"
      uri: "https://10.103.12.249"
      public_addr: "fw12249.teletest.snwl.com"
      insecure_skip_verify: true
      rewrite:
        redirect:
        - "10.103.12.249"
        - "fw12249.teletest.snwl.com"
      labels:
        env: "prod"
      commands:
      - name: "os"
        command: ["/usr/bin/uname"]
        period: "5s"
```



```shell
"$ curl -X POST https://usage.teleport.dev -F OS=${machine} -F use-case=\"access my ...\" -F email=\"alice@example.com\""
```

### 1. 添加群辉节点

1. 创建 token

   ```shell
   # 添加用户
   tctl users add xbu --roles=editor,access --logins=root,xbu
   # 生成一个node和app类型的token
   tctl tokens add --type=node,app --ttl=1d
   # 查看 token
   tctl tokens ls
   ```

2. 安装证书

   下载证书

   ```bash
   mkdir -p /etc/teleport/certs/
   openssl s_client -showcerts -connect 374339855.xyz:443 -servername 374339855.xyz < /dev/null 2>/dev/null | openssl x509 -outform PEM > /etc/teleport/certs/374339855.xyz.crt
   
   echo | openssl s_client -CAfile /etc/teleport/certs/374339855.xyz.crt -connect 374339855.xyz:443 -servername 374339855.xyz
   ```

   

   ```shell
   # 将 CA复制到/var/db/ca-certificates/ 目录下，这是群辉的自动更新目录，群辉不是/usr/local/share/ca-certificates/
   sh-4.4# cat /usr/syno/bin/update-ca-certificates.sh
   CERTSDIR=/usr/share/ca-certificates
   USERCERTSDIR=/usr/syno/etc/security-profile/ca-bundle-profile/ca-certificates/
   AUTOUPDATECERTSDIR=/var/db/ca-certificates/
   
   cp cacert.pem /var/db/ca-certificates/xbu-cacert.crt
   
   # 执行更新，群辉是个shell脚本
   sh-4.4# update-ca-certificates.sh
   rehash: warning: skipping ca-certificates.crt,it does not contain exactly one certificate or CRL
   138 added, 0 removed; done.
   
   # 查看更新结果
   sh-4.4# wget https://374339855.xyz
   --2023-11-27 23:32:35--  https://374339855.xyz/
   Resolving 374339855.xyz... 190.92.232.47
   Connecting to 374339855.xyz|190.92.232.47|:443... connected.
   HTTP request sent, awaiting response... 302 Found
   Location: /web [following]
   --2023-11-27 23:32:36--  https://374339855.xyz/web
   Reusing existing connection to 374339855.xyz:443.
   HTTP request sent, awaiting response... 200 OK
   Length: 698 [text/html]
   Saving to: 'index.html'
   
   index.html                          100%[=================================================================>]     698  --.-KB/s    in 0s
   
   2023-11-27 23:32:36 (368 MB/s) - 'index.html' saved [698/698]
   ```

   

3. 添加的 yaml

   ```yaml
   version: v3
   teleport:
     nodename: synology
     data_dir: /var/lib/teleport
     ca_pin: "sha256:f908bcd836f586cc1822c184b05e7196afc06c64d6dc4902fdc93b993d2082b1"
     auth_token: 2dd260b199fdcf62ed6decf1151d8880
     proxy_server: 374339855.xyz:443
     log:
       output: stderr
       severity: INFO
       format:
         output: text
     diag_addr: ""
   
   auth_service:
     enabled: "no"
   ssh_service:
     enabled: "yes"
     commands:
       - name: hostname
         command: [ hostname ]
         period: 1h0m0s
   proxy_service:
     enabled: "no"
     https_keypairs: []
     acme: {}
     
   app_service:
       enabled: yes
       debug_app: true
       apps:
       - name: "synology-dashboard"
         uri: "http://192.168.2.106:5000/"
         public_addr: "synology.374339855.xyz"
         insecure_skip_verify: true
         rewrite:
           redirect:
           - "192.168.2.106"
           - "synology.374339855.xyz"
         labels:
           env: "prod"
         commands:
         - name: "os"
           command: ["/usr/bin/uname"]
           period: "5s"
       - name: "emby"
         uri: "http://192.168.2.106:8096/"
         public_addr: "emby.374339855.xyz"
         insecure_skip_verify: true
         rewrite:
           redirect:
           - "192.168.2.106"
           - "emby.374339855.xyz"
         labels:
           env: "prod"
         commands:
         - name: "os"
           command: ["/usr/bin/uname"]
           period: "5s"
   ```

   

4. 启动应用

   ```shell
   teleport start -c teleport-config.yml
   
   nohup SSL_CERT_FILE=/root/cacert.pem teleport start -c config.yml > ./nohup.log 2>&1 &
   ```

   使用命令添加节点
   
   ```bash
   teleport start \
      --roles=node \
      --nodename=ip-192-168-200-10 \
      --token=2dd260b199fdcf62ed6decf1151d8880 \
      --auth-server=374339855.xyz:443 --insecure
   ```
   
   

```yaml
version: v3
teleport:
  nodename: node
  data_dir: /var/lib/teleport
  ca_pin: "sha256:f908bcd836f586cc1822c184b05e7196afc06c64d6dc4902fdc93b993d2082b1"
  auth_token: 2dd260b199fdcf62ed6decf1151d8880
  proxy_server: 374339855.xyz:443
  log:
    output: stderr
    severity: INFO
    format:
      output: text
  diag_addr: ""

auth_service:
  enabled: "no"
ssh_service:
  enabled: "yes"
  commands:
    - name: hostname
      command: [ hostname ]
      period: 1h0m0s
proxy_service:
  enabled: "no"
  https_keypairs: []
  acme: {}
  
app_service:
    enabled: yes
    debug_app: true
    resources:
    - labels:
        "*": "*"
    apps:
    - name: "grafana"
      uri: "http://10.103.12.238:3000/"
      public_addr: "grafana.374339855.xyz"
      insecure_skip_verify: true
      rewrite:
        redirect:
        - "10.103.12.238"
        - "grafana.374339855.xyz"
        headers:
        - "Origin: https://grafana.374339855.xyz" 
        - "Host: grafana.374339855.xyz"
      labels:
        env: "prod"
      commands:
      - name: "os"
        command: ["/usr/bin/uname"]
        period: "5s"
```

