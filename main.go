package main

import (
    "time"
    "context"
    "github.com/chromedp/chromedp"
    "github.com/rs/zerolog/log"
)

func logActionFunc(msg string) chromedp.ActionFunc {
    return chromedp.ActionFunc(func(c context.Context) error{
        log.Debug().Msg(msg)
        return nil
    })
}
    

func main() {
    url := "https://practicetestautomation.com/practice-test-login"
    username := "student"
    password := "Password123"

    opts := append(chromedp.DefaultExecAllocatorOptions[:],
        chromedp.Flag("headless", true),
    )

    execAlloc, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
    defer cancel()

    ctx, cancel := chromedp.NewContext(execAlloc) // a chromedp context, not the std one
    defer cancel()

    ctx, cancel = context.WithTimeout(ctx, time.Duration(300) * time.Second)
    defer cancel()

    if err := chromedp.Run(ctx,
        logActionFunc("navigate url"),
        chromedp.Navigate(url),
        
        logActionFunc("waitvisible username"),
        chromedp.WaitVisible("input[name=username]"),
        
        logActionFunc("sendkeys username"),
        chromedp.SendKeys("input[name=username]", username),

        logActionFunc("sendkeys password"),
        chromedp.SendKeys("input[name=password]", password),

        logActionFunc("click #submit.btn"),
        chromedp.Click("#submit.btn", chromedp.NodeVisible),

        logActionFunc("waitvisible h1"),
        chromedp.WaitVisible("//h1[contains(.,'Logged In Successfully')]"),

    ); err != nil {
        log.Fatal().Err(err).Msg("failure encountered")
        return
    }
}
