module vnet

go 1.16

require (
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
	lib/ethernet v0.0.0
	lib/water v0.0.0
)

replace lib/ethernet => ./lib/ethernet
replace lib/water => ./lib/water
