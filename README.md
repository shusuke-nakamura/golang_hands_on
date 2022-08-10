# golang_hands_on

## Fyne
https://developer.fyne.io/
【前提】
・「gwsl」がインストール済みであること
・ 以下のパッケージをインストールすること
sudo apt install x11-apps
sudo apt install libxcursor-dev
sudo apt install libxrandr-dev
sudo apt install libxinerama-dev
sudo apt install libxi-dev
sudo apt install libglx-dev
sudo apt install libgl-dev
sudo apt install libxxf86vm-dev

### 1. インストール方法
コマンド: go get fyne.io/fyne/v2
### 2. 開発方法
go mod init MODULE_NAME
go get fyne.io/fyne/v2
go mod tidy