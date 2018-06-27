# li
li - Like `ls` but for EC2 instances

## Usage 

```
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
~> li redis sentinel
Redis Sentinel   165.91.65.2
```
   
IP Address, instance status and availability-zone
```
~> li -sz redis sentinel
Redis Sentinel   165.91.65.2   running   us-west-1a

```

Detailed Information
IP Address, instance status and availability-zone
```
~> li -d redis sentinel
Redis Sentinel   165.91.65.2   i-xxxxxxxxxxxxxxxxx   running   t2.small   us-west-1a

```
