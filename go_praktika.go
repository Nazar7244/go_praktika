
package main

import (
    "fmt"
    "go_praktika/greetings"
    "go_praktika/logger"
)

func main() {
    message := greetings.Hello("Nazar")
    fmt.Println(message)

    logger.Info("Бағдарлама іске қосылды")
    logger.Error("Файл табылмады")
}
