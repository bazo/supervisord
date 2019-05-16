#!/bin/bash
# Purpose: Display various options to operator using menus
# Author: Vivek Gite < vivek @ nixcraft . com > under GPL v2.0+
# ---------------------------------------------------------------------------

# display message and pause 
pause(){
	local m="$@"
	echo "$m"
	#read -p "Press [Enter] key to continue..." key
}

# set an 
while :
do
	# show menu
	clear
	
	# take action
	
	echo '*** Top 10 Memory eating process:'; ps -auxf | sort -nr -k 4 | head -10; 
	echo; echo '*** Top 10 CPU eating process:';ps -auxf | sort -nr -k 3 | head -10; 
	echo;  pause;
	sleep 5
	
done