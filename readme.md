<h1 align="center">Change Scroll Direction for Windows</h1>

### Program modifies the scroll direction for the 🐭 on Window's operating systems. 
#### Automates the process of going in to regedit, finding your mouse id, and changing it.
#### Alternates between Reverse Scroll and Natural Scroll.

---

<a href="https://github.com/carter4299/windows-scroll-direction/blob/main/ChangeScroll.exe">Download Here</a>

<h3 align="center"> 🔭 How It Works ⚡ </h3>

- Windows 10/11 os
- Have a somewhat modern mouse

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
