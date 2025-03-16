# GoTaskTracker
It is usual CLI task tracker. Made using Go and library cobra. Application uses JSON file as storage.
It is store data using map {"username": {"task name": "progress"}}

	https://roadmap.sh/projects/task-tracker


## Using
GoTaskTracker has 4 available commands:

	add
	delete
	show
	mark
	
Each command can receive up to 3 flags:
	
```username		--user -u```

---

```name of task	--task -t	  ```

---

```progress-mark 	--mark -m```


### add

This command adds tasks to you'r json file

Usage:

``` TaskTracker add --user {Username default "Me"} --task {Name of task default "None"} --mark {Mark of task default "in progress"}```

Examples:

``` TaskTracker add --user Me --task "go outside" --mark "in progress"```

``` TaskTracker add -u Me -t "go outside"```

``` TaskTracker add -t "go outside"```


### delete

This command deletes tasks or users

Usage:
``` TaskTracker delete --user {Username default "all"} --task {Name of task default "all"}```

Examples:

``` TaskTracker delete --user Me --task "go outside" - deletes task "go outside" from user "Me"```

``` TaskTracker delete -u Me - deletes all tasks of user "Me"```

``` TaskTracker delete -m "Blocked" - deletes all tasks```

### show

This command show you'r tasks from json file
	
#### Usage:

``` TaskTracker show --user {Username default "all"} --task {Name of task default "all"} --mark {Mark of task default "all"}```

#### Examples:
``` TaskTracker show - shows all tasks, users, and progress marks```

``` TaskTracker show -u Me - shows all tasks of user "Me"```

``` TaskTracker show -m Done - shows all tasks with progress mark "Done"```
	

#### Output:

	         Leo
	|Write poem               |In Progress              |
	|Go hiking                |Done                     |
	         Bob
	|Write report             |Done                     |
	         David
	|Clean room               |Done                     |
	|Do homework              |In Progress              |
	|Fix door handle          |Done                     |
	|Send email updates       |In Progress              |



### mark

This command changes task's progress-marks
Usage:

``` TaskTracker mark --user {Username default "all"} --task {Name of task default "all"} --mark {Mark of task default "in progress"}```

Examples:
	
``` TaskTracker mark --user Me --task "go outside" --mark "Blocked" - marks task "go outside" of user "Me" as "Blocked"```

``` TaskTracker mark -u Me -t "go outside" - marks task "go outside" of user "Me" as "Done"```

``` TaskTracker mark -u Me --mark "Blocked" - marks all tasks of user "Me" as "Blocked"```

``` TaskTracker mark -m "Blocked" - marks all tasks as "Blocked"```
	
## Installation
```
git clone https://github.com/ADecentRaccoon/GoTaskTracker.git
cd GoTaskTracker
go build -o TaskTracker
```

