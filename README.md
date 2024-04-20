# Go Gas Station Simulator

## Original assignment

1. Cars arrive at the gas station and wait in the queue for free station
1. Total number of cars and their arrival time is configurable
1. There are 4 types of stations: gas, diesel, LPG, electric
1. Count of stations and their serve time is configurable as interval (e.g. 2–5s) and can be different for each type
1. Each station can serve only one car at a time, serving time is chosen randomly from station's interval
1. After the car is served, it goes to the cash register
1. Count of cash registers and their handle time is configurable
1. After the car is handled (random time from register handle time range) by the cash register, it leaves the station
1. Program collects statistics about the time spent in the queue, time spent at the station and time spent at the cash register for every car
1. Program prints the aggregate statistics at the end of the simulation


## Config `config.yaml`
```yaml
cars:
  count: 500
  arrival_time_min: 1ms   # new car arrives every 1-2ms
  arrival_time_max: 2ms
stations:
  gas:
    count: 2
    serve_time_min: 2ms
    serve_time_max: 5ms
  diesel:
    count: 2
    serve_time_min: 3ms
    serve_time_max: 6ms
  lpg:
    count: 4
    serve_time_min: 4ms
    serve_time_max: 7ms
  electric:
    count: 2
    serve_time_min: 5ms
    serve_time_max: 10ms
registers:
  count: 3
  handle_time_min: 1ms
  handle_time_max: 3ms
```

## Output `output.yaml`
```yaml
Stations:
  gas:
    total_cars: 127
    total_time: 384ms
    avg_queue_time: 3ms
    max_queue_time: 15ms
  diesel:
    total_cars: 123
    total_time: 502ms
    avg_queue_time: 4ms
    max_queue_time: 16ms
  lpg:
    total_cars: 112
    total_time: 583ms
    avg_queue_time: 5ms
    max_queue_time: 31ms
  electric:
    total_cars: 138
    total_time: 964ms
    avg_queue_time: 6ms
    max_queue_time: 16ms
registers:
  total_cars: 500
  total_time: 9s
  avg_queue_time: 19ms
  max_queue_time: 32ms
```
