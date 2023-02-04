# WoW ElvUI addon updater/Installer

This project is just for update your version of ElvUI addon of World of Warcraft or install it quickly

For use it, place the binary in: `World of Warcraft\Interface\<here>\Addons`. You can create shortcut for the program and launch it everywhere !

## TODO

- [x] Check remote version on ElvUI website
  * Upgrade the check without input of major version
- [x] Check if remote version is different between our and it
- [x] Download `.zip` file if the previous condition is true
- [x] Unzip `.zip` file downloaded
- [x] Only replace file and no delete the older ones

## Enhancement program

I try to use the **CurseForge** API to update all addons already installed

## TDOD

- [x] Update addons already installed
  - [x] List all addons
  - [ ] Only get the addon dir, not the annexes dirs
  - [ ] Add **CurseForge** API to update addon
- [ ] Search and install addons
  - [ ] CLI tool ? Or UI tool ?
  - *Working ....*
