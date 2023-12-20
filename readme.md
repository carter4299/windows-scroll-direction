<h1 align="center">Change Windows Scroll Direction</h1>

### Program changes the 🐭 to use the alternate scroll direction
#### The program automates the process of going in to regedit, finding your mouse id, and changing it.

---

<a href="ChangeScroll.exe"><h3>Click to Download</h3></a>

<h3 align="center"> 🔭 How It Works ⚡ </h3>

- Be on a Windows PC
- Right click to run program as Administrator
- Have a somewhat modern mouse
- Restart your computer after running the program.

1. Get mouse ID
```powershell
Get-WmiObject Win32_PointingDevice | Select-Object DeviceID
```
2. Read in previous scroll pattern
```powershell
Get-ItemProperty -Path "HKLM:\\SYSTEM\\CurrentControlSet\\Enum\\%s\\Device Parameters" -Name "FlipFlopWheel"
```
3. Change value
```powershell
script := fmt.Sprintf(`Set-ItemProperty -Path "HKLM:\\SYSTEM\\CurrentControlSet\\Enum\\%s\\Device Parameters" -Name "FlipFlopWheel" -Value %d`, user.id, user.new_val)
```

---

<h3 align="center"> 💻 Compile Your Own 🖥 </h3>

If you dont feel comfortable downloading an .exe -> [Source Code](/source/)

### Install

```bash
git clone https://github.com/carter4299/windows-scroll-direction.git
cd windows-scroll-direction/install
go build -o ChangeScroll.exe
```
