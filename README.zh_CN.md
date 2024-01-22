# palworld-save-backup

一个用于备份《幻兽帕鲁》服务器存档的工具，配合 Windows 任务计划程序，或者 Linux Crontab 使用。

《幻兽帕鲁》服务器并不支持自动备份，所以我做了这个工具来让你保护你的存档。

## 使用

在跟随下面的步骤前，请下载正确的 Release 包：

* Windows（或者 Windows Server）：下载文件名包含 `windows-amd64` 的 zip 压缩包
* Linux：下载文件名包含 `linux-amd64` 的 tar.gz 档案

> 对于那些使用 arm64 架构的操作系统的人，我很抱歉目前我不提供 arm64 版本的 Release 包。

### Linux

> 在这个范例中：
>
> * 我会解压 tar.gz 档案到 `/home/steam/palworld-save-backup-tool`
> * 《幻兽帕鲁》服务器已经下载到了 `/home/steam/Steam/steamapps/common/PalServer`
> * 我会把备份存放在 `/home/steam/palworld-save-backups`
>
> 这些可以根据你的情况更改。

1. `tar -zxf palworld-save-backup-linux-amd64-${version}.tar.gz -C /home/steam/palworld-save-backup-tool`
2. 创建配置文件 `/etc/palworld-save-backup/config.toml` 并填入下面的内容，参数根据你的情况更改
   ```toml
   # SavePath
   #   《幻兽帕鲁》服务器的绝对路径
   SavePath = "/home/steam/Steam/steamapps/common/PalServer/Pal/Saved"
   
   # BackupPath
   #   你要保存备份的目录的绝对路径，程序会自动创建它
   BackupPath = "/home/steam/palworld-save-backups"
   
   # DaysKeep
   #   设置一个你需要保存备份的“天数”
   #   任何早于这个天数以前的备份都会被清理
   #   默认设置保存最近 5 天的备份
   #
   #   例如，现在的时间是 2024/01/20 12:00:00，任何在 2024/01/15 12:00:00 以前创建的备份都会被删除
   DaysKeep = 5
   ```
3. 确保你已经安装或启动了 `crontab` 服务，输入 `crontab -e` 编辑定时任务
4. 添加一行 `0 0-23/1 * * * /home/steam/palworld-save-backup-tool/linux-amd64/palworld-save-backup`
   * 这将会添加一个“每小时运行备份工具”的任务，你可以随意更改 cron 表达式，请遵循 crontab 的使用方法来配置
   * 后半部分执行命令可以根据你的情况更改到 `palworld-save-backup` 的绝对路径上
5. 保存并退出编辑器

### Windows / Windows Server

> 在这个范例中：
>
> * 我会解压 zip 压缩包到 `C:\Users\administrator\palworld-save-backup-tool`
> * 《幻兽帕鲁》服务器已经下载到了 `C:\steamcmd\steamapps\common\PalServer`
> * 我会把备份存放在 `C:\palworld-save-backups`
>
> 这些可以根据你的情况更改。

1. 双击 zip 压缩包，把 `windows-amd64` 复制到 `C:\Users\administrator\palworld-save-backup-tool`
2. 创建配置文件 `C:\Users\administrator\palworld-save-backup-tool\windows-amd64\config.toml` 并填入下面的内容，参数根据你的情况更改
   ```toml
   # SavePath
   #   《幻兽帕鲁》服务器的绝对路径
   SavePath = "C:\\steamcmd\\steamapps\\common\\PalServer"
   
   # BackupPath
   #   你要保存备份的目录的绝对路径，程序会自动创建它
   BackupPath = "C:\\palworld-save-backups"
   
   # DaysKeep
   #   设置一个你需要保存备份的“天数”
   #   任何早于这个天数以前的备份都会被清理
   #   默认设置保存最近 5 天的备份
   #
   #   例如，现在的时间是 2024/01/20 12:00:00，任何在 2024/01/15 12:00:00 以前创建的备份都会被删除
   DaysKeep = 5
   ```
3. 打开 `任务计划程序` 并创建一个任务
   * 在右手边的 `操作` 一栏，点击 `创建任务...`
   * 在 `常规` 选项卡，`名称` 填入 `AutoBackupPalworldSaves`
   * 在 `常规 > 安全选项`，选中 `不管用户是否登录都要运行(W)`，然后勾选 `不存储密码(P)。该任务将只有访问本地计算机资源的权限。`
   * 接下来到 `触发器`，创建一个可以间隔一段时间执行的触发器，关于如何设置，请参考微软的任务计划程序的文档
   * 接下来到 `操作`，新建一个操作
      * 新建操作中，`操作(I)` 选择 `启动程序`
      * `程序或脚本(P)` 的输入框，填入 `palworld-save-backup.exe`
      * `起始于(可选)(T)` 填入 `C:\Users\administrator\palworld-save-backup-tool\windows-amd64`，这个值取决于你把 `palworld-save-backup.exe` 放在了哪个文件夹，需要绝对路径
   * 点击 `确定` 新建任务
   * 回到 `任务计划程序`，在右手边的 `操作` 一栏，点击 `启用所有任务历史记录`，这将允许你跟踪任务计划的运行情况
