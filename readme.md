# 网卡信息获取
> 提供上层语言获取网卡相关信息的方法和接口

## 内置方法
- [vela.ifconfig(cnd)](###获取网卡) &emsp;获取网卡信息
- [vela.ifconfig.mac(...)](###指定mac地址网卡) &emsp;获取指定Mac网卡信息
- [vela.ifconfig.name(...)](###指定name地址网卡) &emsp;获取指定Mac网卡信息
- [vela.ifconfig.addr(...)](###指定addr地址网卡) &emsp;获取指定addr网卡信息
- [vela.ifconfig.pipe(...)](###遍历网卡) &emsp;遍历网卡信息
- [vela.ifconfig.update()](###更新缓存) &emsp;刷新缓存网卡信息
- [vela.ifconfig.flow(time , pipe)](###流量监控) &emsp;流量监控

### 获取网卡
> [][interface](###网卡信息) = vela.ifconfig(cnd) <br />
> 返回slice结构 参数:cnd 是通用过滤条件  支持 windows linux
```lua
    local s = vela.ifconfig("addr = 127.0.0.1")
    print(s.size)
    for i = 1, s.size do
        print(s[i]) 
    end
```

### 指定name的网卡
> [][interface](###网卡信息) = vela.ifconfig.name(...string) <br />
> 返回slice结构 参数:name地址信息   支持 windows linux
```lua
    local s = vela.ifconfig.name("eth0" , "eth1")
    print(s.size)
    for i = 1, s.size do
        print(s[i]) 
    end
```

### 指定addr的网卡
> [][interface](###网卡信息) = vela.ifconfig.addr(...string) <br />
> 返回slice结构 参数:addr地址信息   支持 windows linux
```lua
    local s = vela.ifconfig.addr("eth0" , "eth1")
    print(s.size)
    for i = 1, s.size do
        print(s[i]) 
    end
```

### 指定Mac的网卡
> [][interface](###网卡信息) = vela.ifconfig.mac(...string) <br />
> 返回slice结构 参数:mac地址信息   支持 windows linux
```lua
    local s = vela.ifconfig.mac("mac-xx-xxx" , "mac-yy-yyy")
    print(s.size)
    for i = 1, s.size do
        print(s[i]) 
    end
```

### 遍历网卡
> vela.ifconfig.pipe(pipe) <br />
> 遍历所有的网卡 内置参数: [interface](###网卡信息)
```lua
    vela.ifconfig.pipe(function(interface)
        print(interface.name)
        print(interface.mac)
        print(interface.ipv4)
        print(interface.ipv6)
    end)
```

### 流量监控
> procdata = vela.ifconfig.flow(time ,pipe) <br /> 返回任务对象
> 监控网卡流量 time:监控周期 单位:毫秒     pipe:处理当前流量大小逻辑  内置参数:[flow](###流量信息)
```lua
    vela.ifconfig.flow(1000 , function(flow)
        print(flow. 
    end)

```

### 流量信息
flow 流量监控是pipe获取到内置参数 &nbsp;字段信息:
- in_bytes
- in_packets
- in_error
- in_dropped
- in_bps
- in_pps
- out_bytes
- out_packets
- out_error
- out_dropped
- out_bps
- out_pps


```lua
    vela.ifconfig.flow(1000 , function(flow)
        print(in_bytes)
        print(in_packets)
        print(in_error)
        print(in_dropped)
        print(in_bps)
        print(in_pps)
        print(out_bytes)
        print(out_packets)
        print(out_error)
        print(out_dropped)
        print(out_bps)
        print(out_pps)
    end)
```


### 网卡信息
> interface 封装了网卡的信息

内置字段
- name
- flag
- index
- mac
- mtu
- ipv4    &emsp; 输出所有IPv4地址 逗号分割
- ipv6    &emsp; 输出所有IPv6地址 逗号分割

内置方法
- addr(int) &emsp;获取网卡下第几个IP地址

```lua
    --[[
        网卡1: 192.168.1.1/24,192.168.1.2/24 mac:mac-xx-xxx-xxxx name:本地网卡 flags:up|broadcast mtu:1500
        网卡2: 127.0.0.1/24 mac:mac-xx-xxx-xxxx name:回环网卡   flags:up|loopback    mtu:1500
    ]]
    
    local v = vela.ifconfig() --[网卡1 ,网卡2]
    print(v[1].name) -- 本地网卡
    print(v[1].flag) -- up|loopback 
    print(v[1].ipv4) -- 192.168.1.1,192.168.1.2
    
    print(v[1].addr(1)) -- 192.168.1.1    

```