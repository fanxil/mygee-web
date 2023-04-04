module mygee

go 1.19

//require：引用哪些包
require gee v0.0.0
//replace：替换一些包的下载和引用路径
replace gee => ./web/gee
//exclude：不下载和引用哪些包