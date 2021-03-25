## Tool to generate interface, struct, constructor, function at once
## How to Use
1. clone this repository and cd ./mgen
2. ```go install```
3. run ```mgen -c```
4. type information, following guide
5. output is shown in stdout. if you specified ```-c``` option, output is copied to clipboard.

## Options
```
 % mgen -h                                                                 
Usage of ifgen:
  -c    copy to clipboard
```

## Output is like below
```sh
 % mgen -c
struct name: User
Field(if nothing type ! and enter): Name
Type: string
Field(if nothing type ! and enter): Age
Type: int64
Field(if nothing type ! and enter): URL
Type: string
Field(if nothing type ! and enter): !
```

```go
type User struct{ 
	Name string
	Age int64
	URL string
}

func NewUser(
	name string, age int64, uRL string, 
)User{
	return User{ 
		Name: name,
		Age: age,
		URL: uRL,
  }
}
```