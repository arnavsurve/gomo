# gomo

A minimal CLI [Pomodoro](https://en.wikipedia.org/wiki/Pomodoro_Technique) utility written in Go.

### stack

[bubbletea](https://github.com/charmbracelet/bubbletea)
[cobra](https://github.com/spf13/cobra)


## TODO

### config

- [ ] help menu at bottom (q/ctrl+c/esc to quit)


### start

- [ ] (backlog) select focus/short/long?

- [ ] initialize config before start (new init command?)

- [ ] pomodoro cycle (25 -> 10 -> 25 -> 10 -> 25 -> 10 -> 25 -> 20 -> repeat) 
    - [ ] sound on timer end
    - [ ] read current duration and change accordingly for subsequent timers
        - [ ] store focus/short/long (session) history
            - [store in temp file](https://gobyexample.com/temporary-files-and-directories)
        - [ ] display session history

- [ ] mm:ss display in `StartModel.go`
    - [ ] make it look nice/big

- [ ] help menu at bottom (q/ctrl+c/esc to quit)


### styling

- [ ] different focus color (blue? or red)
- [ ] ascii art? pomo tomato?
