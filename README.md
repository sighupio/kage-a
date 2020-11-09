Service needed to demo project ONE

Will expect to be deployed in an environment with another service `B`, see https://git.sighup.io/kage/one-b

Will expect a variable environment `SERVICE_B_HOST` that contains the host of its associated service called `B`

Will expose `GET /do-stuff` route, will call service B, expecting a number, will reply adding 1 to that number
