#!/bin/bash
pkill -f 'node_modules/.bin/vite' 2>/dev/null || true
cd /app/src/frontend
npm install --silent
nohup npm run dev > /tmp/vite.log 2>&1 &
disown $!
