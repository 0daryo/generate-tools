## Tool to generate interface, struct, constructor, function at once
## How to Use
1. clone this repository and cd ./ifgen
2. ```go install```
3. run ```ifgen -c```
4. type information, following guide
5. output is shown in stdout. if you specified ```-c``` option, output is copied to clipboard.

## Options
```
 % ifgen -h                                                                 
Usage of ifgen:
  -c    copy to clipboard
```

## Output is like below
```sh
pkg name: service
interface name: interface name: User
Function Name(if nothing type ! and enter): Get
Arg(if nothing type ! and enter): id
Type: string
Arg(if nothing type ! and enter): !
Return(if nothing type ! and enter): user
Type: User
Return(if nothing type ! and enter): !
Function Name(if nothing type ! and enter): !
```

```go
package service

type User interface {
	Get(
		id string,
	) (
		user User,
	)
}

type user struct{}

func NewUser() *user {
	return &user{}
}

func (s *user) Get(
	id string,
) (
	user User,
) {
	return
}

```