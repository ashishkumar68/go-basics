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

  // Add IST timezone
  timezones["IN"] = "Asia/Kolkata"

  fmt.Println("Map before deleting elements")
  printMap(timezones)

  // delete key from map
  delete(timezones, key)
  fmt.Println("Map after deleting elements")
  printMap(timezones)
}

func printMap(timezones map[string]string) {
  for key, val := range timezones {
    fmt.Println(key, val)
  }
}
