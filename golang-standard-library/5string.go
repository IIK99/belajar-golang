package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Iik Muhammad Ikmal", "Iik"))
	// true
	fmt.Println(strings.Split("Iik Muhammad Ikmal", " "))
	// [Iik Muhammad Ikmal]
	fmt.Println(strings.Join([]string{"Iik", "Muhammad", "Ikmal"}, " "))
	// Iik Muhammad Ikmal
	fmt.Println(strings.ToUpper("Iik Muhammad Ikmal"))
	// IIK MUHAMMAD IKMAL
	fmt.Println(strings.ToLower("Iik Muhammad Ikmal"))
	// iik muhammad ikmal
	fmt.Println(strings.TrimSpace("     Iik Muhammad Ikmal    "))
	// Iik Muhammad Ikmal
	fmt.Println(strings.Trim("Iik Muhammad Ikmal", "Iik"))
	//  Muhammad Ikmal
	fmt.Println(strings.Clone("Iik Muhammad Ikmal"))
	// Iik Muhammad Ikmal
	fmt.Println(strings.Index("Iik Muhammad Ikmal", "Iik"))
	// 0
	fmt.Println(strings.LastIndex("Iik Muhammad Ikmal", "Iik"))
	// 8
	fmt.Println(strings.Replace("Iik Muhammad Ikmal", "Iik", "Ikmal", 1))
	// Ikmal Muhammad Ikmal
	fmt.Println(strings.Repeat("Iik", 3))
	// IikIikIik
	fmt.Println(strings.ReplaceAll("Dia milik saya, dia bersama saya", "saya", "kamu"))
	// Dia milik kamu, dia bersama kamu
	fmt.Println(strings.Count("Iik Muhammad Ikmal", "Iik"))
	// 1	
}