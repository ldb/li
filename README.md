# li
li - Like `ls` but for EC2 instances

## Usage 

```
-a, --address   show IP addresses
-d, --details   show all information
-h, --help      show help
-i, --id        show instance-ID
-s, --status    show instance status
-t, --type      show instance type
-z, --zone      show availability-zone
```

## Examples

IP Address only
```
~> li -a redis sentinel
Redis Sentinel   165.91.65.2
```
   
IP Address, instance status and availability-zone
```
~> li -asz redis sentinel
Redis Sentinel   165.91.65.2   running   us-west-1a

```
