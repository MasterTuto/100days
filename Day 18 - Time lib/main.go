package main

import (
	"fmt"
	"time"
)

func RememberMe(channel <-chan time.Time, t *time.Ticker) {
	currentTime := <-channel
	fmt.Printf("I have been remembered, at %v\n", currentTime)
	t.Stop()
}

func NotifyMe(ticker *time.Ticker) {
	for currentTime := range ticker.C {
		fmt.Printf("I have been notified, at %v\n", currentTime)
	}
}

func testTickAndAfter() {
	fmt.Println("[Test time.Tick && time.After]")
	ticker := time.NewTicker(1 * time.Second)
	go NotifyMe(ticker)

	afterChannel := time.After(10 * time.Second)
	go RememberMe(afterChannel, ticker)

	time.Sleep(11 * time.Second)
	fmt.Print("\n\n")
}

func showDuration(duration time.Duration, name string) {
	fmt.Printf(
		"[%s]\nHours:%f\nMinutes:%f\nSeconds: %f\nMilliseconds: %d\nMicroseconds: %d\nNanosecond: %d\n\n",
		name,
		duration.Hours(),
		duration.Minutes(),
		duration.Seconds(),
		duration.Milliseconds(),
		duration.Microseconds(),
		duration.Nanoseconds(),
	)
}

func testDuration(initialTime time.Time) {
	fmt.Println("[Test duration]")
	duration1, _ := time.ParseDuration("1h20m")
	duration2, _ := time.ParseDuration("2h5m")
	showDuration(duration1, "Duration 1")
	showDuration(duration2, "Duration 2")

	fmt.Printf("Time elapsed since beginning of program: %v\n", time.Since(initialTime))
	fmt.Printf("Time elapsed since beginning until now: %v\n", time.Until(time.Now()))
	fmt.Printf("Time subtracted from to beginning %v\n", time.Now().Sub(initialTime))
	fmt.Print("\n\n")
}

func testMonth() {
	fmt.Println("[Test month]")
	for i := 1; i <= 12; i++ {
		fmt.Printf("Month %d: %s\n", i, time.Month(i))
	}

	fmt.Print("\n\n")
}

func testTime() {
	currentDate := time.Now()
	firstDayOfTheYear := time.Date(
		currentDate.Year(), time.January, 1, 0, 0, 0, 1, currentDate.Location(),
	)

	fmt.Printf(
		"Time elapsed since the beginning of the year: %v\nDays: %f, Months: %f\n",
		currentDate.Sub(firstDayOfTheYear),
		currentDate.Sub(firstDayOfTheYear).Hours()/24,
		currentDate.Sub(firstDayOfTheYear).Hours()/24/30,
	)

	fmt.Println("Two hours from now will be: ", currentDate.Add(2*time.Hour))
	fmt.Println("One year and half from now will be: ", currentDate.AddDate(1, 6, 0))

	if time.Now().After(currentDate) {
		fmt.Println("time.Now() is after the initial time of the function")
	} else if time.Now().Before(currentDate) {
		fmt.Println("time.Now() is before the initial time of the function")
	}

	h, m, s := time.Now().Clock()
	fmt.Printf("Now it is: %d:%d:%d\n", h, m, s)
	fmt.Print("\n\n")
}

func testTimer() {
	fmt.Println("[Test time.Timer]")
	timer1 := time.NewTimer(5 * time.Second)
	<-timer1.C

	fmt.Println("Timer 1 has ended!")

	timer2 := time.AfterFunc(5*time.Second, func() {
		fmt.Println("Timer 2 has ended!")
	})

	time.Sleep(3 * time.Second)
	timer2.Stop()
	fmt.Println("Timer 2 has been forced to stop")
	fmt.Print("\n\n")
}

func main() {
	initialTime := time.Now()
	testTickAndAfter()
	testDuration(initialTime)
	testMonth()
	testTime()
	testTimer()
}
