# palworld-save-backup

A Saves backup tool for PalServer, works with Windows Task Scheduler or Linux Crontab.

PalServer doesn't support auto backup, so I made this tool to let you secure your server's saves.

## Usage

Before following the instruction, please download the correct release package.

* Windows / Windows Server: download zip archive which filename contains `windows-amd64`
* Linux: download tar.gz archive which filename contains `linux-amd64`

> For those who using arm64 operating system, I'm sorry I don't offer an arm64 version release at this time.

### Linux

> In this case:
> 
> * I will extract package to `/home/steam/palworld-save-backup-tool`
> * PalServer has been downloaded at `/home/steam/Steam/steamapps/common/PalServer`
> * I will save backups at `/home/steam/palworld-save-backups`
> 
> These paths can be replaced with your situation.

1. `tar -zxf palworld-save-backup-linux-amd64-${version}.tar.gz -C /home/steam/palworld-save-backup-tool`
2. Create config file to `/etc/palworld-save-backup/config.toml`
   ```toml
   # SavePath
   #   the absolute path to PalServer server saves
   SavePath = "/home/steam/Steam/steamapps/common/PalServer/Pal/Saved"
   
   # BackupPath
   #   the absolute path to where you want to keep the backup files
   BackupPath = "/home/steam/palworld-save-backups"
   
   # DaysKeep
   #   how long of days that you want to keep the recently backups
   #   any backups earlier than the days will be deleted
   #   as default we set it to 5 days
   #
   #   for example, now is 20/01/2024 12:00:00, those backups created before 15/01/2024 12:00:00 will be deleted
   DaysKeep = 5
   ```
3. Make sure you've already installed or opened `crontab` service, typing `crontab -e`
4. Add a line with `0 0-23/1 * * * /home/steam/palworld-save-backup-tool/linux-amd64/palworld-save-backup`
   * this will add a new task running on every hour to save backup, you can modify the cron expr to what ever you want, please do following crontab usage
5. Save and Exit the editor 

### Windows / Windows Server

> In this case:
>
> * I will extract package to `C:\Users\administrator\palworld-save-backup-tool`
> * PalServer has been downloaded at `C:\steamcmd\steamapps\common\PalServer`
> * I will save backups at `C:\palworld-save-backups`
>
> These paths can be replaced with your situation.

1. Double click zip archive, copy `windows-amd64` to `C:\Users\administrator\palworld-save-backup-tool`
2. Create config file to `C:\Users\administrator\palworld-save-backup-tool\windows-amd64\config.toml`
   ```toml
   # SavePath
   #   the absolute path to PalServer server saves
   SavePath = "C:\\steamcmd\\steamapps\\common\\PalServer"
   
   # BackupPath
   #   the absolute path to where you want to keep the backup files
   BackupPath = "C:\\palworld-save-backups"
   
   # DaysKeep
   #   how long of days that you want to keep the recently backups
   #   any backups earlier than the days will be deleted
   #   as default we set it to 5 days
   #
   #   for example, now is 20/01/2024 12:00:00, those backups created before 15/01/2024 12:00:00 will be deleted
   DaysKeep = 5
   ```
3. Open `Task Scheduler` and create a scheduled task
   * In `Actions` on the right hand side, click `Create Task...`
   * In `General`, typing `AutoBackupPalworldSaves` in `Name`
   * In `General > Security options`, checked `Run whether user is logged on or not`, then checked `Do not store password. The task will only have access to local compute resources.`
   * Next to `Triggers`, create a trigger and setup with what ever you want to repeat task, for more details please visit Microsoft Task Scheduler documents
   * Next to `Actions`, create a new action
     * In this new action, select `Start a program` in `Action`
     * Typing `palworld-save-backup.exe` in `Program/script` input box
     * Typing the absolute folder path of `palworld-save-backup.exe` in `Start in (optional)` input box, in this case I set it to `C:\Users\administrator\palworld-save-backup-tool\windows-amd64`
   * Click `OK` to create this task
   * Back to `Task Scheduler`, click `Enable All Tasks History` on the right hand side, this will enable all task logs and allow you to track tasks
