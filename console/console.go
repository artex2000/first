package console

import (
        "time"
        "github.com/gdamore/tcell/v2"
       )

type Console struct {
    screen tcell.Screen
    events chan Event
    quit   chan struct{}
}

type Event interface {
    When() time.Time
}

type KeyEvent struct {
    timestamp time.Time
    KeyCode int16
    ModMask int16
    Char    rune
}

func (key KeyEvent) When() time.Time {
    return key.timestamp
}

func NewConsole() (*Console, error) {
    var con Console
    var err error
    con.screen, err = tcell.NewScreen()
    if err != nil {
        return nil, err
    }

    err = con.screen.Init()
    if err != nil {
        return nil, err
    }

    con.events = make(chan Event, 10)
    con.quit = make(chan struct{})
    con.screen.Clear()

    go pollEvents(&con)

    return &con, nil
}

func (c *Console) Close() {
    c.screen.Fini()
    close(c.quit)
}

func pollEvents(c *Console) {
    for {
        select {
            case <-c.quit:
                close(c.events)
                return
            case <-time.After(time.Millisecond * 10):
        }
        var ev Event
        _ev := c.screen.PollEvent()
        switch e := _ev.(type) {
            case *tcell.EventKey:
                ev = KeyEvent{ e.When(), int16(e.Key()), 
                        int16(e.Modifiers()), e.Rune() }
        }
        c.events <- ev
    }
}
