# gomo

A minimal CLI [Pomodoro](https://en.wikipedia.org/wiki/Pomodoro_Technique) utility written in Go.

### stack

[bubbletea](https://github.com/charmbracelet/bubbletea)
[cobra](https://github.com/spf13/cobra)


## TODO

### config

- [ ] help menu at bottom (q/ctrl+c/esc to quit)


### start

- [ ] initialize config before start (new init command?)
- [ ] pomodoro cycle (25 -> 10 -> 25 -> 10 -> 25 -> 10 -> 25 -> 20 -> repeat) 
    - [ ] sound on timer end
    - [ ] store current duration and change accordingly for subsequent timers
- [ ] mm:ss display in `StartModel.go`
    - [ ] make it look nice
- [ ] help menu at bottom (q/ctrl+c/esc to quit)
