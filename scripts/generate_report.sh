#! /usr/bin/bash

go build -o bin/invoker main.go

mkdir reports

rps=(5 10 15 20 25 30 35 40 45 50)
duration=(100 200 300 400 500 600 700 800 900 1000)

for r in "${rps[@]}"; do
  for d in "${duration[@]}"; do
    echo "Running with RPS=$r and Duration=$d"
    bin/invoker run --rps $r --run_time 12000 --duration $d ${ENDPOINT:-"localhost:8080"} --outputPath "reports/invoke_rps${r}_dur${d}.log"
  done
done

source visualization/venv/bin/activate
pip install -r visualization/requirements.txt

for r in "${rps[@]}"; do
  for d in "${duration[@]}"; do
    python3 visualization/visualization.py "reports/invoke_rps${r}_dur${d}.log" "reports/report_rps${r}_dur${d}.png"
  done
done

rm bin/invoker
