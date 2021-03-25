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
```go
package service

type User interface {
	Get(
		userID string, 
	) (
		User User, 
	)
}

type user struct{}

func NewUser()*user {
	return &user{}
}

func (s *user) Get(
	userID string, 
) (
	User User, 
) {
	return
}
```