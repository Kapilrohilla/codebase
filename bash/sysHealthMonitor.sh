#!/bin/bash

function getCPUUsage() {
    cpuUsage=$(top -n 1 | awk '/Cpu\(s\):/ {print 100-$8}')
    echo "CPU: $cpuUsage"
}

function getRamUsage() {
    ramUsage=$(free | awk '/Mem/ {print $3 / $2  * 100}')
    echo "RAM: $ramUsage"
}

function getTopProcess() {
    pro=$(top -n 1 | awk 'NR==8, NR==18 {print "\nProcessId: " $2 ", User: " $3 ", CPU Usage: " $10 ", Memory Usage: " $11 ", Process:" $12}')
    echo $pro
}

getCPUUsage
getRamUsage
getTopProcess
