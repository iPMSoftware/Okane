module okane

go 1.15

replace okane/database/utils => ./database/utils

replace okane/database/ware => ./database/ware

replace okane/menu => ./menu

replace okane/hardware => ./hardware

require (
	github.com/eiannone/keyboard v0.0.0-20200508000154-caf4b762e807
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect

)
