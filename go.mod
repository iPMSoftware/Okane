module okane

go 1.15

replace okane/database/utils => ./database/utils

replace okane/database/ware => ./database/ware

replace okane/menu => ./menu

replace okane/hardware => ./hardware

replace okane/mock => ./mock

replace okane/sim => ./simulator

replace okane/reports => ./reports

replace okane/configuration => ./configuration

replace okane/texts => ./texts

require (
	github.com/eiannone/keyboard v0.0.0-20200508000154-caf4b762e807
	github.com/llgcode/draw2d v0.0.0-20200930101115-bfaf5d914d1e
	golang.org/x/image v0.0.0-20210220032944-ac19c3e999fb // indirect
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect

)
