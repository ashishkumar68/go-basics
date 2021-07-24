package main

import (
  "fmt"
)

func main() {
  key := "UK"
  timezones := map[string]string{
    "AU" : "Australia/Sydney",
    "NZ" : "Pacific/Auckland",
    "UK" : "Europe/London",
  }

  fmt.Println(timezones)
  if timezone, ok := timezones[key]; ok {
    fmt.Printf("Timezone with key: %s was found: %s", key, timezone)
    fmt.Println()
  }

}
