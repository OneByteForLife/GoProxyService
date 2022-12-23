package main

import (
	"GoProxyService/internal/app"
	"GoProxyService/pkg"
)

func init() {
	pkg.ConfigLog()
	// Проверка подключения к базе
}

// https://hidemy.name/ru/proxy-list/?country=AFALARAMAUATAZBDBYBZBTBOBRBGKHCACLCNCOCDCRCIHRCWCYCZDOECEGFIFRGEDEGTGNHNHKHUINIDIRIQILITJPKZKRLVLBLYLTMYMXMDMNMEMZNLNGNOPKPSPEPHPLPTPRRORURSSGSKZAESSECHTWTHTRUGUAAEGBUSUZVEVNVGZM&maxtime=100&type=h&anon=1#list
func main() {
	app.Run()
}
