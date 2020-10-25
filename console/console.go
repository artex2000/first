package console

import (
        "github.com/gdamore/tcell/v2"
       )

type Console struct {
    screen tcell.Screen
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
    con.screen.Clear()
    return &con, nil
}

func (c *Console) Close() {
    c.screen.Fini()
}
