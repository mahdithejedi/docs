`sudo kill $(sudo ps -aux | grep <APP_NAME>|grep -o '[[:digit:]]*' | head -1)`

kill the first numerica result of ps with specific app name