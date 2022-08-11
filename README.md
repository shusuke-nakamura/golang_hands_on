# golang_hands_on

### vscode
・補完機能が効かない場合の対処「setting.json or workspace」の設定で以下の1行を追加する。
"editor.suggest.snippetsPreventQuickSuggestions": false,

【参考サイト】
https://zenn.dev/wim/articles/vscode-go-autocomplete-not-work-parameter

### Fyne
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