### 1. win10增加暂停更新的日期

在Windows powreshell(管理员)下执行下面命令，暂停更新的日期就可以增加5000天

```
reg add "HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\WindowsUpdate\UX\Settings" /v FlightSettingsMaxPauseDays /t reg_dword /d 5000 /f
```

