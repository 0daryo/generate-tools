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
	name string, age int64, url string, 
)User{
	return User{ 
		Name: name,
		Age: age,
		URL: url,
  }
}

func TestUser_NewUser(t *testing.T){
	t.Parallel()
	type in struct{ name string
		age int64
		url string
		
	}
	tests := []struct{
		name     string
		in       in
		out      User
	}{
		{
			name:    "success",
			in:    in{ 
				name: "name001",age: 15,url: "url001",
			},
			out:   User{ 
				Name: "name001",Age: 15,URL: "url001",
			},
		},
	}
	for _,tt :=range tests{
		tt := tt
		t.Run(tt.name, func(t *testing.T){
			t.Parallel()
			out := NewUser(
				"name001",15,"url001",
			)
			assert.Equal(t, tt.out, out)
		})
	}
}
```