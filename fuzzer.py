#!/usr/bin/python3

import socket
import sys
import time

if len(sys.argv) < 3:
    print("\nUsage: "+ sys.argv[0] + "<HOST><PORT>\n")
    sys.exit()

host = sys.argv[1]
port = sys.argv[2]
length = 100

while (length < 10000):
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.connect((host, int(port)))
    s.recv(1024)
    s.send(b"A" * length)
    s.recv(1024)
    s.close()
    time.sleep(0)
    print("Buffer sent: "+ str(length))
    length += 100